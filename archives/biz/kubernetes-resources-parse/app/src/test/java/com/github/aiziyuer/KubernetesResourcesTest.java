package com.github.aiziyuer;

import java.io.IOException;
import java.nio.charset.Charset;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.stream.Collectors;

import org.apache.commons.io.IOUtils;
import org.apache.commons.lang3.StringUtils;
import org.junit.Ignore;
import org.junit.Test;
import org.yaml.snakeyaml.DumperOptions;
import org.yaml.snakeyaml.Yaml;
import org.yaml.snakeyaml.representer.Representer;

import com.fasterxml.jackson.annotation.JsonInclude.Include;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.SerializationFeature;
import com.fasterxml.jackson.dataformat.yaml.YAMLFactory;
import com.jayway.jsonpath.Configuration;
import com.jayway.jsonpath.JsonPath;
import com.jayway.jsonpath.TypeRef;
import com.jayway.jsonpath.spi.json.JacksonJsonProvider;
import com.jayway.jsonpath.spi.mapper.JacksonMappingProvider;

import io.kubernetes.client.models.V1ConfigMap;
import io.kubernetes.client.models.V1ConfigMapVolumeSource;
import io.kubernetes.client.models.V1Container;
import io.kubernetes.client.models.V1Secret;
import io.kubernetes.client.models.V1Volume;
import io.kubernetes.client.models.V1beta2Deployment;
import lombok.Data;
import lombok.ToString.Exclude;
import lombok.extern.log4j.Log4j2;

@Log4j2
public class KubernetesResourcesTest {

	/** json对象操作 */
	private static final ObjectMapper JSON_OBJECT_MAPPER;

	/** yaml对象操作 */
	private static final ObjectMapper YAML_OBJECT_MAPPER;

	/** 美化后的打印 */
	private static final ObjectWriter PRETTY_WRITER;

	/** 普通打印 */
	private static final ObjectWriter WRITER;

	/** yaml 解析器 */
	private static final Yaml YAML;

	static {

		JSON_OBJECT_MAPPER = new ObjectMapper();
		JSON_OBJECT_MAPPER.setSerializationInclusion(Include.NON_NULL);

		YAML_OBJECT_MAPPER = new ObjectMapper(new YAMLFactory());
		YAML_OBJECT_MAPPER.enable(SerializationFeature.INDENT_OUTPUT);

		PRETTY_WRITER = JSON_OBJECT_MAPPER.writerWithDefaultPrettyPrinter();
		WRITER = JSON_OBJECT_MAPPER.writer();

		Representer representer = new Representer();
		// yaml反解对象时跳过空属性
		representer.getPropertyUtils().setSkipMissingProperties(true);
		DumperOptions options = new DumperOptions();
		options.setDefaultFlowStyle(DumperOptions.FlowStyle.BLOCK);
		options.setPrettyFlow(true);
		YAML = new Yaml(representer, options);

	}

	@Ignore
	@Test
	public void test001() throws JsonProcessingException {
		log.info("test001", PRETTY_WRITER.writeValueAsString(""));
	}

	/** 应用信息 */
	@Data
	public static class AppInfo {

		private String name = StringUtils.EMPTY;

		private Map<String, String> env = new HashMap<>();

		private Map<String, String> readableConfig = new HashMap<>();

		private Map<String, String> unreadableConfig = new HashMap<>();

		@Exclude
		private Object resource = null;
	}

	private String toJson(Object o) {

		try {
			return WRITER.writeValueAsString(o);
		} catch (JsonProcessingException e) {
			throw new RuntimeException(e);
		}

	}

	// @Ignore
	@Test
	public void resourceParse() throws IOException {

		String yamlContent = IOUtils.toString(//
				getClass().getClassLoader()//
						.getResourceAsStream("demo/deployments.yml"),
				Charset.forName("utf8"));

		// {类型: 资源的Java类}
		Map<String, Class<?>> resourceTypeMap = new HashMap<>();
		resourceTypeMap.put("Deployment", V1beta2Deployment.class);
		resourceTypeMap.put("Secret", V1Secret.class);
		resourceTypeMap.put("ConfigMap", V1ConfigMap.class);

		// {类型: {名称: 资源}
		Map<String, Map<String, Object>> resourceMap = new HashMap<>();
		resourceTypeMap.keySet().forEach(s -> resourceMap.put(s, new HashMap<>()));

		// 根据'---'作为分隔符来分割yaml的文档
		for (String resourceContent : yamlContent.split("---")) {

			if (StringUtils.isBlank(resourceContent))
				continue;

			// yaml text -> java map object -> jsonpath object
			String contentJson = toJson(YAML.load(resourceContent));
			log.info(String.format("resource: %s", contentJson));
			Object contentJsonObj = Configuration.defaultConfiguration().jsonProvider().parse(contentJson);
			// resource type
			String kind = JsonPath.read(contentJsonObj, "$.kind");
			// resource name
			String name = JsonPath.read(contentJsonObj, "$.metadata.name");

			Map<String, Object> resourceNameMap = resourceMap.get(kind);
			// 没有注册的类型不解析
			if (resourceNameMap == null) {
				log.warn(String.format("kind(%s) unsupported, skip.", kind));
				continue;
			}

			Object resource = YAML.loadAs(resourceContent, resourceTypeMap.get(kind));

			// Object resource = YAML_OBJECT_MAPPER.readValue(resourceContent,
			// resourceTypeMap.get(kind));
			resourceNameMap.put(name, resource);

		}

		log.trace(String.format("resourceMap: %s", resourceMap));

		// TODO 组件配置: 环境变量, ConfigMap, Secret
		List<AppInfo> appInfos = resourceMap.get("Deployment").entrySet().stream().map(entry -> {

			AppInfo info = new AppInfo();
			info.name = entry.getKey();
			info.resource = entry.getValue();

			Configuration conf = Configuration.builder() //
					.mappingProvider(new JacksonMappingProvider(JSON_OBJECT_MAPPER)) //
					.jsonProvider(new JacksonJsonProvider(JSON_OBJECT_MAPPER)) //
					.build();

			// 处理Volumes
			Map<String, String> voluemeMap = new HashMap<>();
			List<V1Volume> volumes = JsonPath// 实现参考: http://www.ibloger.net/article/2329.html
					.using(conf) //
					.parse(toJson((V1beta2Deployment) info.resource)) //
					.read("$.spec.template.spec.volumes[0:]", new TypeRef<List<V1Volume>>() {
					});
			volumes.stream()//
					.parallel() // 并行计算
					.peek(v -> {
						// ConfigMap
						V1ConfigMapVolumeSource configMap = v.getConfigMap();
						if (configMap == null)
							return;

						// volume的名称
						v.getName();
						// configmap的名称
						configMap.getName();

						voluemeMap.put(v.getName(), configMap.getName());

					}) //
					.peek(v -> {
						// TODO Secrets
						v.getSecret();
					}) //
			;

			// 处理容器
			List<V1Container> containers = JsonPath//
					.using(conf) //
					.parse(toJson((V1beta2Deployment) info.resource)) //
					.read( //
							"$.spec.template.spec.containers[0:]", // [0:]可以让取出来的数据是map而不是数组
							new TypeRef<List<V1Container>>() {
							});
			// 暂时只处理一个容器
			containers.stream().limit(1).peek(container -> {

				log.trace(String.format("containers.stream().limit(1).peek(container->: %s", toJson(container)));

				// TODO 环境变量
				container.getEnv().forEach(env -> {

				});
				container.getEnvFrom().forEach(envFrom -> {

				});

				// TODO 挂载点信息
				container.getVolumeMounts().forEach(volumeMount -> {

				});
				container.getVolumeDevices();

			});

			return info;

		}).collect(Collectors.toList());

		log.info(String.format("appInfos: %s", toJson(appInfos)));

		// TODO 持久化目录(非必须)

	}

}

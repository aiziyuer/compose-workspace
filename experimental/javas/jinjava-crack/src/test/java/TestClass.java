import com.google.common.base.Charsets;
import com.google.common.io.Resources;
import com.hubspot.jinjava.Jinjava;
import com.hubspot.jinjava.JinjavaConfig;
import org.junit.Test;
import org.yaml.snakeyaml.Yaml;

import java.io.IOException;
import java.io.InputStream;

public class TestClass {

    @Test
    public void testJinja2() throws IOException {

        String template =
                Resources.toString(Resources.getResource("template.j2"), Charsets.UTF_8);

        try (
                InputStream inputStream = this.getClass()
                        .getClassLoader()
                        .getResourceAsStream("input.yaml");

        ) {
            String ret = new Jinjava(
                    JinjavaConfig.newBuilder()
                            .withFailOnUnknownTokens(true)
                            .build()
            ).render(template, new Yaml().load(inputStream));
            System.out.println(ret);
        }


    }
}

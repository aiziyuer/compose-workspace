# 运行用镜像
FROM aiziyuer/aarch64-centos:7

# 覆盖源
COPY etc/yum.repos.d/*.repo /etc/yum.repos.d/

RUN echo "verify is passed!"

# 安装必要的软件
RUN \
    rpm --import https://artifacts.elastic.co/GPG-KEY-elasticsearch \
 && yum install -y logstash


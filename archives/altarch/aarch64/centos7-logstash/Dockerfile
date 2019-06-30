# 运行用镜像
FROM aiziyuer/aarch64-centos:7

# 覆盖源
COPY etc/yum.repos.d/*.repo /etc/yum.repos.d/

# 安装加速软件
RUN \
    yum install -y \
            axel yum-plugin-fastestmirror

# 安装java并设置JAVA_HOME
RUN yum install -y \
            java

# # 安装必要的软件
# RUN \
#     yum install -y logstash


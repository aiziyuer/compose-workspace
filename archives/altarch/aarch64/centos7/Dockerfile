# 打包用镜像
FROM golang:1.11 as builder

# 下载异构用的解释器
RUN \
curl -L -k \
https://github.com/multiarch/qemu-user-static/releases/download/v4.0.0/qemu-aarch64-static.tar.gz \
| tar -xvz -C /usr/bin/

# 运行用镜像
FROM centos/aarch64:7

COPY --from=builder /usr/bin/qemu-aarch64-static /usr/bin/qemu-aarch64-static

# 覆盖源
COPY etc/yum.repos.d/*.repo /etc/yum.repos.d/

# 如果是x86_64宿主机下需要采用 docker run --rm --privileged multiarch/qemu-user-static:register 来预先注册
RUN echo "verify is passed!"

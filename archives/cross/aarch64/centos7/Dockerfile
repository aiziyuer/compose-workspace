#FROM centos:7 as builder

# 设置源
#COPY ./x86_64/etc/yum.repos.d/* /etc/yum.repos.d/

# 安装golang来编译crossbuild
#RUN yum install -y \
#               golang

FROM golang:1.11 as builder

# 下载异构用的解释器
RUN \
curl -L -ko /usr/bin/qemu-aarch64-static \
https://github.com/multiarch/qemu-user-static/releases/download/v4.0.0/qemu-aarch64-static

ENV GOPATH=/

WORKDIR $GOPATH/src/github.com/KurtStam/add-aarch64-qemu

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' cross-build.go

FROM centos/aarch64:7

COPY --from=builder /usr/bin/qemu-aarch64-static /usr/bin/qemu-aarch64-static
COPY --from=builder /src/github.com/KurtStam/add-aarch64-qemu/cross-build /usr/bin/cross-build

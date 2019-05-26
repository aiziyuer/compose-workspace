#FROM centos:7 as builder

# 设置源
#COPY ./x86_64/etc/yum.repos.d/* /etc/yum.repos.d/

# 安装golang来编译crossbuild
#RUN yum install -y \
#               golang

FROM golang:1.11 as builder

ENV GOPATH=/tmp \
    GOOS=linux \
    GOARCH=amd64

WORKDIR $GOPATH/src/github.com/KurtStam/add-aarch64-qemu
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...

FROM centos/aarch64:7

# 下载异构用的解释器
RUN \
curl -L -ko /usr/bin/qemu-aarch64-static \
https://github.com/multiarch/qemu-user-static/releases/download/v4.0.0/qemu-aarch64-static


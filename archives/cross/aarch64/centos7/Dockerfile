# 打包用镜像
FROM golang:1.11 as builder

# 下载异构用的解释器
RUN \
curl -L -k \
https://github.com/multiarch/qemu-user-static/releases/download/v4.0.0/qemu-aarch64-static.tar.gz \
| tar -xvz -C /usr/bin/

ENV GOPATH=/

WORKDIR $GOPATH/src/github.com/KurtStam/add-aarch64-qemu

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' cross-build.go
RUN \ 
   ln -s cross-build cross-build-start \
&& ln -s cross-build cross-build-end


# 运行用镜像
FROM centos/aarch64:7

COPY --from=builder /usr/bin/qemu-aarch64-static /usr/bin/qemu-aarch64-static
COPY --from=builder /src/github.com/KurtStam/add-aarch64-qemu/cross-build* /usr/bin/

RUN [ "cross-build-start" ]

RUN echo "verify is passed!"

RUN [ "cross-build-end" ]

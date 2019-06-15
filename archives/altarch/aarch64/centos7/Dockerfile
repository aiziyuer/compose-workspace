# 打包用镜像
FROM golang:1.11 as builder

# 下载异构用的解释器
RUN \
curl -L -k \
https://github.com/multiarch/qemu-user-static/releases/download/v4.0.0/qemu-aarch64-static.tar.gz \
| tar -xvz -C /usr/bin/

# 运行用镜像
FROM centos/aarch64:7

# 注册解析执行器, 主要用户异构场景
COPY --from=builder /usr/bin/qemu-aarch64-static /usr/bin/qemu-aarch64-static

# 覆盖源
COPY etc/yum.repos.d/*.repo /etc/yum.repos.d/

# 如果是x86_64宿主机下需要采用 docker run --rm --privileged multiarch/qemu-user-static:register 来预先注册
RUN (cd /lib/systemd/system/sysinit.target.wants/; for i in *; do [ $i == systemd-tmpfiles-setup.service ] || rm -f $i; done); \
rm -f /lib/systemd/system/multi-user.target.wants/*;\
rm -f /etc/systemd/system/*.wants/*;\
rm -f /lib/systemd/system/local-fs.target.wants/*; \
rm -f /lib/systemd/system/sockets.target.wants/*udev*; \
rm -f /lib/systemd/system/sockets.target.wants/*initctl*; \
rm -f /lib/systemd/system/basic.target.wants/*;\
rm -f /lib/systemd/system/anaconda.target.wants/*;

### systemd requirements - to cleanly shutdown systemd, use SIGRTMIN+3
STOPSIGNAL SIGRTMIN+3

CMD ["/usr/sbin/init"]

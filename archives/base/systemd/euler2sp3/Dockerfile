# 制作指导: https://developer.huawei.com/ict/cn/site-euleros/article/yum
FROM centos:7 as builder

# 覆盖源
RUN rm -rf /etc/yum.repos.d && mkdir -p /etc/yum.repos.d
COPY etc/yum.repos.d/*.repo /etc/yum.repos.d/

ENV ROOTFS_DIR=/rootfs

# 初始化rpmdb
RUN rpm --root $ROOTFS_DIR --initdb

# 制作rootfs来
RUN \
   yum install -y --releasever=/ --installroot=$ROOTFS_DIR \
	   --setopt=tsflags='nodocs' --setopt=override_install_langs=en_US.UTF-8 \
	   bash yum \
	   acl binutils \
	   coreutils cronie cronie-anacron \
	   file filescheck groff-base \
	   hostname iproute iptables iputils \
	   less net-tools euleros-release-2.0SP3 \
	   passwd rootfiles shim sudo \
	   vim-minimal which \
	   libcroco libgomp libunistring tcp_wrapper-libs tar \
&& rm -rf $ROOTFS_DIR/var/cache/yum

FROM scratch

COPY --from=builder /rootfs/ /
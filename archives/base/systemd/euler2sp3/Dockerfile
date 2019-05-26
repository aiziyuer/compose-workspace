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
	   libcroco libgomp libunistring tcp_wrapper-libs tar 

FROM scratch

COPY --from=builder /rootfs/ /

# 定制源
RUN rm -rf /etc/yum.repos.d || true && mkdir -p /etc/yum.repos.d
COPY etc/yum.repos.d/*.repo /etc/yum.repos.d/

# 必要的定位用工具
RUN yum -y install vim net-tools tcpdump telnet lsof wget curl \
&& yum clean all

RUN (cd /lib/systemd/system/sysinit.target.wants/; for i in *; do [ $i == systemd-tmpfiles-setup.service ] || rm -f $i; done); \
rm -f /lib/systemd/system/multi-user.target.wants/*;\
rm -f /etc/systemd/system/*.wants/*;\
rm -f /lib/systemd/system/local-fs.target.wants/*; \
rm -f /lib/systemd/system/sockets.target.wants/*udev*; \
rm -f /lib/systemd/system/sockets.target.wants/*initctl*; \
rm -f /lib/systemd/system/basic.target.wants/*;\
rm -f /lib/systemd/system/anaconda.target.wants/*;

# 避免出现僵尸进程: https://developers.redhat.com/blog/2016/09/13/running-systemd-in-a-non-privileged-container/
STOPSIGNAL SIGRTMIN+3

CMD ["/usr/sbin/init"]
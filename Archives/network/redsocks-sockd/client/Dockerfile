FROM centos:7 as builder

# 覆盖源
COPY etc/yum.repos.d/*.repo /etc/yum.repos.d/

RUN yum install -y gcc make libevent libevent-devel

COPY ./dists/redsocks-release-0.5.tar.gz /tmp/redsocks-release.tar.gz

RUN \
mkdir -p /workspace \
&& tar xzvf /tmp/redsocks-release.tar.gz -C /workspace \
&& cd /workspace/redsocks* \
&& make

FROM centos:7

# 覆盖源
COPY etc/yum.repos.d/*.repo /etc/yum.repos.d/

RUN  \
rpm --rebuilddb  \
&& yum -y install net-tools tcpdump telnet lsof wget curl \
&& yum clean all

# 增加bash命令补全
RUN  \
rpm --rebuilddb  \
&& yum -y install mlocate bash-completion bash-completion-extras \
&& yum clean all

# 安装redsocks
RUN yum install -y iptables libevent
COPY --from=builder /workspace/redsocks*/redsocks /usr/local/bin/redsocks

# 加入系统service
COPY ./etc/sysconfig/* /etc/sysconfig/
COPY ./etc/systemd/system/*.service /etc/systemd/system/
# 拷贝自定义脚本
COPY ./usr/local/bin/* /usr/local/bin/
RUN chmod +x /usr/local/bin/*

RUN (cd /lib/systemd/system/sysinit.target.wants/; for i in *; do [ $i == systemd-tmpfiles-setup.service ] || rm -f $i; done); \
rm -f /lib/systemd/system/multi-user.target.wants/*;\
rm -f /etc/systemd/system/*.wants/*;\
rm -f /lib/systemd/system/local-fs.target.wants/*; \
rm -f /lib/systemd/system/sockets.target.wants/*udev*; \
rm -f /lib/systemd/system/sockets.target.wants/*initctl*; \
rm -f /lib/systemd/system/basic.target.wants/*;\
rm -f /lib/systemd/system/anaconda.target.wants/*;

RUN \
    systemctl enable container-environment-inject \
&&  systemctl enable redsocks


# 避免出现僵尸进程: https://developers.redhat.com/blog/2016/09/13/running-systemd-in-a-non-privileged-container/
STOPSIGNAL SIGRTMIN+3

CMD ["/usr/sbin/init"]
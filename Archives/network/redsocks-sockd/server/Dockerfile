FROM centos:7

# 覆盖源
COPY etc/yum.repos.d/*.repo /etc/yum.repos.d/

# 增加网络定位用的工具
RUN  \
rpm --rebuilddb  \
&& yum -y install net-tools tcpdump telnet lsof wget curl \
&& yum clean all

# 增加bash命令补全
RUN  \
rpm --rebuilddb  \
&& yum -y mlocate bash-completion bash-completion-extras \
&& yum clean all

# 开始安装sorkd
COPY ./dists/tgz-prod.dante-1.4.2-rhel72-amd64-64bit-gcc.tar.gz /tmp/dante.tar.gz
RUN tar xzvf /tmp/dante.tar.gz -C /
COPY ./lib/systemd/system/sockd.service /lib/systemd/system/

RUN (cd /lib/systemd/system/sysinit.target.wants/; for i in *; do [ $i == systemd-tmpfiles-setup.service ] || rm -f $i; done); \
rm -f /lib/systemd/system/multi-user.target.wants/*;\
rm -f /etc/systemd/system/*.wants/*;\
rm -f /lib/systemd/system/local-fs.target.wants/*; \
rm -f /lib/systemd/system/sockets.target.wants/*udev*; \
rm -f /lib/systemd/system/sockets.target.wants/*initctl*; \
rm -f /lib/systemd/system/basic.target.wants/*;\
rm -f /lib/systemd/system/anaconda.target.wants/*;

RUN \
systemctl enable sockd

# 避免出现僵尸进程: https://developers.redhat.com/blog/2016/09/13/running-systemd-in-a-non-privileged-container/
STOPSIGNAL SIGRTMIN+3

CMD ["/usr/sbin/init"]
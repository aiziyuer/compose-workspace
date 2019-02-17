FROM centos/systemd

# initial pass
RUN echo 'root:root' | chpasswd

RUN \
rpm --rebuilddb  \
&& yum -y update  \
&& yum clean all

# tools
RUN yum -y install net-tools tcpdump telnet lsof wget curl \
&& yum -y install openssh-server \
&& yum install -y epel-release \
&& yum clean all

# 复写sshd的service文件
COPY config/usr/lib/systemd/system/sshd.service /usr/lib/systemd/system/sshd.service

RUN (cd /lib/systemd/system/sysinit.target.wants/; for i in *; do [ $i == systemd-tmpfiles-setup.service ] || rm -f $i; done); \
rm -f /lib/systemd/system/multi-user.target.wants/*;\
rm -f /etc/systemd/system/*.wants/*;\
rm -f /lib/systemd/system/local-fs.target.wants/*; \
rm -f /lib/systemd/system/sockets.target.wants/*udev*; \
rm -f /lib/systemd/system/sockets.target.wants/*initctl*; \
rm -f /lib/systemd/system/basic.target.wants/*;\
rm -f /lib/systemd/system/anaconda.target.wants/*;


# 设置开机自启动的服务, 这里就以sshd来做测试
RUN \
	systemctl enable sshd

# 避免出现僵尸进程: https://developers.redhat.com/blog/2016/09/13/running-systemd-in-a-non-privileged-container/
STOPSIGNAL SIGRTMIN+3

CMD ["/usr/sbin/init"]
FROM centos:7


# 设置环境全英文避免出错
ENV LC_ALL=en_US.utf8 \
    LANG=en_US.UTF-8 \
    LC_CTYPE=en_US.UTF-8

# initial pass
RUN echo 'root:root' | chpasswd

# 覆盖源
COPY etc/yum.repos.d/*.repo /etc/yum.repos.d/

RUN \
rpm --rebuilddb  \
&& yum clean all

# 必要的定位用工具
RUN yum -y install vim net-tools tcpdump telnet lsof wget curl \
&& yum -y install openssh-server \
&& yum install -y epel-release \
&& yum clean all

# 复写sshd的service文件
COPY usr/lib/systemd/system/sshd.service /usr/lib/systemd/system/sshd.service

# 加入环境变量注入服务
COPY etc/systemd/system/container-environment-inject.service /etc/systemd/system/container-environment-inject.service
COPY usr/local/bin/* /usr/local/bin/

# 为/usr/local/bin下的文件添加执行权限
RUN chmod u+x /usr/local/bin/*

RUN (cd /lib/systemd/system/sysinit.target.wants/; for i in *; do [ $i == systemd-tmpfiles-setup.service ] || rm -f $i; done); \
rm -f /lib/systemd/system/multi-user.target.wants/*;\
rm -f /etc/systemd/system/*.wants/*;\
rm -f /lib/systemd/system/local-fs.target.wants/*; \
rm -f /lib/systemd/system/sockets.target.wants/*udev*; \
rm -f /lib/systemd/system/sockets.target.wants/*initctl*; \
rm -f /lib/systemd/system/basic.target.wants/*;\
rm -f /lib/systemd/system/anaconda.target.wants/*;

# 定制SSH
RUN printf "\n" >>/etc/ssh/sshd_config \
&&  echo "UseDNS no" >>/etc/ssh/sshd_config \
&&  echo "AllowTcpForwarding yes" >>/etc/ssh/sshd_config

# 设置开机自启动的服务, 这里就以sshd来做测试
RUN \
	systemctl enable sshd \
&&	systemctl enable container-environment-inject

# 避免出现僵尸进程: https://developers.redhat.com/blog/2016/09/13/running-systemd-in-a-non-privileged-container/
STOPSIGNAL SIGRTMIN+3

CMD ["/usr/sbin/init"]
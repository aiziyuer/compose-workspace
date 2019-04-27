FROM centos:7


# 设置环境全英文避免出错
ENV LC_ALL=en_US.utf8 \
    LANG=en_US.UTF-8 \
    LC_CTYPE=en_US.UTF-8

# 覆盖源
COPY etc/yum.repos.d/*.repo /etc/yum.repos.d/
COPY root/.pip/pip.conf /root/.pip/

# 增加bash命令补全
RUN  \
rpm --rebuilddb  \
&& yum -y install mlocate bash-completion bash-completion-extras \
&& yum clean all

# 编译工具
RUN rpm --rebuilddb \
&& yum groupinstall -y 'Development Tools' \
&& yum -y install vim net-tools tcpdump telnet lsof wget curl lrzsz \
&& yum -y install python36-pip python36-devel openssl-devel \
&& yum clean all

# 安装ansible
RUN \
pip3 install ansible

RUN \
rpm --rebuilddb  \
&& yum -y install sshpass \
&& yum clean all

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


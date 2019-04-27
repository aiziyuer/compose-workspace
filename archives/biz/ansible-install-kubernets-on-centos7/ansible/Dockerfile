FROM centos:7


# 设置环境全英文避免出错
ENV LC_ALL=en_US.utf8 \
    LANG=en_US.UTF-8 \
    LC_CTYPE=en_US.UTF-8

# 覆盖源
COPY etc/yum.repos.d/*.repo /etc/yum.repos.d/

RUN rpm --rebuilddb \
&& yum install -y epel-release  \
&& yum -y install vim net-tools tcpdump telnet lsof wget curl lrzsz \
&& yum -y install python3.6-pip \
&& yum clean all

# 增加bash命令补全
RUN  \
rpm --rebuilddb  \
&& yum -y install mlocate bash-completion bash-completion-extras \
&& yum clean all

# 安装ansible
RUN \
pip3 install ansible

RUN \
yum -y install sshpass

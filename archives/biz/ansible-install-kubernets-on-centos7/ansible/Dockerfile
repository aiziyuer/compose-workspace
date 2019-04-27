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

FROM centos:7


# initial pass
RUN echo 'root:root' | chpasswd

RUN \
rpm --rebuilddb  \
&& yum clean all

# tools
RUN yum -y install net-tools tcpdump telnet lsof wget curl \
&& yum -y install openssh-server \
&& yum install -y epel-release \
&& yum clean all

# install server
RUN yum -y install openvpn \
&& yum clean all

# 安装必要的软件
RUN yum install -y nginx squid \
&& yum clean all

# 拷贝vpn客户端文件
COPY config/etc/openvpn/client1.conf /etc/openvpn/client1.conf 

RUN (cd /lib/systemd/system/sysinit.target.wants/; for i in *; do [ $i == systemd-tmpfiles-setup.service ] || rm -f $i; done); \
rm -f /lib/systemd/system/multi-user.target.wants/*;\
rm -f /etc/systemd/system/*.wants/*;\
rm -f /lib/systemd/system/local-fs.target.wants/*; \
rm -f /lib/systemd/system/sockets.target.wants/*udev*; \
rm -f /lib/systemd/system/sockets.target.wants/*initctl*; \
rm -f /lib/systemd/system/basic.target.wants/*;\
rm -f /lib/systemd/system/anaconda.target.wants/*;


# starup service
RUN \
    systemctl enable openvpn@client1 \
&&  systemctl enable squid \
&&  systemctl enable nginx \
&&	systemctl enable sshd

CMD ["/usr/sbin/init"]
FROM centos:7


# initial pass
RUN echo 'root:root' | chpasswd

RUN \
rpm --rebuilddb  \
&& yum clean all

# tools
RUN \
yum -y install net-tools tcpdump telnet lsof wget curl\
&& yum clean all

# install server
RUN \
yum -y install openssh-server \
&& yum clean all


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
	systemctl enable sshd

CMD ["/usr/sbin/init"]
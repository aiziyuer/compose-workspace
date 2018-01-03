FROM centos/systemd


# initial pass
RUN echo 'root:root' | chpasswd

RUN \
rpm --rebuilddb  \
&& yum -y update  \
&& yum clean all

# tools
RUN \
yum -y install net-tools tcpdump telnet lsof wget curl\
&& yum clean all

# 解决/etc/init.d等脚本不存在的bug
RUN \
yum install -y initscripts\
&& yum clean all

# install server
RUN \
wget -O /etc/yum.repos.d/jenkins.repo http://pkg.jenkins-ci.org/redhat-stable/jenkins.repo \
&& rpm --import https://jenkins-ci.org/redhat/jenkins-ci.org.key \
&& yum -y install java jenkins \
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
	systemctl enable jenkins \
	&& systemctl enable sshd

CMD ["/usr/sbin/init"]
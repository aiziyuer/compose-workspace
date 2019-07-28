FROM kong:1.0.2-centos


# 设置环境全英文避免出错
ENV LC_ALL=en_US.utf8 \
    LANG=en_US.UTF-8 \
    LC_CTYPE=en_US.UTF-8

# initial pass
RUN echo 'root:root' | chpasswd

# 覆盖源
COPY etc/yum.repos.d/*.repo /etc/yum.repos.d/


RUN \
    yum install -y git vim
   

RUN \
    luarocks install mobdebug 
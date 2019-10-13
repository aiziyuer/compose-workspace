FROM ubuntu:16.04 

LABEL maintainer="aiziyuer <ziyu0123456789@gmail.com>"

LABEL org.label-schema.name="Asuswrt Merlin Build" \
      org.label-schema.description="Asuswrt Merlin 固件交叉编译环境" \
      org.label-schema.url="https://aiziyuer.github.io" \
      org.label-schema.vcs-url="https://github.com/aiziyuer/compose-workspace/tree/master/experimental/asuswrt-merlin.ng/workspace" \
      org.label-schema.vendor="aiziyuer" \
      org.label-schema.schema-version="1.0"

RUN dpkg --add-architecture i386 \
    && apt-get update -y \
    && apt-get install -y \
        sudo net-tools cron e2fsprogs wget vim openssl curl psmisc git \
        heirloom-mailx autoconf automake bison bzip2 bsdtar diffutils \
        sed file flex g++ gawk gcc-multilib gettext gperf groff-base \
        zsh libncurses-dev libexpat1-dev libslang2 libssl-dev libtool \
        libxml-parser-perl make patch perl pkg-config python shtool tar \
        texinfo unzip zlib1g zlib1g-dev intltool autopoint libltdl7-dev \
        lib32z1-dev lib32stdc++6 automake1.11 libelf-dev:i386 libelf1:i386 \
    && apt-get autoremove -y \
    && apt-get autoclean -y \
    && rm -rf /var/lib/apt/lists/* \
    && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone

# 编译后发现需要补充的包
RUN apt-get update -y \
    && apt-get install -y \
        liblzo2-dev uuid-dev uuid-dev build-essential lzma-dev liblzma-dev lzma binutils-dev \
        texinfo dos2unix u-boot-tools device-tree-compiler qemu \
        bc xsltproc \
        patch cmake autogen intltool libglib2.0-dev \
        libc6-i386 lib32stdc++6 lib32z1 libelf1:i386 libncurses5:i386 libc6-dev-i386 \
    && apt-get autoremove -y \
    && apt-get autoclean -y \
    && rm -rf /var/lib/apt/lists/*

# 编译后发现需要补充的包
RUN apt-get update -y \
    && apt-get install -y \
        docbook-xsl libtool-bin \
    && apt-get autoremove -y \
    && apt-get autoclean -y \
    && rm -rf /var/lib/apt/lists/*

# 系统定制, 强制系统使用bash来解释sh脚本
RUN \
    ln -sf bash /bin/sh

# ng的编译中有要求不允许root用户
RUN adduser --gecos '' --ingroup root --disabled-password asuswrt && adduser asuswrt sudo
RUN echo '%sudo ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers
RUN \
       mkdir -p /{media,opt} \
    && mkdir -p /home/asuswrt/am-toolchains \
    && mkdir -p /home/asuswrt/asuswrt-merlin.ng \
    && chown -R asuswrt:root /media /opt \
    && chown -R asuswrt:root /home/asuswrt

USER asuswrt

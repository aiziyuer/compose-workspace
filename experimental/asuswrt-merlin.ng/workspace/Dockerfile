FROM byrnedo/alpine-curl as builder

# 下载异构用的解释器
RUN \
 curl -L -k \
      https://github.com/multiarch/qemu-user-static/releases/download/v4.0.0/qemu-aarch64-static.tar.gz \
 |  tar -xvz -C /usr/bin/ \
 && chmod +x /usr/bin/qemu-aarch64-static

FROM ubuntu:16.04

LABEL maintainer="aiziyuer <ziyu0123456789@gmail.com>"

LABEL org.label-schema.name="Asuswrt Merlin Build" \
      org.label-schema.description="Asuswrt Merlin 固件交叉编译环境" \
      org.label-schema.url="https://aiziyuer.github.io" \
      org.label-schema.vcs-url="https://github.com/aiziyuer/compose-workspace/tree/master/experimental/asuswrt-merlin.ng/workspace" \
      org.label-schema.vendor="aiziyuer" \
      org.label-schema.schema-version="1.0"

# 注册解析执行器, 主要用户异构场景
COPY --from=builder /usr/bin/qemu-aarch64-static /usr/bin/qemu-aarch64-static

# 替换成速度快一些的华为云mirrors
RUN \
    sed -i "s@http://.*archive.ubuntu.com@http://mirrors.huaweicloud.com@g" /etc/apt/sources.list \
 && sed -i "s@http://.*security.ubuntu.com@http://mirrors.huaweicloud.com@g" /etc/apt/sources.list 

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

# 系统定制, 强制系统使用bash来解释sh脚本
RUN \
    ln -sf bash /bin/sh

# ng的编译中有要求不允许root用户
RUN adduser --gecos '' --ingroup root --disabled-password asuswrt && adduser asuswrt sudo
RUN echo '%sudo ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers

# 下载编译工具链, 下面是加速mirrors, 原始链接: https://github.com/RMerl/am-toolchains.git
RUN git clone --depth=1 \
       https://codehub.devcloud.huaweicloud.com/aiziyuer00001/am-toolchains.git /opt/toolchains \
 && chown -R asuswrt:root /opt/toolchains

# 编译后发现需要补充的包
RUN apt-get update -y \
 && apt-get install -y \
        liblzo2-dev uuid-dev uuid-dev build-essential lzma-dev liblzma-dev lzma binutils-dev \
        texinfo dos2unix u-boot-tools device-tree-compiler \
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

RUN \
    echo "export LD_LIBRARY_PATH=$LD_LIBRARY:/opt/toolchains/crosstools-arm-gcc-5.3-linux-4.1-glibc-2.22-binutils-2.25/usr/lib" >> ~/.profile \
 && echo "export TOOLCHAIN_BASE=/opt/toolchains" >> ~/.profile \
 && echo "PATH=\$PATH:/opt/toolchains/crosstools-arm-gcc-5.3-linux-4.1-glibc-2.22-binutils-2.25/usr/bin" >> ~/.profile \
 && echo "PATH=\$PATH:/opt/toolchains/crosstools-aarch64-gcc-5.3-linux-4.1-glibc-2.22-binutils-2.25/usr/bin" >> ~/.profile

RUN \
       mkdir -p /{media,opt} \
 && mkdir -p /home/asuswrt/am-toolchains \
 && mkdir -p /home/asuswrt/asuswrt-merlin.ng

USER asuswrt

FROM openresty/openresty:1.15.8.1-3-alpine

# 源替换
RUN \
    sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories

# 安装必要的工具
RUN \
    apk add --no-cache \
        gd \
        git \
        geoip \
        libgcc \
        libxslt \
        zlib \
        unzip

# 安装luarocks
RUN apk add --no-cache --virtual .build-deps \
        build-base \
        coreutils \
        make \
 && wget http://luarocks.org/releases/luarocks-3.1.3.tar.gz \
 && tar -xzvf luarocks-*.tar.gz \
 && cd luarocks-*/ \
 && ./configure --prefix=/usr/local/openresty/luajit \
    --with-lua=/usr/local/openresty/luajit/ \
    --lua-suffix=jit \
    --with-lua-include=/usr/local/openresty/luajit/include/luajit-2.1 \
 && make \
 && make install \
 && apk del .build-deps

# 安装开发包并安装luarocks
RUN apk add --no-cache --virtual .build-deps \
        build-base \
        coreutils \
        curl \
        gd-dev \
        geoip-dev \
        libxslt-dev \
        linux-headers \
        make \
        perl-dev \
        readline-dev \
        zlib-dev \
 && luarocks install mobdebug \
 && apk del .build-deps



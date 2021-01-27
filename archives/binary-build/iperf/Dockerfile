ARG VERSION

FROM alpine

ENV VERSION=${VERSION:-3.9}

RUN \
apk add build-base pkgconf autoconf automake libtool \
        git perl linux-headers

WORKDIR /workspace

RUN git clone --depth 1 --branch ${VERSION} https://github.com/esnet/iperf.git .

RUN \
./configure --disable-shared --enable-static \
            --prefix=/usr/bin CXXFLAGS="-std=c++14" CPPFLAGS="--static" LDFLAGS="--static"

RUN \
make install

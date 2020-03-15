FROM centos:7

RUN \
    yum group install -y "Development Tools" 

# install cmake
RUN \
   curl -sL https://github.com/Kitware/CMake/releases/download/v3.17.0-rc3/cmake-3.17.0-rc3-Linux-x86_64.tar.gz \
    | tar zx --keep-old-files --strip-components=1 -C /usr

# install llvm
WORKDIR /opt/llvm-src
RUN \
   curl -sL https://releases.llvm.org/6.0.1/llvm-6.0.1.src.tar.xz \
    | tar xJ --strip-components=1 -C .
WORKDIR /opt/llvm
RUN cmake /opt/llvm-src

RUN cmake --build . 

RUN cmake -DCMAKE_INSTALL_PREFIX=/tmp/llvm -P cmake_install.cmake

# install 
# RUN \
#       git clone https://github.com/TunSafe/TunSafe.git \
#    && cd TunSafe

# RUN \
#    sh ./build_linux.shy

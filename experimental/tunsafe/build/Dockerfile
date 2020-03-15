FROM centos:7

# 参考: https://llvm.org/docs/CMake.html
RUN \
    yum group install -y "Development Tools" 

RUN \
    yum install -y \
        python-devel libffi-devel graphviz-devel \
        elfutils-libelf-devel readline-devel libedit-devel \
        libxml2-devel protobuf-devel gtext-devel doxygen swig

# install cmake
RUN \
   curl -sL https://github.com/Kitware/CMake/releases/download/v3.17.0-rc3/cmake-3.17.0-rc3-Linux-x86_64.tar.gz \
    | tar zx --keep-old-files --strip-components=1 -C /usr

# install llvm
# 参考: https://blog.csdn.net/weixin_33782386/article/details/88839625
WORKDIR /opt/llvm-src

ENV LLVM_VERSION=6.0.1

RUN \
   curl -sL https://releases.llvm.org/${LLVM_VERSION}/llvm-${LLVM_VERSION}.src.tar.xz \
    | tar xJ --strip-components=1 -C .
RUN \
   mkdir -p ./tools/clang ./tools/lld ./tools/lldb ./tools/polly ./tools/clang/tools/extra; \
   curl -sL https://releases.llvm.org/${LLVM_VERSION}/cfe-${LLVM_VERSION}.src.tar.xz\
   | tar xJ --strip-components=1 -C ./tools/clang; \
   curl -sL https://releases.llvm.org/${LLVM_VERSION}/lld-${LLVM_VERSION}.src.tar.xz \
   | tar xJ --strip-components=1 -C ./tools/lld; \
   curl -sL https://releases.llvm.org/${LLVM_VERSION}/lldb-${LLVM_VERSION}.src.tar.xz \
   | tar xJ --strip-components=1 -C ./tools/lldb; \
   curl -sL https://releases.llvm.org/${LLVM_VERSION}/polly-${LLVM_VERSION}.src.tar.xz \
   | tar xJ --strip-components=1 -C ./tools/polly; \
   curl -sL https://releases.llvm.org/${LLVM_VERSION}/clang-tools-extra-${LLVM_VERSION}.src.tar.xz \
   | tar xJ --strip-components=1 -C ./tools/clang/tools/extra; \
   mkdir -p ./projects/openmp ./projects/libcxx ./projects/libcxxabi ./projects/libunwind ./projects/compiler-rt; \
   curl -sL https://releases.llvm.org/${LLVM_VERSION}/openmp-${LLVM_VERSION}.src.tar.xz \
   | tar xJ --strip-components=1 -C ./projects/openmp; \
   curl -sL https://releases.llvm.org/${LLVM_VERSION}/libcxx-${LLVM_VERSION}.src.tar.xz \
   | tar xJ --strip-components=1 -C ./projects/libcxx; \
   curl -sL https://releases.llvm.org/${LLVM_VERSION}/libcxxabi-${LLVM_VERSION}.src.tar.xz \
   | tar xJ --strip-components=1 -C ./projects/libcxxabi; \
   curl -sL https://releases.llvm.org/${LLVM_VERSION}/libunwind-${LLVM_VERSION}.src.tar.xz\
   | tar xJ --strip-components=1 -C ./projects/libunwind; \
   curl -sL https://releases.llvm.org/${LLVM_VERSION}/compiler-rt-${LLVM_VERSION}.src.tar.xz \
   | tar xJ --strip-components=1 -C ./projects/compiler-rt

WORKDIR /opt/llvm
RUN cmake /opt/llvm-src
RUN cmake --build . --config Release -- -j 4
RUN cmake -DCMAKE_INSTALL_PREFIX=/usr -P cmake_install.cmake

# RUN yum install -y \
#         libedit-dev ncurses ncurses-devel libxml2 libxml2-devel

# install 
# RUN \
#       git clone https://github.com/TunSafe/TunSafe.git \
#    && cd TunSafe

# RUN \
#    sh ./build_linux.shy

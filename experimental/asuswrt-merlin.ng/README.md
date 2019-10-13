华硕ASUS 86U固件编译
===

## 环境准备

需要一个可以运行`docker`的环境, 并且预装`docker-compose`等必要小工具


## 编译准备

``` bash
# 下载源码
git clone https://github.com/RMerl/am-toolchains.git /opt/am-toolchains
git clone https://github.com/RMerl/asuswrt-merlin.ng /opt/asuswrt-merlin.ng
cd /opt/asuswrt-merlin.ng && git checkout 384.13_1 -b 384.13_1

# 用容器启动编译环境干干净净
docker-compose run workspace bash

# 修改文件数组
sudo chown asuswrt:root -R /home/asuswrt/asuswrt-merlin.ng

# 设置工具链接
ln -s ~/am-toolchains/brcm-arm-hnd /opt/toolchains
echo "export LD_LIBRARY_PATH=$LD_LIBRARY:/opt/toolchains/crosstools-arm-gcc-5.3-linux-4.1-glibc-2.22-binutils-2.25/usr/lib" >> ~/.profile
echo "export TOOLCHAIN_BASE=/opt/toolchains" >> ~/.profile
echo "PATH=\$PATH:/opt/toolchains/crosstools-arm-gcc-5.3-linux-4.1-glibc-2.22-binutils-2.25/usr/bin" >> ~/.profile
echo "PATH=\$PATH:/opt/toolchains/crosstools-aarch64-gcc-5.3-linux-4.1-glibc-2.22-binutils-2.25/usr/bin" >> ~/.profile

# 设置输出目录
mkdir -p /media/ASUSWRT/
ln -s ~/asuswrt-merlin.ng /media/ASUSWRT/asuswrt-merlin.ng

# 开始编译
cd ~/asuswrt-merlin.ng/release/src-rt-5.02hnd
#export CC=/opt/toolchains/crosstools-aarch64-gcc-5.5-linux-4.1-glibc-2.26-binutils-2.28.1/bin/aarch64-linux-gcc
make rt-ac86u


```
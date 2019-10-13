华硕ASUS 86U固件编译
===

## 环境准备

需要一个可以运行`docker`的环境, 并且预装`docker-compose`等必要小工具


## 编译准备

``` bash
# 下载源码
git clone --depth=1 -b 384.13_1 \
    https://github.com/aiziyuer/asuswrt-merlin.ng.git /media/asuswrt-merlin.ng

# 宿主机注册aarch64的cpu异构钩子
docker run --rm --privileged multiarch/qemu-user-static:register
# 取消注册钩子
# docker run --rm --privileged multiarch/qemu-user-static:register --reset

# 用容器启动编译环境干干净净
docker-compose run workspace bash

# 修改文件数组
sudo chown asuswrt:root -R /media/ASUSWRT/asuswrt-merlin.ng

# 开始编译
cd /media/ASUSWRT/asuswrt-merlin.ng/release/src-rt-5.02hnd
make rt-ac86u | tee output.log

# 这一步非常久, 我这里大概是3个小时

```

## 系统定制

### IPV6的NAT开关打开

``` bash
# 默认官方的梅林固件的ipv6是没有打开NAT功能的, 重新编译来支持
# 文件路径: release/src-rt-5.02hnd/hostTools/scripts/defconfig-bcm.template
# CONFIG_NF_NAT_MASQUERADE_IPV6
# CONFIG_IP6_NF_NAT
# CONFIG_IP6_NF_TARGET_MASQUERADE
# 以上几个开关设置为y

# 成功后可以在内核模块中看到nf_nat_ipv6.ko文件

```


## 蛋疼网络卡住了镜像构建

``` bash
# 切换到需要编译的Dockerfile所在目录, 配置一下代理让构建过程飞起来
docker build --build-arg http_proxy=http://192.168.50.136:3128 --build-arg https_proxy=http://192.168.50.136:3128 .
```

## FAQ

- [mritd/asuswrt-merlin-build](https://hub.docker.com/r/mritd/asuswrt-merlin-build/dockerfile)
- [Compile Firmware from source using Ubuntu](https://github.com/RMerl/asuswrt-merlin/wiki/Compile-Firmware-from-source-using-Ubuntu)
功能说明
===

制作cpu异构的容器


### 测试功能

``` bash

# 注册宿主机操作系统级解释器
docker run --rm --privileged multiarch/qemu-user-static:register --reset
docker run --rm --privileged multiarch/qemu-user-static:register

curl -L -ko /usr/bin/qemu-aarch64-static \
https://github.com/multiarch/qemu-user-static/releases/download/v4.0.0/qemu-aarch64-static

```
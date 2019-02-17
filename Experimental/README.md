实验性质
===

这个目录中的所有制作的镜像过程都是实验性质(不可用于生产), 里面可能是一些容器镜像的一些探索, 也可能是一些生产镜像制作的一些前期试水记录



## 命令合集

``` bash

# 编译某个应用的镜像
docker-compose build centos7-systemd-environment

# 编译某个应用的镜像-不缓存
docker-compose build  --no-cache centos7-systemd-environment

# 启动某个应用
docker-compose up -d centos7-systemd-environment

# 进入某个应用容器
docker-compose exec centos7-systemd-environment bash

```
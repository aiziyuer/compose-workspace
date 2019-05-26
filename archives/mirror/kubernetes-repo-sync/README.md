功能说明
===

kubernetes的官方源同步


### 测试功能

```
# 启动应用
docker-compose build && docker-compose up -d 

# 进入容器
docker-compose exec slave bash

# 设置代理(如果有需要的话)
export http_proxy=http://192.168.224.1:3128
export https_proxy=http://192.168.224.1:3128

# 从互联网上同步数据
reposync --plugins --repoid=kubernetes \
--allow-path-traversal \
--download_path=/mirrors/tmp/kubernetes

# 同步到目的目录
mkdir -p /mirrors/kubernetes
rsync -avzP --delete /mirrors/tmp/pool /mirrors/kubernetes/


# 重新生成索引
craterepo -v /mirrors/kubernetes/
```
简单http网关
===


``` bash

# 安装初始化的工具
yum install -y redis unix2dos


# 初始化redis的
cat <<EOL | unix2dos | redis-cli --pipe
FLUSHALL
SET backend:example.com:80 11.11.11.101:80
EOL

# 样例nginx服务从gateway访问
curl -X GET \
--url http://localhost \
-H "Host:example.com"

# 原始的样例nginx服务直接访问
curl -X GET \
--url http://11.11.11.101

```


## FAQ

-   [Dynamic Routing Based On Redis](https://openresty.org/en/dynamic-routing-based-on-redis.html)
-   [Nginx内置绑定变量](https://wiki.jikexueyuan.com/project/openresty/openresty/inline_var.html)
-   [Redis操作指南](https://redis.io/topics/mass-insert)
-   [lua与redis结合应用于nginx的动态upstream](http://www.rendoumi.com/luayu-redisjie-he-ying-yong-yu-nginxde-dong-tai-upstream/)
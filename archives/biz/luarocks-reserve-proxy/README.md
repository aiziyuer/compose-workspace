简单反向代理
===

本代理将支持代理任意的网站

``` bash

# 启动应用
docker-compose down -v --remove-orphans; docker-compose up -d; docker-compose ps;

# 安装初始化的工具
yum install -y redis unix2dos

# 初始化redis的
cat <<EOL | unix2dos | redis-cli --pipe
FLUSHALL
SET backend:example.com:80 "202.38.95.110:80 mirrors.ustc.edu.cn"
EOL

# 样例nginx服务从gateway访问
curl -X GET \
--url http://localhost/alpine/v3.10/releases/x86_64/alpine-extended-3.10.0-x86_64.iso.sha256 \
-H "Host:example.com" 

# 原始的样例nginx服务直接访问
curl http://mirrors.ustc.edu.cn/alpine/v3.10/releases/x86_64/alpine-extended-3.10.0-x86_64.iso.sha256

```


## FAQ

-   [OpenResty最佳实践-正则表达式](https://moonbingbing.gitbooks.io/openresty-best-practices/lua/re.html)
-   [Dynamic Routing Based On Redis](https://openresty.org/en/dynamic-routing-based-on-redis.html)
-   [Nginx内置绑定变量](https://wiki.jikexueyuan.com/project/openresty/openresty/inline_var.html)
-   [Redis操作指南](https://redis.io/topics/mass-insert)
-   [lua与redis结合应用于nginx的动态upstream](http://www.rendoumi.com/luayu-redisjie-he-ying-yong-yu-nginxde-dong-tai-upstream/)
-   [Environment Variables in Nginx Config](https://blog.doismellburning.co.uk/environment-variables-in-nginx-config/)
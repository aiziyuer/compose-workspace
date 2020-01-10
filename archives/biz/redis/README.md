简易Redis环境
===


### 预置数据

``` bash

# 初始化redis的
cat <<EOL | unix2dos | redis-cli --pipe
FLUSHALL
SET backend:example.com:80 "202.38.95.110:80 mirrors.ustc.edu.cn"
EOL

```

### 备份还原

``` bash

# 启动redis
docker-compose up -d redis

# 预置数据
cat <<EOL | docker run --rm mor1/dos2unix unix2dos | docker run --rm redis  redis-cli -h 172.17.0.1 --pipe
FLUSHALL
SET backend:example.com:80 "202.38.95.110:80 mirrors.ustc.edu.cn"
EOL

# 备份
docker run --rm aiziyuer/redis-dump:0.4.0 redis-dump -u 172.17.0.1:6379 > db_full.json

# 还原
cat db_full.json | docker run --rm aiziyuer/redis-dump:0.4.0 redis-load -u 172.17.0.1:6379 > db_full.json 


```

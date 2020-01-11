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



```

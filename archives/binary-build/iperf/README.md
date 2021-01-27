编译说明
---

```
# 编译
docker build . -t iperf:build

# 取出二进制
docker run --rm iperf:build -v $PWD:$PWD cp /usr/bin/iperf3 $PWD/

# 检查产物
ldd $PWD/iperf3
```
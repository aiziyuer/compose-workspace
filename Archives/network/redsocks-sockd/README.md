功能说明
===

ss-redir的精简版本


### 测试功能

```
# 服务器端添加一张别名网卡, 网段假设是31.31.31.31/32
ifconfig lo:0 31.31.31.31/32

# 用nc模拟tcp服务器
yum install -y nc pv 
nc -4 -l -k -p 4444 </dev/zero

# 客户端用nc+pv来测试速度
yum install -y nc pv
nc 31.31.31.31 4444 | pv >/dev/null

```


### FAQ

- [Configuring Dante Socks 101](https://www.pahoehoe.net/tag/sockd-service/)
- [dante官网](https://www.inet.no/dante/)
- [dante搭建socks5代理](https://lixingcong.github.io/2018/05/25/dante-socks5/)
- [README-setup-firewall-on-systemd.md](https://gist.github.com/drmalex07/7712d4185b7651747932)
- [Systemd 入门教程：实战篇](http://www.ruanyifeng.com/blog/2016/03/systemd-tutorial-part-two.html)
- [systemd的target依赖关系](https://www.freedesktop.org/software/systemd/man/bootup.html#System%20Manager%20Bootup)
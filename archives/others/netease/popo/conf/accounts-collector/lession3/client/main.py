# -*- coding:utf8 -*-

import socket
from lession2 import parse_member_info

# 参考文档:
# https://python3-cookbook.readthedocs.io/zh_CN/latest/c11/p02_creating_tcp_server.html

# 拦截服务器
s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.connect(('localhost', 20000))

# 告知服务器需要这个时间点统计的人员信息
s.sendall('20190519')
data = s.recv(1024).strip()

# 处理服务器端发过来的人员清单信息
parse_member_info(data)

s.close()

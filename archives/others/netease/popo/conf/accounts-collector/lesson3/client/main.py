# -*- coding:utf8 -*-

import socket


# 参考文档:
# https://python3-cookbook.readthedocs.io/zh_CN/latest/c11/p02_creating_tcp_server.html

s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.connect(('localhost', 20000))

s.sendall('20190519')
data = s.recv(1024).strip()

print data

s.close()

# -*- coding:utf8 -*-

import socket

client_addr = ('localhost', 65432)
sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
sock.connect(client_addr)

while True:
    inp = input("please input data:")
    sock.sendall(str(inp))

sock.close()

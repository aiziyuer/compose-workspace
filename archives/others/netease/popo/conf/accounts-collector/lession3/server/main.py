# -*- coding:utf8 -*-

from SocketServer import StreamRequestHandler, ThreadingTCPServer
from lession1 import parse_raw_data_by_date


# 参考文档:
# https://python3-cookbook.readthedocs.io/zh_CN/latest/c11/p02_creating_tcp_server.html

class MyHandler(StreamRequestHandler):
    def handle(self):
        parameter_date = self.request.recv(1024).strip()
        data = parse_raw_data_by_date(parameter_date)
        print data
        self.request.sendall(data)


if __name__ == '__main__':
    serv = ThreadingTCPServer(('0.0.0.0', 20000), MyHandler)
    serv.serve_forever()

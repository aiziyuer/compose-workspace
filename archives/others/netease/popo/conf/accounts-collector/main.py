# -*- coding:utf8 -*-

import Queue
import os
import re

_CUR_DIR = 'data'

# 这里存储所有满足条件的文件
py_files = []

# 模拟递归查找文件
q = Queue.Queue()
q.put('data')
while not q.empty():
    dir_name = q.get()
    fileList = os.listdir(dir_name)
    for f in fileList:
        f = os.path.join(dir_name, f)
        if os.path.isdir(f):
            q.put(f)
        else:
            if f.endswith('.py'):
                py_files.append(os.path.normpath(f))

# 开始处理文件内容找出需要的个人信息
personal_info_regex = re.compile(r"(?:name\s+=\s+u[\\]*')([^'\\]+)[\\]*'\s+(?:popo\s+=\s+u[\\]*')([^'\\]+)[\\]*'")
pre_department = ''
for f in py_files:
    department = os.path.dirname(f)
    if pre_department != department:
        pre_department = department
        print os.path.basename(pre_department)
    with open(f, 'r') as content_file:
        content = content_file.read()
        m = personal_info_regex.search(content)
        if m:
            print m.group(1)
            print m.group(2)

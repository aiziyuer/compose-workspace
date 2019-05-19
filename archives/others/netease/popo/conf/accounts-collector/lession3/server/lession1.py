# -*- coding:utf8 -*-

import Queue
import os
import re
import StringIO


def parse_raw_data_by_date(date='20190519'):
    _DATA_DIR = 'input' + '/' + date

    # 这里存储所有满足条件的文件
    py_files = []

    # 模拟递归查找文件
    q = Queue.Queue()
    q.put(_DATA_DIR)
    while not q.empty():
        dir_name = q.get()
        file_list = os.listdir(dir_name)
        for f in file_list:
            f = os.path.join(dir_name, f)
            if os.path.isdir(f):
                q.put(f)
            else:
                if f.endswith('.py'):
                    py_files.append(os.path.normpath(f))

    # 开始处理文件内容找出需要的个人信息
    personal_info_regex = re.compile(r"(?:name\s+=\s+u[\\]*')([^'\\]+)[\\]*'\s+(?:popo\s+=\s+u[\\]*')([^'\\]+)[\\]*'")
    pre_department = ''

    output = StringIO.StringIO()
    for f in py_files:
        department = os.path.dirname(f)
        if pre_department != department:
            pre_department = department
            # print os.path.basename(pre_department)
            output.write(os.path.basename(pre_department))
            output.write('\n')
        with open(f, 'r') as content_file:
            content = content_file.read()
            m = personal_info_regex.search(content)
            if m:
                # print m.group(1)
                output.write(m.group(1))
                output.write('\n')
                # print m.group(2)
                output.write(m.group(2))
                output.write('\n')
    contents = output.getvalue()
    output.close()

    return contents


if __name__ == '__main__':
    _OUTPUT_DIR = 'output'
    data = parse_raw_data_by_date()
    fb = open(_OUTPUT_DIR + '/name.txt', 'w')
    fb.write(data)

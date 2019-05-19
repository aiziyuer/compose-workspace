# -*- coding:utf8 -*-

import xlwt

_INPUT_DIR = 'input'
_OUTPUT_DIR = 'output'


# 人员信息

class Person:

    def __init__(self, user_depart, user_name, popo_email):
        self.user_depart = user_depart
        self.user_name = user_name
        self.popo_email = popo_email


file_path = _INPUT_DIR + '/name.txt'
one_list = []
with open(file_path) as fp:
    line = fp.readline().rstrip('\n')
    cnt = -1

    # 临时计数
    nobody = Person('', '', '')

    while line:

        # 判断是否是部门信息, 如果是部门, 行号计数清零, 否则行号最后累加
        if line.istitle():
            cnt = 0
            nobody.user_depart = line
        # 判断是否是姓名, 累计行号模2余1
        elif cnt % 2 == 1:
            nobody.user_name = line
        # 判断是否是邮箱, 累计行号模2余0, 封装结构
        elif cnt % 2 == 0:
            nobody.popo_email = line
            someone = Person(nobody.user_depart, nobody.user_name, nobody.popo_email)
            one_list.append(someone)

        # 固定循环
        line = fp.readline().rstrip('\n')
        cnt += 1

# 3. 人员列表分部门存储
depart_dict = {}
for one in one_list:
    team = depart_dict.get(one.user_depart)
    team = team if team is not None else []
    depart_dict[one.user_depart] = team
    team.append(one)

# 4. 输出为excel
wbk = xlwt.Workbook(encoding='utf-8')
sheet = wbk.add_sheet(u'Sheet1', cell_overwrite_ok=True)
team_count = 0
for team_name, team in depart_dict.items():

    team_count += 1

    # 录入部门信息
    line_no = 0
    sheet.write(line_no, (team_count - 1) * 2, team_name,
                style=xlwt.easyxf('font: color-index red, bold on,height 280, name SimSun'))
    line_no += 1

    for one in team:
        sheet.write(line_no, (team_count - 1) * 2, one.user_name,
                    style=xlwt.easyxf('pattern: pattern solid, fore_colour yellow;'
                                      'font: italic 1, color-index red, bold on, height 220, name Microsoft YaHei'))
        sheet.write(line_no, (team_count - 1) * 2 + 1, one.popo_email,
                    style=xlwt.easyxf('font: italic 1, color-index green, bold off, height 220, name SimSun'))
        line_no += 1

        pass

    pass

wbk.save(_OUTPUT_DIR + '/name.xls')

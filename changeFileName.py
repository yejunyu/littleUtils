# -*- coding: utf-8 -*-
# @author yejunyu
# @date 18-7-20.

import os
import re

"""
有的电影或者下载的资源都带着很多前缀后缀,去除掉
"""

path = "/Users/yejunyu/Downloads/222937210_神级第六人/MOOC/golang/基于Golang协程实现流量统计系统"
pattern = re.compile("(【.*?】)")
for _, dirs, files in os.walk(path):
    dir_list = dirs
    file_list = files
    break
files = dir_list + file_list
print(files)
for file in files:
    new_file_name = pattern.sub(r"", file)
    os.rename(path + "/" + file, path + "/" + new_file_name)

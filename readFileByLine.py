#-*- coding: utf-8 -*- 
# @author yejunyu
# @date 18-7-20.

import time
import sys

def read_file_by_line(rsc,dst):
    """
    逐行读文件,用于日志传递或者kafka测试
    :param rsc:
    :param dst:
    :return:
    """
    with open(rsc,"r") as r:
        w = open(dst,'a')
        for line in r:
            w.write(line)
            print(line)
            time.sleep(1)
        w.close()
    r.close()

if __name__ == '__main__':
    read_file_by_line(sys.argv[1],sys.argv[2])

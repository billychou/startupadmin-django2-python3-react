#!/usr/bin/env python
# -*- coding: utf-8 -*-
# 
# Copyright (c) 2018 All Rights Reserved
# 

"""
File: __init__.py
Author: songchuan.zhou
Date: 2018/8/24 22
"""


def insertion_sort(array):
    # 插入排序算法
    for j in range(2, len(array)+1):
        key = array[j]
        i = j - 1
        while i>0 and array[i]>key:
            array[i+1] = array[i]
            i = i -1
        array[i+1] = key

if __name__ == "__main__":
    print(insertion_sort([5, 3, 4, 6, 1]))

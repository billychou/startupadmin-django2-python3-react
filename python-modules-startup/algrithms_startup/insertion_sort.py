#!/usr/bin/env python
# -*- coding: utf-8 -*-
# 
# Copyright (c) 2018 All Rights Reserved
# 

"""
File: insertion_sort.py
Author: songchuan.zhou
Date: 2018/8/24 22
"""


def insertion_sort(array):
    # 插入排序算法
    # 从第二个key 进行排序
    for j in range(1, len(array)):
        # 循环 key
        key = array[j]
        i = j - 1
        while i >= 0 and array[i] > key:
            array[i+1] = array[i]
            array[i] = key
            i -= 1

if __name__ == "__main__":
    array = [5, 3, 4, 6, 1]
    insertion_sort(array)
    print(array)
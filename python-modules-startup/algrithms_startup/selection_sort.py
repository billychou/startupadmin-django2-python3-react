#!/usr/bin/env python
# -*- coding: utf-8 -*-
# 
# Copyright (c) 2018 All Rights Reserved
# 

"""
File: findMin.py
Author: songchuan.zhou
Date: 2018/8/25 09
"""


def find_min(a):
    # find the minest value
    min = a[0]
    for i in range(1, len(a)):
        if (a[i] < min):
            min = a[i]
    return min


def selection_sort(a):
    for j in range(0, len(a) - 1):
        min = i
        # min_value = min(a[i:])
        # if a[j] > min_value:
        #     a[0] = min_value
        for i in range(j+1, len(a)):


    return a

if __name__ == "__main__":
    a = [6, 8, 10, 3, 5]
    selection_sort(a)
    print(a)

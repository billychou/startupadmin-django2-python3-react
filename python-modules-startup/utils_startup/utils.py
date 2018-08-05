#!/usr/bin/env python
# -*- coding: utf-8 -*-
# 
# Copyright (c) 2018 All Rights Reserved
# 

"""
File: utils
Author: songchuan.zhou
Date: 2018/8/2 23
"""


def more(text, numlines=15):
    lines = text.splitlines()  # 效果类似split('\n'), 不过不用在末尾加''
    while lines:
        chunk = lines[:numlines]

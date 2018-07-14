#!/usr/bin/env python
# -*- coding: utf-8 -*-
# 
# Copyright (c) 2018. All Rights Reserved
# 

"""
File: convert.py
Author: songchuan.zhou (songchuan.zhou)
Date: 2018/7/11 14-15
Usage: 提供常用的类型转换
"""

import itertools
import socket
import struct


def unicode2utf8(_value):
    """
    unicode 转化为 utf8
    :param _value:
    :return:
    """
    if isinstance(_value, unicode):
        _value = _value.encode('utf-8')
    else:
        _value = str(_value)
    return _value


def utf82unicode(content):
    if isinstance(content, utf8):
        pass
    else:
        pass






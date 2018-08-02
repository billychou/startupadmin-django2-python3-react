#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2018 alibaba-inc. All Rights Reserved
# 
########################################################################
 
"""
File: socker_demo_client.py
Author: songchuan.zhou(songchuan.zhou@alibaba-inc.com)
Date: 2018/08/02 09:27:24
"""
import socket
s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.connect(('127.0.0.1', 8080))
data = b'welsklskssjkssss'
s.send(data)
data_get = s.recv(1024)    # 进程会卡在这里，收不到响应
print(data_get)
s.close()

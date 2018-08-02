#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2018 alibaba-inc. All Rights Reserved
# 
########################################################################
 
"""
File: socket_demo_server.py
Author: songchuan.zhou(songchuan.zhou@alibaba-inc.com)
Date: 2018/08/02 09:25:12
"""
import socket
server = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
server.bind(('127.0.0.1', 8080))
server.listen(50)


while True:
    data, addr = server.accept()
    info = b'welcome'
    while True:
        buffer = data.recv(1024)    # 如果客户端发送 b''，进程将卡在这里，下面无法执行
        print(buffer)
        print(buffer==b'')
        if buffer == b'':
            info = b'Bad request'
            break
    data.send(info)
    data.close()

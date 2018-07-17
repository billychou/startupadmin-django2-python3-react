#!/usr/bin/env python
# -*- coding: utf-8 -*-
# 
# Copyright (c) 2018 All Rights Reserved
# 

"""
File: loghandler.py
Author: songchuan.zhou
Date: 2018/7/15 00
"""

from libs.common.mixin.loghandler import StartupLogFileHandlerMixin


class StartupLogFileHandler(StartupLogFileHandlerMixin):
    """
    进程日志处理
    """
    def __init__(self, filename, when="h", interval=1, backupCount=20, encoding=None, delay=False, utc=False):
        self.delay = delay
        super(StartupLogFileHandler, self).__init__(filename, when, interval, backupCount, encoding, delay, utc)

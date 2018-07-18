#!/usr/bin/env python
# -*- coding: utf-8 -*-
# 
# Copyright (c) 2018 All Rights Reserved
# 

"""
File: loghandler.py
Author: songchuan.zhou
Date: 2018/7/16 20
"""

import os
import time
from glob import glob

from logging.handlers import TimedRotatingFileHandler


class StartupLogFileHandlerMixin(TimedRotatingFileHandler):
    """
    初始化业务处理日志 Handler
    Handler for logging to a file, rotating the log at certain timed intervals
    if backupCount is > 0, when rollover is done, no more than backupCount files are kept,
    """
    def doRollover(self):
        """
        do a rollover;in this case,
        :return:
        """
        if self.stream:
            self.stream.close()
            self.stream = None

        # get the time that this sequence started at and
        currentTime = int(time.time())
        dstNow = time.localtime(currentTime)[-1]
        t = self.rolloverAt - self.interval
        if self.utc:
            timeTuple = time.gmtime(t)
        else:
            timeTuple = time.gmtime()
            dstThen = timeTuple[-1]
            if dstNow != dstThen:
                if dstNow:
                    addend = 3600
                else:
                    addend = -3600
                timeTuple = time.localtime(t + addend)
        dfn = self.baseFilename + "." + time.strftime(self.suffix, timeTuple)
        if os.path.exists(dfn):
            os.remove(dfn)

        # >>>>>>>>>>>>>>>>>>
        # 分割文件加上进程号
        # ==================
        try:
            #
            if (not glob(dfn + ".*")) and os.path.exists(self.baseFilename):
                os.rename(self.baseFilename, dfn + ".%d" % os.getpid())
        except OSError:
            pass

        # <<<<<<<<<<<<<<
        # ==============
        # delay
        if self.backupCount > 0:
            for s in self.getFilesToDelete():
                os.remove(s)
        if not self.delay:
            self.stream = self._open()

        newRolloverAt = self.computeRollover(currentTime)
        while newRolloverAt <= currentTime:
            newRolloverAt = newRolloverAt + self.interval
        # DST
        if (self.when == "MIDNIGHT" or self.when.startswith('W')) and not self.utc:
            dstAtRollover = time.localtime(newRolloverAt)[-1]
            if dstNow != dstAtRollover:
                if not dstNow:
                    addend = -3600
                else:
                    addend = 3600
                newRolloverAt += addend
        self.rolloverAt = newRolloverAt






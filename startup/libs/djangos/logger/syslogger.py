#!/usr/bin/env python
# -*- coding: utf-8 -*-
# 
# Copyright (c) 2018 All Rights Reserved
# 

"""
File: syslogger.py
Author: songchuan.zhou
Date: 2018/7/15 00
"""

import traceback
import logging
import logging.config
import sys

from django.conf import settings


class SysLogger(object):
    """
    system logger
    """
    INFO_LOGGER = logging.getLogger(settings.PROJECT_INFO_LOG)
    ERROR_LOGGER = logging.getLogger(settings.PROJECT_ERROR_LOG)
    EXCEPTION_LOGGER = logging.getLogger(settings.PROJECT_EXCEPTION_LOG)

    @classmethod
    def debug(cls, msg):
        """
        logging debug message
        :param msg:
        :return:
        """
        extra = {
            "realLocation": repr(traceback.format_stack(limit=2)[0])
        }
        # repr函数将对象转化为可读的形式
        cls.INFO_LOGGER.debug(msg, extra=extra)

    @classmethod
    def info(cls, msg):
        """
        logging info message,
        :param msg:
        :return:
        """
        cls.INFO_LOGGER.info(msg)

    @classmethod
    def warn(cls, msg):
        """
        logging warn message
        :param msg:
        :return:
        """
        extra = {
            "realLocation": repr(traceback.format_stack(limit=2)[0]),
        }
        cls.INFO_LOGGER.warn(msg, extra=extra)

    @classmethod
    def error(cls, msg):
        """
        logging error message
        :param msg:
        :return:
        """
        extra = {
            "realLocation": repr(traceback.format_stack(limit=2)[0])
        }
        cls.INFO_LOGGER.error(msg, extra=extra)

    @classmethod
    def exception(cls, exp, request=None):
        """
        loggig exception message
        :param msg:
        :return:
        """
        extra = {
            "realLocation": repr(traceback.format_stack(limit=2)[0]),
            "request": request
        }
        #  format_stack A shorthand for format_list(extra)
        #  A shorthand for format_list(extract_stack(f, limit)).
        #  format_list
        cls.INFO_LOGGER.error(exp, extra=extra)
        if sys.version_info >= (2, 7, 7):
            cls.EXCEPTION_LOGGER.exception(exp, extra=extra)
        else:
            cls.EXCEPTION_LOGGER.exception(exp)
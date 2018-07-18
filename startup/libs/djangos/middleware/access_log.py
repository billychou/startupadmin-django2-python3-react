#!/usr/bin/env python
# -*- coding: utf-8 -*-
# 
# Copyright (c) 2018 All Rights Reserved
# 

"""
File: access_log.py
Author: songchuan.zhou
Date: 2018/7/17 20
Usage: 访问日志中间件,记录的日志是解密组装处理后的,即使 POST 请求参数也在日志内
"""

import socket
import time
import re
import threading
import logging


from django.conf import settings

from ..logger import SysLogger
from ..misc import get_clientip

ACCESS_LOG = logging.getLogger(settings.PROJECT_ACCESS_LOG)


class InitialParams(threading.local):
    """
    初始化线程变量
    """
    eagleeye_trace_id = None
    eagleeye_rpc_id = None
    eagleeye_user_data = None
    call_counter = 0

INITIAL_PARAMS = InitialParams()


class AccessLogMiddleware(object):
    """
    记录请求日志,和返回状态码
    """
    local_ip = socket.gethostbyname(socket.gethostname())

    #  AttributeError: 'NoneType' object has no attribute '_closable_objects'
    #
    def __init__(self, get_response):
        self.get_response = get_response
        self.start_time = 0
        self.end_time = 0

    def __call__(self, request):
        self.process_request(request)
        response = self.get_response(request)
        self.process_response(request, response)
        return response

    def process_request(self, request):
        self.start_time = time.time()

    def process_response(self, request, response):
        # request
        host = get_clientip(request)
        method = request.method
        scheme = request.scheme
        path = request.get_full_path()
        # response
        status_code = response.status_code
        content_length = len(response.content)
        self.end_time = time.time()
        time_delta = self.end_time - self.start_time
        time_delta = round(time_delta*1000)
        user_agent = request.META.get("HTTP_USER_AGENT", "")
        message = '{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}'.format(
            host,
            method,
            scheme,
            path,
            status_code,
            time_delta,
            user_agent,
            content_length,
            self.local_ip
        )
        ACCESS_LOG.info(message)
        return response

    def eagleeye_header(self, request):
        pass








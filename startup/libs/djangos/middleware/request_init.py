#!/usr/bin/env python
# -*- coding: utf-8 -*-
# 
# Copyright (c) 2018 All Rights Reserved
# 

"""
File: request_init.py
Author: songchuan.zhou
Date: 2018/7/13 15
Usage: 请求初始化预处理
    1. 参数统一整合到 paraments
"""

import urllib
import traceback
from django.http import HttpResponseBadRequest, QueryDict


class RequestInitMiddleware(object):
    """
        请求初始化预处理
        1. 参数合并
    """
    def __init__(self, get_response):
        """
        webserver 启动调用一次
        :param get_response:
        """
        self.get_response = get_response

    def __call__(self, request):
        """
        每一条请求都会调用一次
        :param request:
        :return:
        """
        self.process_request(request)
        response = self.get_response(request)
        return response

    def process_request(self, request):
        try:
            # request 对象增加parameters 属性
            # GET.POST 整合到 parameters
            request.parameters = request.GET.copy()
            raw_body = ''
            if request.method == "POST":
                raw_body = request.body
                request.parameters = request.GET.copy()
                print("raw_body", raw_body)
                if request.META['CONTENT_TYPE'].startswith('multipart/form-data'):
                    raw_body = ''
                # startswith first arg must be bytes or a tuple of bytes, not str
                elif raw_body.startswith(b'<') or raw_body.startswith(b'{') or raw_body.startswith(b'['):
                    raw_body = ''
                elif b'=' in raw_body:
                    request.parameters.update(QueryDict(raw_body))
                    raw_body = ''

                if not request.parameters.get('ent') and not request.parameters.get('in'):
                    for k in request.POST:
                        request.parameters.setlist(k, request.POST.getlist())
            return None

            # 2. 请求参数统一解密
            # ent = request.parameters.get('ent')
            # in_ = request.parameters.get('in')
            # q = request.parameters.get('q')
            # cipher = in_ or q
            # if ent or cipher:

        except Exception as e:
            print(e)
            response = HttpResponseBadRequest()
            return response

    def process_view(self, request, view_func, view_args, view_kwargs):
        pass















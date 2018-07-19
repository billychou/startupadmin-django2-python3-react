#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
#
# Copyright (c) 2018. All Rights Reserved
#
########################################################################

import json
import decimal
import time


class ApiStatus(object):
    """
    接口成功与否的状态枚举，与日常说的http状态码不是一个概念
    """
    SUCCESS = "success"
    FAILED = "failed"


class DecimalEncoder(json.JSONEncoder):
    """
    decimal型浮点数的序列化
    """
    def default(self, obj):
        if isinstance(obj, decimal.Decimal):
            return float(obj)
        return super(DecimalEncoder, self).default(obj)


class ResponseBuilder(object):
    """
    构造响应内容
    """
    def __init__(self):
        pass

    @classmethod
    def response_json(self, context, ensure_ascii=True, indent=0):
        """
        构造请求
        :param context:
        :param ensure_ascii:
                If ``ensure_ascii`` is false, then the return value can contain non-ASCII
                characters if they appear in strings contained in ``obj``. Otherwise, all
                such characters are escaped in JSON strings.
        :param indent:
                If ``indent`` is a non-negative integer, then JSON array elements and
                object members will be pretty-printed with that indent level. An indent
                level of 0 will only insert newlines. ``None`` is the most compact
                representation.
        :return:
        """
        return json.dumps(context, ensure_ascii=ensure_ascii, indent=indent)

    def __call__(self, context=None, status_code=None, code=None, msg=None, version='',
                 result=None, add_response=False):
        if not context:
            context = {}

        if result is None:
            result = True if status_code == 200 else False
        else:
            result = True if code == 200 else False

        response = {
            "version": version,
            "code": code if code is not None else 100,
            "message": msg,
            "timestamp": int(time.time()),
            "result": result,
        }

        if add_response:
            response.update({"data": context})
        else:
            pass
        return response



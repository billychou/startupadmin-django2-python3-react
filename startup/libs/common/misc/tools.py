#!/usr/bin/env python
# -*- coding: utf-8 -*-
# 
# Copyright (c) 2018 All Rights Reserved
# 

"""
File: tools.py
Author: songchuan.zhou
Date: 2018/7/31 19
"""

import json
import time

from ..error.statuscode import CommonStatus


class ResponseBuilder(object):

    def response_json(self, context, ensure_ascii=True, indent=0):
        return json.dumps(context, ensure_ascii=ensure_ascii, indent=indent)

    def __call__(self, context=None, statuscode=None, code=None, msg=None, version='', result=None,
                 add_response=False):
        if not context:
            context = {}

        if result is None:
            if statuscode is not None:
                result = True if statuscode == CommonStatus.SUCCESS else False
            else:
                result = True if code == CommonStatus.SUCCESS.code else False

        response = {
            "version": version,
            "code": code if code is not None else statusc-----------------ode.code,
            "message": msg or statuscode.msg,
            "timestamp": int(time.time()),
            "result": result
        }
        if add_response:
            response.update({"response": context})
        else:
            response.update(context)
        return response
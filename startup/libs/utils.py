#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
#
# Copyright (c) 2018. All Rights Reserved
#
########################################################################

import json
import decimal


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


print(json.dumps({"key": 123, "skks": 234.45}))
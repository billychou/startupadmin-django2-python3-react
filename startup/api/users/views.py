#!/usr/bin/env python
# -*- coding: utf-8 -*-
# 
# Copyright (c) 2018 All Rights Reserved
# 

"""
File: views.py
Author: songchuan.zhou
Date: 2018/7/24 09
"""

import logging
import traceback

from django.views import View
from django.http import HttpResponse
from django.conf import settings

from libs.utils import ResponseBuilder


class UserCheckLoginView(View):
    """
    检查用户是否登录
    """
    def get(self, request, *args):
        result = {
            "message": "test",
            "data": [

            ]
        }
        response = ResponseBuilder.response_json(result)
        return HttpResponse(response)
#!/usr/bin/env python
# -*- coding: utf-8  -*-

from datetime import datetime
from functools import wraps
from fnmatch import fnmatch

from django.conf import settings
from django.http import HttpResponseForbidden, HttpResponseNotAllowed
from django.utils.functional import SimpleLazyObject
from django.contrib.auth.models import AnonymousUser

import time


def validate_signature(request, assign_list, role):
    """
    签名算法
    """
    pass


def webservice_auth_required(assign_list=[], allow_anonymous=False, method="GET"):
    """请求基础校验装饰器；同时支持捕获APIError异常并响应
    assign_list  -- 待签名参数列表
    allow_anonymous  -- 是否允许匿名访问
    method  -- 允许的http method, 为None表示不限制method
    """

    def decorator(view_func):
        @wraps(view_func)
        def _check_authenticate(request, *args, **kwargs):
            if method and (request.method != method.upper()):
                return HttpResponseNotAllowed([method.upper()])
            
            # 获取签名的结果
            sign_check = validate_signature(request, assign_list, role='all')
            if sign_check == CommonStatus.SUCCESS:
                try:
                    if allow_anonymous:
                        return view_func(request, *args, **kwargs)
                    elif request.user.is_active and request.user.is_authenticated():
                        return view_func(request, *args, **kwargs)
                    else:
                        statuscode = CommonStatus.NOT_LOGIN
                        if request.path.rstrip('/') in getattr(settings, "LOG_SESSION_PATH", []):
                            log_session_info(request)
                except APIError as ex:
                    statuscode = ex.statuscode
                except Exception as ex:
                    SysLogger.exception(ex, request)
                    statuscode = APIError().statuscode
            else:
                statuscode = sign_check
            return ResponseContext()(request, statuscode=statuscode)

        return _check_authenticate

    return decorator

#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
#
# Copyright (c) 2018 All Rights Reserved
#
########################################################################
import logging
import traceback

from django.views import View
from django.http import HttpResponse

from django.conf import settings
TRACE_LOG = logging.getLogger(settings.PROJECT_EXCEPTION_LOG)

# Demo 不要为失败找理由,要为成功找方法


class ApiView(View):
    """
    Class as views验证
    """
    def get(self, request, *args, **kwargs):
        """
        test for use
        :param request:
        :param args:
        :param kwargs:
        :return:
        """
        try:
            a = request.parameters.get('a')
            b = request.parameters.get('b')
            c = request.parameters.get('c')

        except Exception as e:
            TRACE_LOG.info(traceback.format_exc())

        return HttpResponse("Hello world! a = %s, b = %s, c=%s" % (a, b, c))

    def post(self, request, *args, **kwargs):
        """
        post 处理逻辑
        :param request:
        :param args:
        :param kwargs:
        :return:
        """
        try:
            c = request.parameters.get('c')
            d = request.parameters.get('d')

        except Exception as e:
            TRACE_LOG.info(traceback.format_exc())

        return HttpResponse("You are a post request! c = %s, d= %s"%(c,d))


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
from libs.utils import ResponseBuilder
from django.conf import settings
TRACE_LOG = logging.getLogger(settings.PROJECT_EXCEPTION_LOG)

# 不要为失败找理由,要为成功找方法
# 不难,要你干嘛
# 目标导向
# 唯有热爱,才能抵挡危机


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

            context = {
                "a": a,
                "b": b
            }
            res_build = ResponseBuilder()
            response = res_build(context=context, status_code=200, add_response=True)
            response = ResponseBuilder.response_json(response)
        except Exception as e:
            TRACE_LOG.info(traceback.format_exc())
        return HttpResponse(response)

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


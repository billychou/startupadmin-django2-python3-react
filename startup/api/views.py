#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
#
# Copyright (c) 2018 All Rights Reserved
#
########################################################################

from django.views import View
from django.http import HttpResponse


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

        except Exception as e:
            print(e)
            a = 565
            b = 555
        return HttpResponse("Hello world! a = %s, b = %s" % (a, b))

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
            print(e)

        return HttpResponse("You are a post request! c = %s, d= %s"%(c,d ))


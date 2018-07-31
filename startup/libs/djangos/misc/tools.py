#!/usr/bin/env python
# -*- coding: utf-8 -*-
# 
# Copyright (c) 2018 All Rights Reserved
# 

"""
File: tools.py
Author: songchuan.zhou
Date: 2018/7/17 22
"""
import os
import time

import django
from django.conf import settings
from django.http import HttpResponse
from django.shortcuts import _get_queryset




def get_clientip(request, distinct=True):
    """
    获得客户端 ip
    :param request:
    :param distinct:
    :return:
    """
    serverip = request.META.get('HTTP_NS_CLIENT_IP')
    if not serverip or serverip.lower() == 'unknown':
        serverip = request.META.get('HTTP_X_FORWARDED_FOR') or ''

    if not serverip or serverip.lower() == "unknown":
        serverip = request.META.get('HTTP_PROXY_CLIENT_IP') or ''

    if not serverip or serverip.lower() == "unknown":
        serverip = request.META.get('HTTP_WL_PROXY_CLIENT_IP') or ''

    if not serverip or serverip.lower() == "unknown":
        serverip = request.META.get('REMOTE_ADDR') or ''

    if serverip and serverip.lower() != "unknown":
        if distinct:
            serverip_list = []
            for ip in serverip.split(','):
                ip = ip.strip()
                if ip and ip not in serverip_list:
                    serverip_list.append(ip)
            serverip = ','.join(serverip_list)
        return serverip
    return ''



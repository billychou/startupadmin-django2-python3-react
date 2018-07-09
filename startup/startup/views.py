#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2018 songchuan.zhou. All Rights Reserved
# 
########################################################################
 
"""
File: views.py
Author: songchuan.zhou(songchuan.zhou)
Date: 2018/07/06 09:27:36
"""
from django.shortcuts import render
from django.shortcuts import render_to_response
from django.http import HttpResponse


def index(request):
    return render(request, 'index.html')

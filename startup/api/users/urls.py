#!/usr/bin/env python
# -*- coding: utf-8 -*-
# 
# Copyright (c) 2018 All Rights Reserved
# 

"""
File: urls.py
Author: songchuan.zhou
Date: 2018/7/31 18
"""

from django.urls import path
from . import views

urlpatterns = [
    path('check-login', views.UserCheckLoginView.as_view(), name="check-login"),
]
#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
#
# Copyright (c) 2018 All Rights Reserved
#
########################################################################

from django.urls import include
from django.urls import path
from . import views

urlpatterns = [
    path('', views.ApiView.as_view(), name="api-index"),
]

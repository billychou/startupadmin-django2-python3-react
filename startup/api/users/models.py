#!/usr/bin/env python
# -*- coding: utf-8 -*-
# 
# Copyright (c) 2018 All Rights Reserved
# 

"""
File: models.py
Author: songchuan.zhou
Date: 2018/7/24 09
"""

from django.db import models
from libs.djangos.model.base import BaseModel
from datetime import datetime


class User(BaseModel):
    """用户主表"""
    username = models.CharField(max_length=100, unique=True)
    nickname = models.CharField(max_length=100, default='')
    is_active = models.BooleanField(default=True)
    date_joined = models.DateTimeField(default=datetime.now)
    last_update = models.DateTimeField(auto_now=True)

    class Meta:
        db_table = 'account_user'
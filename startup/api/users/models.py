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
from libs.djangos.model.base import BaseManager
from datetime import datetime


class User(BaseModel):
    """用户主表"""
    ROLE_CHOICES = (
        ("farmer", "农场主"),
        ("consumer", "消费者")
    )
    username = models.CharField(verbose_name="用户名", max_length=100, unique=True)
    nickname = models.CharField(verbose_name="昵称", max_length=100, default='')
    age = models.IntegerField(verbose_name="年龄")
    email = models.EmailField(verbose_name="邮箱")
    phone = models.CharField(verbose_name="电话")
    is_active = models.BooleanField(default=True)
    role = models.CharField(verbose_name="角色", max_length=100, choices=ROLE_CHOICES, default="consumer")
    date_joined = models.DateTimeField(verbose_name="注册时间", default=datetime.now)
    last_update = models.DateTimeField(verbose_name="更新时间", auto_now=True)

    @classmethod
    def user_add(cls, username, nickname, age, email, phone):
        user_obj = cls.objects.create(username=username, nickname=nickname, age=age, email=email, phone=phone)
        user_obj.save()

    class Meta:
        db_table = 'account_user'

class UserManager(BaseManager):
    pass
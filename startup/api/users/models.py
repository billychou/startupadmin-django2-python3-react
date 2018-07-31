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
from datetime import datetime

from django.db import models
from django.contrib.auth.models import AbstractBaseUser,BaseUserManager, update_last_login

from libs.djangos.model.base import BaseModel
from libs.djangos.model.base import BaseManager
from .constants import UserGenderType


class UserManager(BaseManager):
    """
    自定义 objects 方法
    """
    def create_user(self, username, password, request=None, **other_fields):
        """
        新用户注册
        :param username:
        :param password:
        :param request:
        :param other_fields:
        :return:
        """
        if not username:
            raise ValueError("The username must be given")
        email, phone = other_fields.pop("gender", "")


    def with_counts(self):
        """
        demo
        :return:
        """
        from django.db import connection
        with connection.cursor() as cursor:
            cursor.execute("""
                select * from account_user
            """)
        result_list = []
        for row in cursor.fetchall():
            p = self.model()
        return result_list


class User(BaseModel, AbstractBaseUser):
    """
    设计用户主表
    1. 用手机号登录、注册
    2. 三方账号微信登录
    """
    ROLE_CHOICES = (
        ("farmer", "农场主"),
        ("consumer", "消费者")
    )
    userid = models.IntegerField("用户 id", unique=True)
    username = models.CharField(verbose_name="用户名", max_length=100, unique=True)
    gender = models.SmallIntegerField(default=UserGenderType.UNKNOWN)
    birthday = models.DateField(null=True)
    email = models.EmailField(verbose_name="邮箱", default='')
    mobile = models.CharField(verbose_name="电话", max_length=12, default='')
    is_active = models.BooleanField(verbose_name="是否是活跃用户", default=True)
    role = models.CharField(verbose_name="角色", max_length=100, choices=ROLE_CHOICES, default="consumer")
    date_joined = models.DateTimeField(verbose_name="注册时间", default=datetime.now)
    last_update = models.DateTimeField(verbose_name="更新时间", auto_now=True)

    objects = UserManager()

    class Meta:
        db_table = 'account_user'

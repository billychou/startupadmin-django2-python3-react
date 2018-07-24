#!/usr/bin/env python
# -*- coding: utf-8 -*-
# 
# Copyright (c) 2018 All Rights Reserved
# 

"""
File: base.py
Author: songchuan.zhou
Date: 2018/7/24 09
"""

from django.db import models
from django.db.models.query import QuerySet
from django.db import router


class BaseQuerySet(QuerySet):
    def use_master(self):
        """
        返回当前 model 对应主库的 queryset
        :return:
        """
        write_db = router.db_for_write(self.model)


class BaseManager(models.Manager):
    use_for_related_fields = True

    def get_query_set(self):
        return BaseQuerySet(self.model)

    def use_master(self):
        return self.get_query_set().use_master()


class BaseModel(models.Model):
    objects = BaseManager()

    class Meta:
        abstract = True

    @classmethod
    def reverse_query_by_uid(cls, uid):
        """通过 uid 反查当前 class 对应的 queryset"""
        return cls.objects.filter(user_id=uid)

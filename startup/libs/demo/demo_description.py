#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2018 All Rights Reserved
# 
########################################################################
 
"""
File: demo_description.py
Author: songchuan.zhou(songchuan.zhou)
Date: 2018/08/01 16:34:06
"""

from functools import wraps

class Circle(object):
    """
    circle 圆形
    """
    def __init__(self, radius):
        self.radius = radius

    @property
    def get_area(self):
        """
        计算面积
        """
        ret = 3.14 * self.radius * self.radius
        return ret

    @property
    def get_circumference(self):
        """
        计算周长
        """
        ret = 2 * 3.14 * self.radius
        return ret

# 使用描述符__get__实现 Lazy 延迟初始化 

class Lazy(object):
    """
    使用描述符,实现延迟初始化
    """
    def __init__(self, func):
        self.func = func 
    
    def __get__(self, instance, cls):
        value = self.func(instance)
        setattr(instance, self.func.__name__, value)
        return value


class CircleLazy(object):
    """
    circle lazy
    """
    def __init__(self, radius):
        self.radius = radius
    
    @Lazy
    def get_area(self):
        print 'evalute'
        return 3.14 * self.radius * self.radius

    def get_circumference(self):
        print("no lazy")
        return 2 * 3.14 * self.radius
    

# 使用Property 实现延迟

def lazy_property(func):
    attr_name = "_lazy_" + func.__name__
    # 属性名
    # 装饰器
    @property
    def wrapper(self):
        if not hasattr(self, attr_name):
            setattr(self, attr_name, func(self))
        return getattr(self, attr_name)    
    return wrapper


class CircleLazyProperty(object):
    """property"""
    def __init__(self, radius):
        self.radius = radius 
    
    @lazy_property
    def get_area(self):
        return 2 * 3.14 * self.radius 
    

if __name__ == "__main__":
    c = CircleLazy(4)
    print(c.get_area)
    print "======="
    print(c.get_area)
    print(c.get_area)
    print(c.get_area)
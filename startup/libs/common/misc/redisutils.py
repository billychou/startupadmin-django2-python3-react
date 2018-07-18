#!/usr/bin/env python
# -*- coding: utf-8 -*-
# 
# Copyright (c) 2018  All Rights Reserved
# 

"""
File: redisutils.py.py
Author: songchuan.zhou (songchuan.zhou)
Date: 2018/7/12 21-22
"""

from functools import wraps

def dump_keys(client, keys, use_codis=False):
    transaction = False if use_codis else True
    pipe = client.pipeline(transaction=transaction)
    for key in keys:
        pipe.pttl(key)
        pipe.dump(key)
    return pipe.execute()





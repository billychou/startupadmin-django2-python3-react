#!/usr/bin/env python
# -*- coding: utf-8 -*-
# 
# Copyright (c) 2018 All Rights Reserved
# 

"""
File: logger_conf.py
Author: songchuan.zhou
Date: 2018/7/16 00
"""

import os
from os.path import dirname as d

lib_name = os.path.basename(d(d(d(__file__))))

# logging
LOG_ROOT = '.'

PROJECT_ACCESS_LOG = "access_log"
PROJECT_INFO_LOG = "trace_log"
PROJECT_ERROR_LOG = "trace_error_log"
PROJECT_EXCEPTION_LOG = "trace_exception_log"
PROJECT_BASESERVICE_LOG = "baseservice_log"

# The full documention for dictConf
LOGGING = {
    'version': 1,
    'disable_existing_loggers': False,
    'formatters': {
        'verbose': {
            'format': '[%(asctime)s] %(levelname)s: %(message)s \n',
        },
        'exception': {
            'format': '[%(asctime)s] %(levelname)s %(module)s Line:%(lineno)d:\n',
        },
        'trace_service': {
            'format': '%(message)s %(levelname)s %(modules)s line:%(lineno)d:\n',
        },
    },
    'filters': {
    },
    'handlers': {
        'null': {
            'level': 'DEBUG',
            'class': 'logging.NullHandler',
        },
        'console': {
            'level': 'DEBUG',
            'class': 'logging.StreamHandler',
            'formatter': 'verbose'
        },
        PROJECT_INFO_LOG: {
            'level': 'DEBUG',
            'class': lib_name + '.djangos.logger.loghandler.StartupLogFileHandler',
            'filename': os.path.join(LOG_ROOT, "logs/trace_log.log"),
            'formatter': 'verbose',
            'when': 'midnight',
        },
        PROJECT_ACCESS_LOG: {
            'level': 'DEBUG',
            'class': lib_name + '.djangos.logger.loghandler.StartupLogFileHandler',
            'filename': os.path.join(LOG_ROOT, "logs/access_log.log"),
            'formatter': 'verbose',
            'when': 'midnight',
        },
        PROJECT_EXCEPTION_LOG: {
            'level': 'DEBUG',
            'class': lib_name + '.djangos.logger.loghandler.StartupLogFileHandler',
            'filename': os.path.join(LOG_ROOT, "logs/trace_exception_log.log"),
            'formatter': 'verbose',
            'when': 'midnight',
        },
    },
    'loggers': {
        PROJECT_ACCESS_LOG: {
            'handlers': [PROJECT_ACCESS_LOG],
            'level': 'INFO',
        },
        PROJECT_INFO_LOG: {
            'handlers': [PROJECT_INFO_LOG],
            'level': 'INFO',
        },

    }

}
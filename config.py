# -*- coding: utf-8 -*-
CONFIG_DEBUG = False
ISMYSQL = True
WEB_PORT = 5000
WEB_ADDRESS = "0.0.0.0"


# 数据库配置
LOCAL_PATH = 'E:\\项目\\Ero\\'
DB_PATH = LOCAL_PATH + "novel.db"


# DEBUG == False时下列才有效
MYSQL_HOST = "0.0.0.0"
MYSQL_DATABASE = "ero_novel"
MYSQL_USERNAME = "root"
MYSQL_PASSWD = "bakabie"

# 加密类
SALT = "bakabie"
ACCESS_KEY = "BAKABIE9BIE"


# SMTP邮件发送类
SMTP_SERVER = ''
SMTP_PORT = 25
SMTP_USERNAME = ''
SMTP_PASSWD = ''


# 错误代码表
'''
0,成功
-1,用户名或者密码错误
-2,缺少参数
-4,内部错误（执行SQL操作时出错）
-8,下载点数不足
-10,权限不足
-12,下载错误
-14,未开始下载
-16,请先登录
'''

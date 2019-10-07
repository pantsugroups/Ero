# -*- coding: utf-8 -*-

CONFIG_DEBUG = True
WEB_PORT = 5000
WEB_ADDRESS = "0.0.0.0"


# 数据库配置
DB_PATH = "novel.db"


# DEBUG == False时下列才有效
MYSQL_HOST = "0.0.0.0"
MYSQL_DATABASE = "ero_novel"
MYSQL_USERNAME = "root"
MYSQL_PASSWD = "bakabie"

# 节点配置
DOWNLOAD_REMOTE_SERVER = ["http://127.0.0.1:5000",  # 首个为主节点，必须和主站为同一路径
                          ]

# 加密类
SALT = "bakabie"
ACCESS_KEY = "BAKABIE9BIE"

# 储存类
DL_SAVE_ADDRESS = "/NOVELS" #务必修改为DL服务器储存地址的全路径！！！！！
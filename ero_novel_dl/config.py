# coding:utf-8

# 是否为主力下载
'''
设置为True时必须得和web站部署在统一服务器上
或者说是web站有权限访问本目录下的Novels文件夹，已保证小说的成功上传
同时还应当在主站的config.py下添加主站web下载网址
'''
IS_MAIN = True
PORT = 5001
ADDRESS = "0.0.0.0"
# 验证参数，务必保证和主站一样
SALT = "bakabie"

DATA_ADDRESS = "Novels"
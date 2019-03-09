# -*- coding: utf-8 -*-
import sys
from config import CONFIG_DEBUG
sys.path.append('../')
from flask import  redirect, request, url_for
from flask_login import login_user, logout_user, login_required
from app import models
from app.utils import *
from config import *
from . import auth



@auth.route('/login', methods=[ 'GET','POST'])
def login():
    if request.method == 'GET':
        return jsonresp({"code":-16,"msg":"请先登陆"})
    username = request.form["user"]
    passwd = request.form['passwd']
    if not username or not passwd:
        return jsonresp({"code":-2,"msg":"缺少参数"})
    try:
        user = models.User.get(models.User.username == username)
    except Exception as e:
        return jsonresp({"code":-4,"msg":"内部错误", "error": str(e) if CONFIG_DEBUG else ""})
    if user.verify_password(passwd):
        login_user(user)
        return jsonresp({"code":0,"msg":"成功。","data":{"uid":user.id}})
    else:
        return jsonresp({"code":-1,"msg":"用户名或者密码错误"})
# 未完成
@auth.route('/register', methods=['POST'])
def register():
    username =request.form['user']
    passwd = request.form['passwd']
    mail = request.form['mail']
    if not username or not passwd or not mail:
        return jsonresp({"code": -2, "msg": "缺少参数"})
    models.User.create(username=username,password=passwd,mail=mail,lv=0)
    import smtplib
    smtpObj = smtplib.SMTP(SMTP_SERVER,SMTP_PORT)
    smtpObj.login(SMTP_USERNAME,SMTP_PASSWD)
    # smtpObj.sendmail(SMTP)
    return jsonresp({"code": 0, "msg": "成功。", })
@auth.route('/verify', methods=['GET'])
def verify():
    pass

@auth.route('/logout')
@login_required
def logout():
    logout_user()
    return jsonresp({"code": 0, "msg": "成功。", })

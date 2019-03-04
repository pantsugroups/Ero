# -*- coding: utf-8 -*-
from flask import render_template, redirect, request, url_for, flash
from flask_login import login_user, logout_user, login_required
from .. import models
from . import user
@user.route("/api/user/info")
def user_info(User):
    pass
# 修改签名，已完成
@user.route("/describe_change",methods=["POST"])
def describe_change(User):
    pass
# 关注列表，已完成
@user.route("/subscribe_list")
def subscribe_list(User):
    pass
@user.route("/link_qq")
def link_qq(User):
    pass
# 已完成
@user.route("/setting",methods=["POST"])
def setting(User):
    pass

@user.route("/msg_list")
@user.route("/msg_list/<int:page>")
def msg_list(page=1):
    pass
@user.route("/msg_get")
def msg_readed():
    pass
# 发送系统私信用
@user.route("/msg_send")
def msg_send(User):
    pass
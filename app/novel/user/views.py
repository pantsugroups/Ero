# -*- coding: utf-8 -*-
from flask import render_template, redirect, request, url_for, flash
from flask_login import login_required,current_user
from .. import models
from . import user
@user.route("/info")
@login_required
def user_info(User):
    pass
# 修改签名，已完成
@user.route("/describe_change",methods=["POST"])
@login_required
def describe_change(User):
    pass
# 关注列表，已完成
@user.route("/subscribe_list")
@login_required
def subscribe_list(User):
    pass
@user.route("/link_qq")
@login_required
def link_qq(User):
    pass
# 已完成
@user.route("/setting",methods=["POST"])
@login_required
def setting(User):
    pass

@user.route("/msg_list")
@user.route("/msg_list/<int:page>")
@login_required
def msg_list(page=1):
    pass
@user.route("/msg_get")
@login_required
def msg_readed():
    pass
# 发送系统私信用
@user.route("/msg_send")
@login_required
def msg_send(User):
    pass
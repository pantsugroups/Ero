# -*- coding: utf-8 -*-
import sys
sys.path.append('../')
from config import CONFIG_DEBUG
from flask import render_template, redirect, request, url_for, flash
from flask_login import login_required,current_user
from .. import models
from . import user
from ..utils import *
@user.route("/info")
@user.route("/info/<int:uid>")
@login_required
def user_info(uid=-1):
    if uid is -1:
        uid = current_user.id
    try:
        item = models.User.get(id=uid)
    except Exception as e:
        return jsonresp({"code": -4, "msg": "内部错误", "error": str(e) if CONFIG_DEBUG else ""})
    return jsonresp({
        "code": 0,
        "msg": "成功。",
        "data": {"user": {"avatar":item.avatar,
                          "username":item.username,
                          "mail":item.mail,
                          "register_time":item.register_time,
                          "qq":item.qq,
                          "bio":item.bio,
                          "downloads":item.downloads,
                          "pushmail":item.pushmail,
                          "lv":item.lv,
                          "hito":item.hito,
                          # "Subscribe":item.novelsubscribe
                          }}
    })

# 修改签名，已完成
@user.route("/describe_change",methods=["POST"])
@login_required
def describe_change():
    bio = request.form['bio']
    if not bio:
        return jsonresp({"code": -2, "msg": "缺少参数"})
    try:
        models.User\
            .update(bio=bio)\
            .where(models.User.id == current_user.id).execute()
    except Exception as e:
        return jsonresp({"code": -4, "msg": "内部错误", "error": str(e) if CONFIG_DEBUG else ""})
    return jsonresp({
                        "code": 0,
                        "msg": "成功。"
    })
# 关注列表，已完成
@user.route("/subscribe_list")
@user.route("/subscribe_list/<int:uid>")
@login_required
def subscribe_list(uid=-1):
    if uid is -1:
        uid = current_user.id
    try:
        items = models.User.get(models.User.id ==uid).novelsubscribe
    except Exception as e:
        return jsonresp({"code": -4, "msg": "内部错误", "error": str(e) if CONFIG_DEBUG else ""})
    result = query_to_list(items)
    return jsonresp({
        "code": 0,
        "msg": "成功。",
        "data": {"novels": result}
    })

@user.route("/link_qq")
@login_required
def link_qq(User):
    pass
# 已完成
@user.route("/change_avatar",methods=["POST"])
@login_required
def setting():
    avatar = request.form['avatar']
    if not avatar:
        return jsonresp({"code": -2, "msg": "缺少参数"})
    try:
        models.User \
            .update(avatar=avatar) \
            .where(models.User.id == current_user.id)
    except Exception as e:
        return jsonresp({"code": -4, "msg": "内部错误", "error": str(e) if CONFIG_DEBUG else ""})
    return jsonresp({
        "code": 0,
        "msg": "成功。"
    })

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
# -*- coding: utf-8 -*-
from flask import render_template, redirect, request, url_for, flash
from flask_login import current_user, login_required
from .. import models
from ..utils import *
from . import admin




@admin.route("/novel_delete/<int:nid>")
@login_required
def novel_delete(nid):
    remove_all = request.args.get("remove_all")
    if current_user.lv != 2:
        return jsonresp({
            "code":-10,
            "msg":"权限不足"
        })
    if not nid:
        return jsonresp({"code": -2, "msg": "缺少参数"})
    try:
        if remove_all:
            volumes = models.Novel.get(models.Novel.id == nid).volumes
            obj = __import__("json").loads(volumes)
            for i in obj:
                models.Volume.get(
                    models.Volume.id == i
                ).delete_instance()
        models.Novel.get(models.Novel.id == nid).delete_instance()
    except Exception as e:
        return jsonresp({"code": -4, "msg": "内部错误", "error": str(e) if CONFIG_DEBUG else ""})
    return jsonresp({"code": 0, "msg": "成功。", })

# 已完成
@admin.route("/comment_delete")
@login_required
def comment_delete(User):
    pass


@admin.route('/novel_append_volume')
@login_required
def novel_append_volume(User):
    pass


@admin.route("/novel_info_change", methods=["POST"])
@login_required
def novel_change_info(User):
    pass


@admin.route("/novel_create", methods=["POST"])
@login_required
def novel_create(User):
    pass


@admin.route("/volume_create", methods=["POST"])
@login_required
def volume_create(User):
    pass



@admin.route("/tag_create")
@login_required
def tag_create(User):
    pass


@admin.route("/tag_list")
@admin.route("/tag_list/<int:page>")
@login_required
def tag_list(User, page=1):
    pass


@admin.route("/tag_append")
@login_required
def tag_append(User):
    pass





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
    return jsonresp({"code": 0, "msg": "成功。" })

# 已完成
@admin.route("/comment_delete/<int:cid>")
@login_required
def comment_delete(cid):
    if not cid:
        return jsonresp({"code": -2, "msg": "缺少参数"})
    try:
        models.Comment.get(models.Comment.cid == cid).delete_instance()
    except Exception as e:
        return jsonresp({"code": -4, "msg": "内部错误", "error": str(e) if CONFIG_DEBUG else ""})
    return jsonresp({"code": 0, "msg": "成功。", })

@admin.route('/novel_append_volume/<int:nid>',methods=['POST'])
@login_required
def novel_append_volume(nid):
    vid = request.form['vid']
    if not nid or not vid:
        return jsonresp({"code": -2, "msg": "缺少参数"})
    try:
        l = json.loads(
            models.Novel.get(
                models.Novel.id == nid
            ).volumes
        )
        l.appent(vid)
        models.Novel.update(
            volumes = json.dumps(l)
        ).where(
            models.Novel.id == nid
        ).execute()
    except Exception as e:
        return jsonresp({"code": -4, "msg": "内部错误", "error": str(e) if CONFIG_DEBUG else ""})
    return jsonresp({"code": 0, "msg": "成功。", })


@admin.route("/novel_description_change/<int:nid>", methods=["POST"])
@login_required
def novel_change_info(nid):
    description = request.form['description']
    if not description:
        return jsonresp({"code": -2, "msg": "缺少参数"})
    try:
        models.Novel.update(
            description=description
        ).where(
            models.Novel.id == nid
        ).execute()
    except Exception as e:
        return jsonresp({"code": -4, "msg": "内部错误", "error": str(e) if CONFIG_DEBUG else ""})
    return jsonresp({"code": 0, "msg": "成功。", })

@admin.route("/novel_create", methods=["POST"])
@login_required
def novel_create():
    title = request.form["title"]
    author = request.form['author']
    cover = request.form['cover']
    description = request.form['description']
    tags = request.form["tags"]
    if not title or not author or not cover or not description or not tags:
        return jsonresp({"code": -2, "msg": "缺少参数"})
    try:
        models.Novel.create(
            title=title,
            author=author,
            cover=cover,
            description=description,
            tags=tags
        )
    except Exception as e:
        return jsonresp({"code": -4, "msg": "内部错误", "error": str(e) if CONFIG_DEBUG else ""})
    return jsonresp({"code": 0, "msg": "成功。", })

@admin.route("/volume_create/<int:nid>", methods=["POST"])
@login_required
def volume_create(nid):
    title = request.form['title']
    chapters = request.form["chapters"]
    files = request.form["files"]
    if not nid or not chapters or not files or not title:
        return jsonresp({"code": -2, "msg": "缺少参数"})
    if type(__import__("json").loads(chapters)) is not list:
        return jsonresp({"code": -2, "msg": "参数错误"})
    try:
        models.Volume.create(
            novel = nid,
            title = title,
            chapters = chapters,
            files = files
        )
    except Exception as e:
        return jsonresp({"code": -4, "msg": "内部错误", "error": str(e) if CONFIG_DEBUG else ""})
    return jsonresp({"code": 0, "msg": "成功。", })




@admin.route("/tag_create")
@login_required
def tag_create():
    pass


@admin.route("/tag_list")
@admin.route("/tag_list/<int:page>")
@login_required
def tag_list( page=1):
    pass


@admin.route("/tag_append")
@login_required
def tag_append(User):
    pass





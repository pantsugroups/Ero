# -*- coding: utf-8 -*-
import sys
from config import *

from flask import render_template, redirect, request, url_for, flash
from flask_login import current_user, login_required
from .. import models
from ..utils import *
from . import comment
@comment.route("/list/<int:cid>")
@comment.route("/list/<int:cid>/<int:page>")
def commit_list(cid = 0,page=1):
    if cid is 0 or cid <0:
        return jsonresp({"code": -2, "msg": "缺少参数"})

    try:
        items = models.Comment\
            .select()\
            .where(models.Comment.cid == cid)\
            .paginate(page, 20)
    except Exception as e:
        return jsonresp({"code": -4, "msg": "内部错误", "error": str(e) if CONFIG_DEBUG else ""})
    result = query_to_list(items)
    return jsonresp({
        "code": 0,
        "msg": "成功。",
        "data": {"novel": result}
    })

# 已完成
@comment.route("/post/<int:nid>",methods=['POST'])
@login_required
def post_comment(nid):
    content = request.form["content"]
    rep_cid = request.args.get('rep_cid')
    if nid is 0 or nid <0 or not content:
        return jsonresp({"code": -2, "msg": "缺少参数"})
    try:
        if rep_cid:
            models.Comment.create(
                novel=nid,
                user=current_user.id,
                content=content,
                rep_cid=rep_cid
            )
        else:
            models.Comment.create(
                novel=nid,
                user=current_user.id,
                content=content
            )
    except Exception as e:
        return jsonresp({"code": -4, "msg": "内部错误", "error": str(e) if CONFIG_DEBUG else ""})
    return jsonresp({
        "code": 0,
        "msg": "成功。"
    })

# 点赞，已完成
@comment.route("/like_comment/<int:cid>")
@login_required
def like_comment(cid):
    try:
        models.CommentLike.create(comment=cid,user=current_user.id)
        models.Comment.update(
            liked=models.Comment.liked+1
        ).where(models.Comment.cid == cid).execute()
    except Exception as e:
        return jsonresp({"code": -4, "msg": "内部错误", "error": str(e) if CONFIG_DEBUG else ""})
    return jsonresp({
            "code": 0,
            "msg": "成功。"
        })

# 取消赞
@comment.route("/dislike_comment/<int:cid>")
@login_required
def dislike_comment(cid):
    try:
        models.CommentLike.get(comment=cid).delete_instance()
        models.Comment.update(
            liked=models.Comment.liked - 1
        ).where(models.Comment.cid == cid).execute()
    except Exception as e:
        return jsonresp({"code": -4, "msg": "内部错误", "error": str(e) if CONFIG_DEBUG else ""})
    return jsonresp({
            "code": 0,
            "msg": "成功。"
        })

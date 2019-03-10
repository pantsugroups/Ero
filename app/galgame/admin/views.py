from flask import render_template, redirect, request, url_for, flash
from flask_login import current_user, login_required
import os
from app import models
from config import *
from app.utils import *
from . import admin
@admin.route("/create",methods = ["GET","POST"])
@login_required
def create():
    if current_user.lv is not 2:
        return "没权限",403
    if request.method == "POST":
        title = request.form["title"]
        j_title = request.form["j_title"]
        cover = request.form['cover']
        content = request.form["content"]
        tag = request.form["tag"]
        primary_str = request.form["primary_str"]
        if not title or not j_title or not cover or not content or not tag or not primary_str:
            return "少了什么东西",500

        models.Game.create(
            title=title,
            j_title=j_title,
            content=content,
            cover=cover,
            tag=tag,
            primary_str=primary_str
        )
        return jsonresp({
            "code":0,
            "data":"成功"
        })
    return render_template('game/editor.html',title="",j_title="",tag="",cover="",content="",primary_str="")

@admin.route("/delete/<int:id>",methods=["GET"])
@login_required
def delete(id=0):
    if id is 0:
        return "少了点什么",403
    if current_user.lv is not 2:
        return "错了啦",403
    models.Game.get(models.Game.id == id).delete_instance()
    return jsonresp({
        "code": 0,
        "data": "成功"
    })



@admin.route("/change/<int:id>",methods = ["GET","POST"])
@login_required
def change(id = 0):
    if current_user.lv is not 2:
        return "没权限",403
    if request.method == "POST":
        title = request.form["title"]
        j_title = request.form["j_title"]
        cover = request.form['cover']
        content = request.form["content"]
        tag = request.form["tag"]
        primary_str = request.form["primary_str"]
        if not title or not j_title or not cover or not content or not tag or not primary_str:
            return "少了什么东西",500

        models.Game.update(
            title=title,
            j_title=j_title,
            content=content,
            cover=cover,
            tag=tag,
            primary_str=primary_str
        ).where(
            models.Game.id == id
        ).execute()
        return jsonresp({
            "code":0,
            "data":"成功"
        })
    item = models.Game.get(
        models.Game.id == id
    )
    return render_template('game/editor.html',title=item.title,j_title=item.j_title
                           ,tag=item.tag,cover=item.cover,content=item.content,primary_str=item.primary_str)



@admin.route("/upload",methods = ["POST"])
@login_required
def upload():
    if request.method == 'POST':
        f = request.files['file']
        ext = f.filename[f.filename.rfind("."):]
        filename = str(time.time()) + ext
        upload_path = os.path.join(LOCAL_PATH, 'static/game/gameCover', filename)
        f.save(upload_path)
        return jsonresp({
            "coded": 0,
            "msg": "成功",
            "downloads": '/static/game/gameCover/'+filename
        })
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
    return render_template('game/editor.html',title="",j_title="",tag="",cover="",content="",primary_str="")

@admin.route("/change/<int:id>",methods = ["GET","POST"])
@login_required
def change(id = 0):
    pass

@admin.route("/change_primary/<int:id>",methods=["GET",'POST'])
@login_required
def cjamge_primary(id=0):
    pass

@admin.route("/upload",methods = ["POST"])
@login_required
def upload():
    if request.method == 'POST':
        f = request.files['file']
        ext = f.filename[f.filename.find("."):]
        upload_path = os.path.join(LOCAL_PATH, 'static/game/gameCover', str(time.time()) + ext)
        f.save(upload_path)
        return jsonresp({
            "coded": 0,
            "msg": "成功",
            "downloads": 'static/game/gameCover/'+str(time.time()) + ext
        })
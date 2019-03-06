# -*- coding: utf-8 -*-
import sys
sys.path.append('../')
from config import *
from flask import render_template, redirect, request, url_for, flash
from flask_login import login_user, current_user, login_required
from .. import models
from . import stream
from ..utils import *
@stream.route('/api/upload', methods=['POST'])
@login_required
def upload():
    pass
@stream.route("/api/novel/download/<vid>")
@stream.route("/api/novel/download/<vid>/<auto>")
def download(vid, auto=""):
    servers = int(request.args.get("servers"))
    if not vid:
        return jsonresp({"code": -2, "msg": "缺少参数"})
    if not servers or int(servers) > DOWNLOAD_REMOTE_SERVER:
        servers = DOWNLOAD_REMOTE_SERVER[0]
    else:
        servers = DOWNLOAD_REMOTE_SERVER[servers]
    try:
        links = models.Volume.get(models.Volume.id == vid).files
        token = generate_token(links[links.rfind("/") + 1:])
        if models.User.get(models.User.id == current_user.id).downloads <=0:
             return jsonresp({
                 "code":-8,
                 "msg":"下载点数不足"
             })
        models\
            .User\
            .update(
            downloads=models
                          .User\
                          .downloads-1)\
            .where(
                models\
                    .User\
                    .id == current_user.id
        )
        results = {"token": token,
                   "hash":  random_string(32),
                   "name":  models.Volume.get(models.Volume.id == vid).title }
    except Exception as e:
        return jsonresp({
            "code": -4,
            "msg": "内部错误",
            "error": str(e) if CONFIG_DEBUG else ""
        })

    return jsonresp({
        "coded":0,
        "msg":"成功",
        "downloads":results
    })
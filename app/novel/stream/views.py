# -*- coding: utf-8 -*-
import sys,os
sys.path.append('../')
from config import *
from werkzeug.utils import secure_filename
from flask import  request
from flask_login import  current_user, login_required
from app import models
import time
from ..conf import config
from . import stream
from ..utils import *
import threading
schedule = {}
def sun_delete_schedule(token):
    time.sleep(60)
    if token in schedule:
        schedule.pop(token)

@stream.route("/schedule/get")
def get_schedule():
    global schedule
    token = request.args.get("token")
    if not token:
        return jsonresp({"code": -2, "msg": "缺少参数"})
    if token not in schedule:
        return jsonresp({"code": -14, "msg": "未开始下载"})
    else:
        return jsonresp({"code":0,"msg":"成功","data":{"download": schedule[token][0], "all": schedule[token][1], "data": ""}})

@stream.route('/upload_cover', methods=['POST'])
@login_required
def upload_cover():
    if request.method == 'POST':
        f = request.files['file']
        ext = f.filename[f.filename.find("."):]
        upload_path = os.path.join('/root/Ero','static/cover',str(time.time())+ext)
        f.save(upload_path)
        return jsonresp({
            "coded":0,
            "msg":"成功",
            "downloads":upload_path
        })
@stream.route('/upload_volume', methods=['POST'])
@login_required
def upload_volume():
    if request.method == 'POST':
        title = request.args.get('title')
        f = request.files['file']
        if not title or not f:
            return jsonresp({"code": -2, "msg": "缺少参数"})
        ext = f.filename[f.filename.find("."):]
        upload_path = os.path.join(config.DL_SAVE_ADDRESS,secure_filename(title)+ext)
        f.save(upload_path)
        return jsonresp({
            "coded":0,
            "msg":"成功",
            "downloads":secure_filename(title)+ext
        })
@stream.route("/download/<vid>")
@login_required
def download(vid ):
    servers = request.args.get("servers")
    if not vid:
        return jsonresp({"code": -2, "msg": "缺少参数"})
    if not servers or int(servers) > len(config.DOWNLOAD_REMOTE_SERVER):
        servers = config.DOWNLOAD_REMOTE_SERVER[0]
    else:
        servers = config.DOWNLOAD_REMOTE_SERVER[int(servers)]
    try:
        links = models.Volume.get(models.Volume.id == vid).files
        token = generate_token(links)
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
                   "name":  models.Volume.get(models.Volume.id == vid).title,
                   "file":models.Volume.get(models.Volume.id == vid).files,
                   "server":servers}
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

@stream.route("/<file>")
def manage(file):
    token = request.args.get("token")
    rand = request.args.get("hash")
    name = request.args.get("name")
    if token != generate_token(file) or not rand or len(rand) <= 16:
        return "Permission denid!", 403
    obj = os.path.join(config.DL_SAVE_ADDRESS, file)
    if os.path.isfile(obj):
        def send_():
            with open(obj, "rb") as target_file:
                global schedule
                schedule[token] = [0, os.path.getsize(obj), int(
                    time.time()) // 600]  # 已下载,总大小,限制时间
                while True:

                    chunk = target_file.read(1024)
                    schedule[token][0] += len(chunk)
                    if not chunk:
                        t = threading.Thread(
                            target=sun_delete_schedule, args=(token,))
                        break
                    yield chunk

        response = Response(
            send_(), content_type='application/octet-stream')
        if name:

            from urllib.parse import quote
            name = quote(name)

            response.headers["Content-Disposition"] = 'attachment; filename="{}";'.format(
                name.encode("utf-8"))
            response.headers['Content-Disposition'] += "; filename*=utf-8''{}".format(
                name)
        # return Response(
        #     send_(), content_type='application/octet-stream')
        return response

    else:
        return jsonresp({
            "coded": -12,
            "msg": "下载错误",
        })
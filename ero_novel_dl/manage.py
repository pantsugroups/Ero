# coding:utf-8
from .config import *
from flask import Flask, request, render_template, jsonify, make_response, send_from_directory, Response

import hashlib
import threading
import time
import requests
import os
from urllib.parse import quote
app = Flask(__name__)

schedule = {}
@app.route("/dic_config")
def get_diconfig():
    key = request.args.get("key")
    if key != SALT:
        return "key is error."
    else:
        return os.path.join(os.getcwd(), DATA_ADDRESS)

def sun_delete_schedule(token):
    time.sleep(60)
    if token in schedule:
        schedule.pop(token)


def generate_token(_hash):
    ts = int(time.time()) // 600
    raw = str(ts) + _hash + SALT
    return hashlib.md5(raw.encode("ascii")).hexdigest()


@app.route("/schedule/get")
def get_schedule():
    global schedule
    token = request.args.get("token")
    if not token:
        return jsonify({"download": 0, "all": 0, "data": "not token"})
    if token not in schedule:
        return jsonify({"download": 0, "all": 0, "data": "no download threading."})
    else:
        return jsonify({"download": schedule[token][0], "all": schedule[token][1], "data": ""})


@app.route("/<file>")
def manage(file):
    token = request.args.get("token")
    rand = request.args.get("hash")
    name = request.args.get("name")
    if token != generate_token(file) or not rand or len(rand) <= 16:
        return "Permission denid!", 403
    obj = os.path.join(os.getcwd(), DATA_ADDRESS, file)
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
            name = quote(name)
            response.headers["Content-Disposition"] = 'attachment; filename="{}";'.format(
                name.encode("utf-8"))
            response.headers['Content-Disposition'] += "; filename*=utf-8''{}".format(
                name)
        # return Response(
        #     send_(), content_type='application/octet-stream')
        return response

    else:
        return "QAQ?????", 404


if __name__ == '__main__':

    app.run(port=PORT,host=ADDRESS)

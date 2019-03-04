# -*- coding: utf-8 -*-
import sys
sys.path.append('../')
from config import CONFIG_DEBUG
from flask import Flask, render_template, request, redirect, url_for
from json import dumps,loads
from .. import models
from ..utils import *
from ..conf.config import *
from . import novel
@novel.route("/novel")
@novel.route("/novel/<int:page>")
def index(page=1):
    try:
        items = models.Novel.select()\
            .order_by(models.Novel.update_time)\
            .paginate(page, 20)
    except Exception as e:
        return jsonresp({
            "code": -4,
            "msg": "内部错误",
            "error": str(e) if CONFIG_DEBUG else ""
        })
    result = query_to_list(items)
    return jsonresp({
        "code": 0,
        "msg": "成功。",
        "data": {"novel": result}
    })
@novel.route("/api/novel/detail/<int:nid>")
def detail(nid=0):
    if nid ==0:
        return jsonresp({"code": -2, "msg": "缺少参数"})
    try:
        items = models.Novel.get(models.Novel.id == nid)
    except Exception as e:
        return jsonresp({
            "code": -4,
            "msg": "内部错误",
            "error": str(e) if CONFIG_DEBUG else ""
        })
    result = query_to_list(items)[0]
    return jsonresp({
        "code": 0,
        "msg": "成功。",
        "data": {"novel": result}
    })
@novel.route("/search/<text>")
@novel.route("/search/<text>/<inprofile>")
def search(text, inprofile=None):
    pass
@novel.route("/volumes/<int:nid>")
def volumes(nid=0):
    if nid ==0:
        return jsonresp({"code": -2, "msg": "缺少参数"})
    try:
        items = models.Novel.get(models.Novel.id == nid).volumes
        items = loads(items)
        results = []
        for i in items:
            results.append(models.Volume.get(i))
    except Exception as e:
        return jsonresp({
            "code": -4,
            "msg": "内部错误",
            "error": str(e) if CONFIG_DEBUG else ""
        })
@novel.route("/author/<name>")
def author(name):
    pass
@novel.route("/subscribe/<nid>")
def subscribe(User, nid):
    pass
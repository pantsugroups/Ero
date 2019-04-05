# -*- coding: utf-8 -*-
from flask import Blueprint, request, jsonify, current_app
import json

from ..models import Novel, NTag

bp = Blueprint("novel", __name__)


@bp.route("/", methods=["GET"])
def novel_list():
    """
    小说列表
    ---
    tags:
      - 小说
    parameters:
      - in: query
        name: page
        type: integer
        required: false
        description: 分页页码（默认为1）
    """
    page = int(request.args.get("page", 1))
    novels = Novel.select().order_by(Novel.nid.desc()).paginate(page, 20)
    results = []
    for novel in novels:
        results.append({
            "nid": novel.nid,
            "title": novel.title,
            "author": novel.author,
            "cover": novel.cover,
            "tags": [t.name for t in novel.tags],
            "update_time": novel.update_time.strftime("%Y-%m-%d"),
            "ended": novel.ended
        })
    return jsonify({
        "status": True,
        "data": results
    })


@bp.route("/", methods=["POST"])
def novel_create():
    """
    添加新小说
    ---
    tags:
      - 小说
    parameters:
      - in: body
        name: body
        schema:
          type: object
          required:
            - title
            - author
          properties:
            title:
              type: string
              description: 书名
            author: 
              type: string
              description: 作者
            cover:
              type: string
              description: 封面
            description:
              type: string
              description: 简介
            tags:
              type: array
              items:
                type: string
              description: 标签
    """
    data = request.get_json()
    if not data:
        return jsonify({
            "status": False,
            "msg": "请以Json传入参数"
        })
    try:
        title = data["title"]
        author = data["author"]
        cover = data.get("cover")
        description = data.get("description")
        tags = data.get("tags", [])
    except KeyError:
        return jsonify({
            "status": True,
            "msg": "缺少字段"
        })
    novel = Novel.create(title=title,
                 author=author,
                 cover=cover,
                 description=description)
    for tag in tags:
        t = NTag.select().where(NTag.name == tag)
        if not t:
            t = NTag.create(name=tag)
            t.save()
        else:
            t = t.get()
        novel.tags.add(t)
    novel.save()
    return jsonify({
        "status": True,
        "data": {
            "nid": novel.nid
        }
    })


@bp.route("/<int:nid>", methods=["GET"])
def novel_detail(nid):
    """
    获取小说详细信息
    ---
    tags:
      - 小说
    parameters:
      - in: path
        name: nid
        type: integer
        required: true
        description: 小说id
    """
    try:
        novel = Novel.get(nid)
    except Novel.DoesNotExist:
        return jsonify({
            "status": True,
            "msg": "小说不存在"
        })
    return jsonify({
        "status": True,
        "data": {
            "nid": novel.nid,
            "title": novel.title,
            "author": novel.author,
            "cover": novel.cover,
            "tags": [t.name for t in novel.tags],
            "update_time": novel.update_time.strftime("%Y-%m-%d"),
            "subscribed": novel.subscribed,
            "viewed": novel.viewed,
            "liked": novel.liked,
            "ended": novel.ended
        }
    })


@bp.route("/<int:nid>", methods=["DELETE"])
def novel_delete(nid):
    """
    删除小说
    ---
    tags:
      - 小说
    parameters:
      - in: path
        name: nid
        type: integer
        description: 小说id
        required: true
    """
    try:
        novel = Novel.get(nid)
    except Novel.DoesNotExist:
        return jsonify({
            "status": False,
            "msg": "小说不存在"
        })
    with current_app.db.atomic():
        for tag in novel.tags:
            tag.novels.remove(novel)
        novel.delete_instance()
    return jsonify({
        "status": True
    })
    
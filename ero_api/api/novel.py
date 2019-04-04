# -*- coding: utf-8 -*-
from flask import Blueprint, request, jsonify

from ..models import Novel

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
        schma:
          type: integer
        description: 分页页码（默认为1）
    responses:
      "200":
        description: 返回小说列表
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
            "tags": novel.tags.split("|"),
            "update_time": novel.update_time.strftime("%Y-%m-%d"),
            "ended": novel.ended
        })
    return jsonify({
        "code": 0,
        "data": results
    })


@bp.route("/<int:nid>")
def novel_detail(nid):
    """
    获取小说详细信息
    ---
    tags:
      - 小说
    parameters:
      - in: path
        name: nid
        schma:
          type: integer
        description: 小说id
    responses:
      "200":
        description: 返回小说详细信息
    """
    try:
        novel = Novel.get(nid)
    except Novel.DoesNotExist:
        return jsonify({
            "code": 404,
        })
    return jsonify({
        "code": 0,
        "data": {
            "nid": novel.nid,
            "title": novel.title,
            "author": novel.author,
            "cover": novel.cover,
            "tags": novel.tags.split("|"),
            "update_time": novel.update_time.strftime("%Y-%m-%d"),
            "subscribed": novel.subscribed,
            "viewed": novel.viewed,
            "liked": novel.liked,
            "ended": novel.ended
        }
    })

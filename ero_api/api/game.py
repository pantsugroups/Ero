# -*- coding: utf-8 -*-
import json
import datetime

from flask import Blueprint, request, jsonify, current_app

from ..utils import require_permission
from ..models import Game, GTag, GameTag

bp = Blueprint("game", __name__)


@bp.route("/", methods=["GET"])
def get_games():
    page = int(request.args.get("page", 1))
    games = Game.select().order_by(Game.gid.desc()).paginate(page, 20)
    results = []
    for game in games:
        results.append({
            "gid": game.gid,
            "title": game.title,
            "jp_title": game.jp_title,
            "cover": game.cover,
            "description": game.description,
            "screenshot": json.loads(game.screenshot),
            "publish_time": game.publish_time.strftime(r"%Y-%m-%d")
        })
    return jsonify({
        "status": True,
        "data": results
    })


@bp.route("/", methods=["POST"])
def create_game():
    data = request.get_json()
    if data is None:
        return jsonify({
            "status": False,
            "msg": "请以Json传入参数"
        })
    try:
        title = data["title"]
        jp_title = data.get("jp_title", None)
        cover = data.get("cover")
        description = data.get("description")
        screenshot = data.get("screenshot", [])
        tags = data.get("tags", [])
        download = data.get("download", "暂不提供下载")
    except KeyError:
        return jsonify({
            "status": False,
            "msg": "缺少参数"
        })
    game = Game.create(
        title=title,
        jp_title=jp_title,
        cover=cover,
        description=description,
        screenshot=json.dumps(screenshot),
        download=download
    )
    for tag in tags:
        t = GTag.select().where(GTag.name == tag)
        if not t:
            t = GTag.create(name=tag)
            t.save()
        else:
            t = t.get()
        game.tags.add(t)
    game.save()
    return jsonify({
        "status": True,
        "data": {
            "gid": game.gid
        }
    })


@bp.route("/<int:gid>", methods=["GET"])
def game_detail(gid):
    try:
        game = Game.get(gid)
    except Game.DoesNotExist:
        return jsonify({
            "status": False,
            "msg": "游戏不存在"
        })
    return jsonify({
        "status": True,
        "data": {
            "gid": game.gid,
            "title": game.title,
            "jp_title": game.jp_title,
            "cover": game.cover,
            "description": game.description,
            "screenshot": json.loads(game.screenshot),
            "publish_time": game.publish_time.strftime(r"%Y-%m-%d")
        }
    })


@bp.route("/<int:gid>", methods=["PUT"])
@require_permission(2)
def update_game(gid):
    try:
        game = Game.get(gid)
    except Game.DoesNotExist:
        return jsonify({
            "status": False,
            "msg": "游戏不存在"
        })
    data = request.get_json()
    if data is None:
        return jsonify({
            "status": False,
            "msg": "请以Json传入参数"
        })
    allowed = {"title", "jp_title", "cover", "screenshot", "tags", "download", "description"}
    with current_app.db.atomic():
        for k, v in data.items():
            if k not in allowed:
                return jsonify({
                    "status": False,
                    "msg": "%s字段不存在或不可修改" % k
                })
            if k == "screenshot":
                data["screenshot"] = json.dumps(data["screenshot"])
            elif k == "tags":
                after_tags = set()
                for tag in v:
                    t = GTag.select().where(GTag.name == tag)
                    if not t:
                        t = GTag.create(name=tag)
                    else:
                        t = t.get()
                    after_tags.add(t)
                current_tags = set(game.tags)
                delete = current_tags - after_tags
                add = after_tags - current_tags
                for tag in delete:
                    game.tags.remove(tag)
                    tag.games.remove(game)
                for tag in add:
                    game.tags.add(tag)
        if "tags" in data:
            data.pop("tags")
        Game.update(**data).where(Game.gid == game.gid).execute()
    return jsonify({
        "status": True
    })


@bp.route("/<int:gid>", methods=["DELETE"])
@require_permission(2)
def delete_game(gid):
    try:
        game = Game.get(gid)
    except Game.DoesNotExist:
        return jsonify({
            "status": False,
            "msg": "游戏不存在"
        })
    with current_app.db.atomic():
        for tag in game.tags:
            tag.games.remove(game)
        game.delete_instance()
    return jsonify({
        "status": True
    })

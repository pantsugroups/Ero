# -*- coding: utf-8 -*-
import datetime
import json

from flask import Blueprint, request, jsonify, current_app
from flask_restful import Resource

from ..models import Novel, NTag
from ..utils import require_permission


class NovelItem(Resource):

    def get(self, nid):
        try:
            novel = Novel.get(nid)
        except Novel.DoesNotExist:
            return jsonify({
                "status": False,
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

    @require_permission(2)
    def delete(self, nid):
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
    
    @require_permission(2)
    def put(self, nid):
        try:
            novel = Novel.get(nid)
        except Novel.DoesNotExist:
            return jsonify({
                "status": False,
                "msg": "小说不存在"
            })
        data = request.get_json()
        if data is None:
            return jsonify({
                "status": False,
                "msg": "请以Json传入参数"
            })
        allowed = {"title", "author", "cover", "ended", "tags", "hide", "description"}
        with current_app.db.atomic():
            for k, v in data.items():
                if k not in allowed:
                    return jsonify({
                        "status": False,
                        "msg": "%s字段不存在或不可修改" % k
                    })
                if k == "tags":
                    after_tags = set()
                    for tag in v:
                        t = NTag.select().where(NTag.name == tag)
                        if not t:
                            t = NTag.create(name=tag)
                        else:
                            t = t.get()
                        after_tags.add(t)
                    current_tags = set(novel.tags)
                    delete = current_tags - after_tags
                    add = after_tags - current_tags
                    for tag in delete:
                        novel.tags.remove(tag)
                        tag.novels.remove(novel)
                    for tag in add:
                        novel.tags.add(tag)
            if "tags" in data:
                data.pop("tags")
            data["update_time"] = datetime.datetime.now()
            Novel.update(**data).where(Novel.nid == novel.nid).execute()
        return jsonify({
            "status": True
        })


class NovelList(Resource):

    def get(self):
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
        
    @require_permission(2)
    def post(self):
        data = request.get_json()
        if data is None:
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
            ended = data.get("ended", False)
        except KeyError:
            return jsonify({
                "status": False,
                "msg": "缺少参数"
            })
        novel = Novel.create(title=title,
                    author=author,
                    cover=cover,
                    description=description,
                    ended=ended)
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

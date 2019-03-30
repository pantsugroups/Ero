from flask import render_template, redirect, request, url_for, flash
import os
from flask_login import current_user, login_required
from app import models
from app.utils import *
from config import *
from . import index
import markdown
@index.route('/api/index')
@index.route('/api/index/<int:page>')
def api_indexs(page=1):
    try:
        items = models.Game.select() \
        .order_by(models.Game.post_time) \
        .paginate(page, 20)
        count = models.Game.select().count()
    except Exception as e:
        return "",404
    next=0
    last=0
    if page*20<=count:
        next=1
    elif page != 1:
        last=1
    print(next,last)
    items=query_to_list(items)
    return jsonresp({
        "code": 0,
        "msg": "成功。",
        "data":items,
        "next":next,
        "last":last,
    })

@index.route("/api/view/<int:id>")
def api_game(id=0):

    try:
        item = models.Game.get(
        models.Game.id == id
    )
    except Exception as e:
        return "",404
    if request.args.get("markdown") == "true":
        html = markdown.markdown(item.content)
    else:
        html = item.content
    # return render_template('game/view.html',html=html,title=item.title,id=item.id)
    return jsonresp({
        "code": 0,
        "msg": "成功。",
        "date": html
    })

@index.route("/api/view_primary/<int:id>")
def api_primary_string(id=0):
    try:
        item = models.Game.get(
        models.Game.id==id
    ).primary_str
    except Exception as e:
        return "",404
    if request.args.get("markdown") == "true":
        html = markdown.markdown(item)
    else:
        html = item
    return jsonresp({
        "code": 0,
        "msg": "成功。",
        "date": html
    })
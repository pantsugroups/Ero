from flask import render_template, redirect, request, url_for, flash
import os
from flask_login import current_user, login_required
from app import models
from app.utils import *
from config import *
from . import index
import markdown
@index.route('/')
@index.route('/<int:page>')
def indexs(page=1):
    try:
        items = models.Game.select(
            models.Game.id,
            models.Game.cover,
            models.Game.title,
            models.Game.j_title,
            models.Game.tag
        ) \
        .order_by(models.Game.post_time.desc()) \
        .paginate(page, 20)
        count = models.Game.select().count()
    except Exception as e:
        return "",404
    next,last=False,False
    if page*20<=count:
        next=True
    elif page != 1:
        last=True
    return render_template('game/index.html',items=items,next=next,last=last)

@index.route('/search/')
@index.route('/search/<int:page>')
def search(page=1):
    text=request.args.get("text")
    try:
        items = models.Game.select() \
        .where(models.Game.title ** '%'+text+'%')\
        .order_by(models.Game.post_time) \
        .paginate(page, 20)
        count = models.Game.select() \
            .where(models.Game.title ** '%' + text + '%') \
            .count()

    except Exception as e:
        return "",404
    next, last = False, False
    if page * 20 <= count:
        next = page
    elif page != 1:
        last = page
    return render_template('game/index.html',items=items,next=next,last=last)

@index.route("/view/<int:id>")
def game(id=0):
    try:
        item = models.Game.get(
        models.Game.id == id
    )
    except Exception as e:
        return "",404
    html = markdown.markdown(item.content)
    return render_template('game/view.html',html=html,title=item.title,id=item.id)

@index.route("/view_primary/<int:id>")
def primary_string(id=0):
    try:
        item = models.Game.get(
        models.Game.id==id
    ).primary_str
    except Exception as e:
        return "",404
    html = markdown.markdown(item)
    return html
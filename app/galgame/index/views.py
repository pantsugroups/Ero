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
        items = models.Game.select() \
        .order_by(models.Game.post_time) \
        .paginate(page, 20)
    except Exception as e:
        return "",404
    return render_template('game/index.html',items=items)

@index.route('/search/<text>')
@index.route('/search/<text>/<int:page>')
def search(text,page=1):
    try:
        items = models.Game.select() \
        .where(models.Game.title ** '%'+text+'%')\
        .order_by(models.Game.post_time) \
        .paginate(page, 20)
    except Exception as e:
        return "",404
    return render_template('game/index.html',items=items)

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
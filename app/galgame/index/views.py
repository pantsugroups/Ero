from flask import render_template, redirect, request, url_for, flash
import os
from flask_login import current_user, login_required
from app import models
from app.utils import *
from config import *
from . import index
@index.route('/')
@index.route('/<int:page>')
def indexs(page=1):
    items = models.Game.select() \
        .order_by(models.Game.post_time) \
        .paginate(page, 20)
    return render_template('game/index.html',items=items)

@index.route("/game/<int:id>")
def game(id=0):
    pass

@index.route("/game_primary/<int:id>")
def primary_string(id=0):
    pass
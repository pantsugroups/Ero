# -*- coding: utf-8 -*-
from flask import render_template, redirect, request, url_for, flash
from flask_login import login_user, logout_user, login_required
from .. import models

from . import comment
@comment.route("/list")
@comment.route("/list/<int:page>")
def commit_list(page=1):
    pass

# 已完成
@comment.route("/post",methods=['POST'])
def post_comment():
    pass

# 点赞，已完成
@comment.route("/like_comment")
def like_comment():
    pass

# 取消赞
@comment.route("/dislike_comment")

def dislike_comment():
    pass

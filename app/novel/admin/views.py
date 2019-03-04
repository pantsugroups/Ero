# -*- coding: utf-8 -*-
from flask import render_template, redirect, request, url_for, flash
from flask_login import login_user, logout_user, login_required
from .. import models

from . import admin


@admin.route("/novel_delete")
def novel_delete(User):
    pass


# 已完成
@admin.route("/comment_delete")
def comment_delete(User):
    pass


@admin.route('/novel_append_volume')
def novel_append_volume(User):
    pass


@admin.route("/novel_info_change", methods=["POST"])
def novel_change_info(User):
    pass


@admin.route("/novel_create", methods=["POST"])
def novel_create(User):
    pass


@admin.route("/volume_create", methods=["POST"])
def volume_create(User):
    pass


@admin.route("/workist")
def workist(User):
    pass


@admin.route("/workist_accpet")
def workist_accpet(User):
    pass


@admin.route("/workist_refuse")
def workist_refuse(User):
    pass


@admin.route("/workist_clean")
def workist_clean(User):
    pass


@admin.route("/tag_create")
def tag_create(User):
    pass


@admin.route("/tag_list")
@admin.route("/tag_list/<int:page>")
def tag_list(User, page=1):
    pass


@admin.route("/tag_append")
def tag_append(User):
    pass





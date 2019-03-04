# -*- coding: utf-8 -*-
from flask import render_template, redirect, request, url_for, flash
from flask_login import login_user, logout_user, login_required
from .. import models
from . import stream
@stream.route('/api/upload', methods=['POST'])
def upload():
    pass
@stream.route("/api/novel/download/<vid>")
@stream.route("/api/novel/download/<vid>/<auto>")
def download(User, vid, auto=""):
    pass
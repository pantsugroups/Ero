# -*- coding: utf-8 -*-
from flask import render_template, redirect, request, url_for, flash
from flask_login import login_user, logout_user, login_required
import app.novel.models

from . import auth



@auth.route('/login', methods=['GET', 'POST'])
def login():
    pass

@auth.route('/register', methods=['GET', 'POST'])
def login():
    pass


@auth.route('/logout')
@login_required
def logout():
    logout_user()
    flash(u'您已退出登录')
    return redirect(url_for('auth.login'))

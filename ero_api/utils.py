# -*- coding: utf-8 -*-
from hashlib import sha256

from flask import current_app, session, jsonify, request

from .models import User


def require_login(func):

    def warpper(*args, **kwargs):
        if "uid" not in session or "username" not in session:
            return jsonify({
                "status": False,
                "msg": "请先登录"
            })
        try:
            user = User.get(int(session["uid"]))
        except User.DoesNotExist:
            session.pop("uid", None)
            session.pop("username", None)
            session.pop("login_ts", None)
            return jsonify({
                "status": False,
                "msg": "请先登录"
            })
        request.user = user
        return func(*args, **kwargs)
 
    return warpper


def encrypt_pwd(username, pwd):
    password = sha256((username + \
                       pwd + \
                       current_app.config["SECRET_KEY"]).encode("utf-8")).hexdigest()
    return password
    
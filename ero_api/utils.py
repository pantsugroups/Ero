# -*- coding: utf-8 -*-
from hashlib import sha256
from functools import wraps 

from flask import current_app, session, jsonify, request

from .models import User


def require_login(func):

    @wraps(func)
    def wrapper(*args, **kwargs):
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
 
    return wrapper


def require_permission(permission):
    
    def outer_wrapper(func):
    
        @wraps(func)
        def wrapper(*args, **kwargs):
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
            if user.permission <= permission:
                return jsonify({
                    "status": False,
                    "msg": "操作受限"
                })
            request.user = user
            return func(*args, **kwargs)
        
        return wrapper
 
    return outer_wrapper


def encrypt_pwd(username, pwd):
    password = sha256((username + \
                       pwd + \
                       current_app.config["SECRET_KEY"]).encode("utf-8")).hexdigest()
    return password

# -*- coding: utf-8 -*-
import time
from re import match as re_match
from random import choice
import string

from flask import Blueprint, request, jsonify, current_app, session

from ..models import User
from ..utils import require_login, encrypt_pwd

bp = Blueprint("auth", __name__)


@bp.route("/login", methods=["POST"])
def login():
    """
    登陆
    ---
    tags:
      - 账号
    parameters:
      - in: body
        name: body
        schema:
          type: object
          required:
            - username
            - password
          properties:
            username:
              type: string
              description: 用户名
            password:
              type: string
              description: 密码
    """
    data = request.get_json()
    if data is None:
        return jsonify({
            "status": False,
            "msg": "请用Json传入参数"
        })
    try:
        username = data["username"]
        password_raw = data["password"]
    except KeyError:
        return jsonify({
            "status": False,
            "msg": "缺少字段"
        })
    password = encrypt_pwd(username, password_raw)
    user = User.select().where(User.username == username)
    if not user:
        return jsonify({
            "status": False,
            "msg": "用户名不存在"
        })
    user = user.get()
    if user.password != password:
        return jsonify({
            "status": False,
            "msg": "密码错误"
        })
    session["uid"] = user.uid
    session["username"] = user.username
    session["login_ts"] = int(time.time())
    return jsonify({
        "status": True,
        "data": {
            "uid": user.uid,
            "username": user.username,
            "nickname": user.nickname
        }
    })
    

@bp.route("/logout", methods=["GET"])
@require_login
def logout():
    session.pop("uid", None)
    session.pop("username", None)
    session.pop("login_ts", None)
    return jsonify({
        "status": True
    })


@bp.route("/register", methods=["POST"])
def register():
    """
    注册
    ---
    tags:
      - 账号
    parameters:
      - in: body
        name: body
        schema:
          type: object
          required:
            - username
            - email
            - password
          properties:
            username:
              type: string
              description: 用户名
            email:
              type: string
              description: 邮箱
            password:
              type: string
              description: 密码
    """
    data = request.get_json()
    if data is None:
        return jsonify({
            "status": False,
            "msg": "请用Json传入参数"
        })
    try:
        username = data["username"]
        email = data["email"]
        password_raw = data["password"]
    except KeyError:
        return jsonify({
            "status": False,
            "msg": "缺少字段"
        })
    if len(username) < 3:
        return jsonify({
            "status": False,
            "msg": "用户名过短，最少4个字符"
        })
    elif len(username) > 20:
        return jsonify({
            "status": False,
            "msg": "用户名过长，最多20个字符"
        })
    if len(password_raw) < 8:
        return jsonify({
            "status": False,
            "msg": "密码过短，至少8个字符"
        })
    elif len(password_raw) > 15:
        return jsonify({
            "status": False,
            "msg": "密码过长，最多15个字符"
        })
    if not re_match("^([a-zA-Z0-9_\-\.]+)@([a-zA-Z0-9_\-\.]+)\.([a-zA-Z]{2,5})$", email):
        return jsonify({
            "status": False,
            "msg": "邮件地址不合法"
        })
    if User.select().where(User.username == username):
        return jsonify({
            "status": False,
            "msg": "用户名已存在"
        })
    password = encrypt_pwd(username, password_raw)
    user = User.create(username=username,
                       nickname=username,
                       password=password,
                       email=email)
    user.save()
    session["uid"] = user.uid
    session["username"] = user.username
    session["login_ts"] = int(time.time())
    return jsonify({
        "status": True,
        "data": {
            "uid": user.uid,
            "username": username,
            "nickname": username
        }
    })
    

@bp.route("/create_administrator", methods=["GET"])
def create_administrator():
    """
    创建最高管理员，用户名administrator，密码随机，仅可创建一次
    ---
    tags:
      - 账号
    """
    if User.select().where(User.username == "administrator"):
        return jsonify({
            "status": False,
            "msg": "最高管理员已存在"
        })
    else:
        pwd = "".join(choice(string.ascii_uppercase + string.digits) for _ in range(15))
        User.create(username="administrator",
                    nickname="狗管理",
                    email="administrator@ero.ink",
                    password=encrypt_pwd("administrator", pwd),
                    permission=999).save()
        return jsonify({
            "status": True,
            "data": {
                "password": pwd
            }
        })

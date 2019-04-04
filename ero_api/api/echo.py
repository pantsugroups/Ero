# -*- coding: utf-8 -*-
from flask import Blueprint, request

bp = Blueprint("echo", __name__)


@bp.route("/", methods=["GET"])
def echo():
    """
    测试API
    ---
    tags:
      - 测试
    parameters:
      - in: query
        name: msg
        schma:
          type: string
        description: 返回内容（默认为Hello World!）
    responses:
      200:
        description: 成功
    """
    msg = request.args.get("msg", "Hello World!")
    return msg

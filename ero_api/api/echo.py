# -*- coding: utf-8 -*-
from flask import Blueprint, request

bp = Blueprint("echo", __name__)


@bp.route("/", methods=["GET"])
def echo():
    msg = request.args.get("msg", "Hello World!")
    return msg

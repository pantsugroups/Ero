# -*- coding: utf-8 -*-
from flask import Blueprint, request

from ..models import User

bp = Blueprint("user", __name__)

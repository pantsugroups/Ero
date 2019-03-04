# -*- coding: utf-8 -*-
from flask import Blueprint

auth = Blueprint('comment', __name__)

from . import views

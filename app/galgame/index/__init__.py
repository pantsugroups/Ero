# -*- coding: utf-8 -*-
from flask import Blueprint

admin = Blueprint('index', __name__)

from . import views

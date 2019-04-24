# -*- coding: utf-8 -*-
from flask import Blueprint

admin = Blueprint('novel_admin', __name__)

from . import views,api

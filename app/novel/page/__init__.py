# -*- coding: utf-8 -*-
from flask import Blueprint

page = Blueprint('novel', __name__)

from . import views

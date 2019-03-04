# -*- coding: utf-8 -*-
from flask import Blueprint

novel = Blueprint('novel', __name__)

from . import views

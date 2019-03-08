# -*- coding: utf-8 -*-
from flask import Blueprint

stream = Blueprint('stream', __name__)

from . import views

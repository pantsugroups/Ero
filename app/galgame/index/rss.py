from flask import render_template, redirect, request, url_for, flash
import os
from flask_login import current_user, login_required
from app import models
from app.utils import *
from config import *
from . import index
import markdown
@index.route('/rss')
@index.route('/rss/<int:page>')
def rss(page=1):
    pass
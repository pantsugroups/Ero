import sys
from config import CONFIG_DEBUG
sys.path.append('../')
from flask import  redirect, request, url_for,render_template
from flask_login import login_user, logout_user, login_required
from app import models
from app.utils import *
from config import *
from . import auth

@auth.route('/login2', methods=[ 'GET','POST'])
def login2():
    return render_template("auth/login.html")
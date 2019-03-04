# -*- coding: utf-8 -*-
from flask import Flask
from flask_cors import *
from app.novel import create
app = Flask(__name__)
from app import novel
CORS(app, supports_credentials=True)
app.register_blueprint(novel, url_prefix="/erolight")

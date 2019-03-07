# -*- coding: utf-8 -*-
from flask import Flask
from flask_cors import *
from app.novel import create
from config import *
app = Flask(__name__)
# from app import novel
CORS(app, supports_credentials=True)
# app.register_blueprint(novel.app, url_prefix="/erolight")
if __name__ == "__main__":
    app = create(app)
    app.run(port=WEB_PORT,host=WEB_ADDRESS)

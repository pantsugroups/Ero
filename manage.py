# -*- coding: utf-8 -*-
from flask import Flask
from flask_cors import *
from app import create
from app.models import create_table
from config import *
import sys
app = Flask(__name__)
# from app import novel
CORS(app, supports_credentials=True)
# app.register_blueprint(novel.app, url_prefix="/erolight")
app = create(app)
if __name__ == "__main__":
    if len(sys.argv) is not 1:create_table()

    app.run(port=WEB_PORT,host=WEB_ADDRESS)

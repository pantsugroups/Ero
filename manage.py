# -*- coding: utf-8 -*-
from flask import Flask
from flask_cors import *
from app.novel import create
app = Flask(__name__)
from app import novel
CORS(app, supports_credentials=True)
app.register_blueprint(novel.app, url_prefix="/admin")

# app=Flask(__name__)
# if __name__ == "__main__":
#
#     app = create(app)
#     app.run(host="0.0.0.0",debug=True,threaded=True)

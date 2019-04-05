# -*- coding: utf-8 -*-
from flask import Flask, jsonify
from flask_restful import Api
from flasgger import Swagger
from flask_cors import CORS

from . import api, models


def create_app(config):
    app = Flask(__name__)
    app.config.from_object(config)
    
    CORS(app)
    restful = Api(app)
    Swagger(app)

    @app.before_request
    def _connect_db():
        app.db.connect()
    
    @app.after_request
    def _close_db(response):
        app.db.close()
        return response

    models.db.initialize(config.DB)
    app.db = models.db

    app.register_blueprint(api.echo.bp, url_prefix="/echo")
    app.register_blueprint(api.auth.bp, url_prefix="/auth")
    restful.add_resource(api.novel.NovelList, "/novel")
    restful.add_resource(api.novel.NovelItem, "/novel/<int:nid>")
    return app


def initialize(config):
    models.db.initialize(config.DB)
    models.db.connect()
    models.db.create_tables([
        models.User,
        models.Novel,
        models.UserNovelSubscribe,
        models.NTag,
        models.NovelTag,
        models.Volume,
        models.NovelComment,
        models.Game,
        models.GTag,
        models.GameTag
    ])
    models.db.commit()
    models.db.close()

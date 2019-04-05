# -*- coding: utf-8 -*-
from flask import Flask, jsonify
from flask_swagger import swagger
from flask_swagger_ui import get_swaggerui_blueprint

from . import api, models


def create_app(config):
    app = Flask(__name__)

    app.config.from_object(config)

    @app.before_request
    def connect_db():
        app.db.connect()
    
    @app.after_request
    def close_db(response):
        app.db.close()
        return response

    swaggerui_blueprint = get_swaggerui_blueprint("/swagger",
                                                  "/swagger/spec",
                                                  {"app_name": "Ero API"})

    @swaggerui_blueprint.route("/spec")
    def swagger_json():
        return jsonify(swagger(app))

    models.db.initialize(config.DB)
    app.db = models.db

    app.register_blueprint(api.echo.bp, url_prefix="/echo")
    app.register_blueprint(api.novel.bp, url_prefix="/novel")
    app.register_blueprint(swaggerui_blueprint, url_prefix="/swagger")
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

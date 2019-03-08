# -*- coding: utf-8 -*-
from flask_login import LoginManager

login_manager = LoginManager()
login_manager.session_protection = 'strong'
login_manager.login_view = 'auth.login'

def create(app):

    app.config['SECRET_KEY'] = '123456'
    app.config['MAX_CONTENT_LENGTH'] = 20 * 1024 * 1024
    login_manager.init_app(app)
    from .admin import admin as admin_blueprint
    app.register_blueprint(admin_blueprint, url_prefix="/admin")

    from .stream import stream as stream_blurprint
    app.register_blueprint(stream_blurprint, url_prefix="/stream")
    from .auth import auth as auth_blueprint
    app.register_blueprint(auth_blueprint,url_prefix="/auth")

    from .novel import novel as novel_blueprint
    app.register_blueprint(novel_blueprint,url_prefix="/novel")

    from .user import user as user_blueprint
    app.register_blueprint(user_blueprint, url_prefix="/user")

    from .comment import comment as comment_blueprint
    app.register_blueprint(comment_blueprint, url_prefix="/comment")
    return app



# -*- coding: utf-8 -*-
from flask_login import LoginManager

login_manager = LoginManager()
login_manager.session_protection = 'strong'
login_manager.login_view = 'auth.login'

def create(app):

    app.config['SECRET_KEY'] = '123456'
    login_manager.init_app(app)

    # from .main import main as main_blueprint
    # app.register_blueprint(main_blueprint,url_prefix="/admin")
    #
    # from .auth import auth as auth_blueprint
    # app.register_blueprint(auth_blueprint,url_prefix="/admin")
    return app



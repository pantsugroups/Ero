#coding:utf-8
from flask_login import LoginManager

login_manager = LoginManager()
login_manager.session_protection = 'strong'
login_manager.login_view = 'auth.login'
login_manager.login_message='对不起，您还没有登录'

def create(app):

    app.config['SECRET_KEY'] = '123456'
    app.config['MAX_CONTENT_LENGTH'] = 20 * 1024 * 1024
    login_manager.init_app(app)
    from .auth import auth as auth_blueprint
    app.register_blueprint(auth_blueprint, url_prefix="/auth")

    from .novel.admin import admin as admin_blueprint
    app.register_blueprint(admin_blueprint, url_prefix="/light/admin")

    from .novel.stream import stream as n_stream_blurprint
    app.register_blueprint(n_stream_blurprint, url_prefix="/light/stream")


    from .novel.novel import novel as n_novel_blueprint
    app.register_blueprint(n_novel_blueprint,url_prefix="/light/novel")

    from .novel.user import user as n_user_blueprint
    app.register_blueprint(n_user_blueprint, url_prefix="/light/user")

    from .novel.comment import comment as n_comment_blueprint
    app.register_blueprint(n_comment_blueprint, url_prefix="/light/comment")




    from .galgame.index import index as g_index_bluepring
    app.register_blueprint(g_index_bluepring, url_prefix="/game/index")
    from .galgame.admin import admin as g_admin_bluepring
    app.register_blueprint(g_admin_bluepring, url_prefix="/game/admin")
    return app
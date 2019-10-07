# -*- coding: utf-8 -*
from flask import Flask

from web_api import *
from web_view import *


@app.route("/api/First_Init")
@Grant
def init():
    # 第一次登录到ero
    username = request.args.get("username")
    qqnumber = request.args.get("qqnumber")
    referer = request.args.get("referer")
    if not username or not qqnumber:
        return render_template("register.html",referer=referer)
    else:
        source_uid = request.cookies.get("globaluid")
        if not source_uid:
            return "403",403
        try:
            User.create(source_uid=int(source_uid), username=username.encode("utf-8"), qq=int(qqnumber))
            if referer:
                return redirect(referer)
            else:
                return "OJBK"
        except Exception as e:
            print(e)
            return "ERROR"


@app.route("/User_Setting/<int:uid>")
@Grant
@AccessControl
def User_Setting(User):
    pass


if __name__ == '__main__':
    app.run(port=WEB_PORT,host=WEB_ADDRESS)

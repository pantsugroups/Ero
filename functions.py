# -*- coding: utf-8 -*-
from flask import jsonify, request, redirect, url_for, Flask,make_response
import functools
import hashlib
import time
import string
import random
from config import *
from kernel import *
from urllib.parse import unquote

app = Flask(__name__)


def GetPathAndFileName(path):
    return path[:path.rfind("/")], path[path.rfind("/") + 1:]


def RandomString(long):
    return ''.join(random.sample(string.ascii_letters + string.digits, long))


def __HashValidate(uid, raw):
    if CONFIG_DEBUG:
        return True, "TESTAUTHCODE."
    raw = unquote(raw)
    data = typecho_users.get(typecho_users.uid == uid).authCode
    if not data:
        return False, 0
    else:
        authcode = data
        salt = raw[3:12]
        salt2 = ""
        hashs = ""
        last = ord(authcode[len(authcode) - 1])
        pos = 0
        while pos < len(authcode):
            asc = ord(authcode[pos])
            last = (last * ord(salt[last % asc % 9]) + asc) % 95 + 32
            hashs += chr(last)
            pos += 1
        md5 = hashlib.md5()
        md5.update(hashs.encode("utf-8"))
        # print('$T$' + salt, md5.hexdigest(), raw)
        if '$T$' + salt + md5.hexdigest() == raw:
            return True, authcode
        else:
            return False, 0


def Grant(func):
    # 鉴权
    # 请务必先调用flak的装饰器函数，否则将会出现逻辑错误
    @functools.wraps(func)
    def Verify(*args, **kwargs):
        # 假设获取到了uid,这里的uid指的是typecho的uid
        if CONFIG_DEBUG == True:
            source_uid = 1
        else:
            source_uid = request.cookies.get("globaluid")
            if not source_uid:
                return redirect("https://ero.ink/login.html?referer="+request.url)
            else:
                source_uid = int(source_uid)
            isok, r = __HashValidate(
                source_uid, request.cookies.get("globalauthcode"))
            print(r)
            if not isok:
                return redirect("https://ero.ink/login.html?referer="+request.url)
        if (request.path == "/First_Init"):
            result = func(*args, **kwargs)
            return result
        if not not source_uid:
            try:
                user = User_Get_1(source_uid)
            except Exception as e:
                print(e)
                # 说明是第一次来到ero light,要求进行初始化认证
                return redirect(url_for("init",referer=request.url))
            kwargs["User"] = {"logined": True}
            kwargs["User"]["uid"] = user.uid  # 这里的uid指的是我们本地数据库的uid
            kwargs["User"]["username"] = user.username
            kwargs["User"]["oldriver"] = user.old_driver
            kwargs["User"]["admin"] = user.admin # 0为普通用户，1为管理员
            # kwargs["User"]["token"] = User_Get_Token(user.uid)
        else:
            kwargs["User"] = {"lgoined": False}
        result = func(*args, **kwargs)
        return result
    return Verify


def AccessControl(func):
        # 用户访问控制器
    @functools.wraps(func)
    def Judge(*args, **kwargs):
        # todo 使用上面那个函数进行验证，如果通过则
        # debug..暂时先设置全部通过
        if "User" in kwargs and kwargs["User"]["logined"] == True:
            if "nid" in kwargs:
                Novel_ViewUp(int(kwargs["nid"]))
            return func(*args, **kwargs)
        else:
            return json_resp({}, -999)
    return Judge

# 直接用天台的s


def generate_token(_hash):
    ts = int(time.time()) // 600
    raw = str(ts) + _hash + SALT
    return hashlib.md5(raw.encode("ascii")).hexdigest()


def json_resp(data, code=0, msg=None):
    message = {
        -999: "请先登录",
        - 666: "Permission Error.",
        -100: "请求参数错误",
        -200: "nid不存在",
        - 500: "可用下载点数不足",
        - 111:"未知错误",
        -101: "你懂得",#没有权限浏览老司机模式
    }

    if code == 0:
        response = make_response(jsonify({"code": 0, "data": data}))
    else:
        response=make_response(jsonify({"code": code, "msg": msg or message[code], "data": data}))
    response.headers['Access-Control-Allow-Origin'] = '*'
    response.headers['Access-Control-Allow-Methods'] = 'OPTIONS,HEAD,GET,POST'
    response.headers['Access-Control-Allow-Headers'] = 'x-requested-with'
    return response 

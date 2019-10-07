# -*- coding: utf-8 -*-
from flask import Flask, render_template, request, redirect, url_for
from werkzeug.utils import secure_filename
import os

from functions import *
##########################################
#@@@@          用户部分开始           @@@@@#
##########################################


# 已完成
@app.route("/api/novel/index")
@app.route("/api/novel/index/<int:page>")
@Grant
def index(User,page=1):
    if User["oldriver"] == 0:
        data = Novel_Index(page,False,False)
        # return json_resp(data)
    else:
        data = Novel_Index(page,False,True)
        # results = []
        # for i in data:
        #         results.append({
        #             "nid":i.nid,
        #             "title":i.title,
        #             "author":i.author,
        #             "cover":i.cover,
        #             "description":i.description,
        #             "tags":i.tags,
        #             "update_time":i.update_time,
        #             "subscribed":i.subscribed,
        #             "viewed":i.viewed,
        #             "liked":i.liked,
        #             "ended":i.ended,
        #             "volumes":i.volumes
        #     })
    return json_resp(data)


@app.route("/api/novel/detail/<int:nid>")
@Grant
def detail(User,nid):
    # nid = request.args.get("nid") 
    if nid is None:
        return json_resp([], -100), 500
    try:
        data = Novel_Get(int(nid))
    except ValueError:
        return json_resp([], -100), 500
    if not data:
        return json_resp([], -200), 500
    if User["oldriver"] == 0:
        if data.hide == 1:
            return json_resp([],-101)
    return json_resp({
        "nid": data.nid,
        "title": data.title,
        "author": data.author,
        "cover": data.cover,
        "description": data.description,
        "tags": data.tags.split(","),
        "update_time": data.update_time.strftime("%Y-%m-%d"),
        "subscribed": data.subscribed,
        "viewed": data.viewed,
        "liked": data.liked,
        # "disliked": data.disliked,
        "volumes": json.loads(data.volumes)
    })

# 已完成
@app.route("/api/novel/volumes/<int:nid>")
@Grant
@AccessControl
def volumes(User,nid):
    # nid = request.args.get("nid")
    if nid is None:
        return json_resp([], -100)
    try:
        novel = Novel_Get_Volume_Name(int(nid))
    except ValueError:
        return json_resp([], -100)
    except:
        return json_resp([])
    if not novel:
        return json_resp([], -200)
    results = []
    for vol in Volume_Get(json.loads(novel)):
        results.append({
            "vid": vol.vid,
            "title": vol.title,
            "update_time": vol.update_time.strftime("%Y-%m-%d"),
            "chapters": json.loads(vol.chapters)
        })
    return json_resp(results)

# 已完成
@app.route("/api/novel/download/<vid>")
@app.route("/api/novel/download/<vid>/<auto>")
@Grant
@AccessControl
def download(User, vid, auto=""):
    servers = int(request.args.get("servers"))
    if not servers or int(servers) >DOWNLOAD_REMOTE_SERVER :
        servers =  DOWNLOAD_REMOTE_SERVER[0]
    else:
        servers = DOWNLOAD_REMOTE_SERVER[servers]
    if vid is None:
        return json_resp({}, -100)
    try:
        volume = Volume_Download(int(vid))
    except ValueError:
        return json_resp({}, -100)
    if not volume:
        return json_resp({}, -300)
    token = generate_token(volume[volume.rfind("/") + 1:])
    if User_Sub_Downloads(User["uid"]) == False:
        return json_resp({}, -500)
    results = {"token": token, "hash": RandomString(
        32), "name": Volume_Get_Name(vid)}
    if auto:
        return redirect("%s?token=%s&hash=%s&name=%s" % (servers + volume, token, RandomString(
            32), Volume_Get_Name(vid)))
    return json_resp(results)

# 应该已完成
@app.route("/api/novel/author/<name>")
def author(name):
    # name = request.args.get("name")
    if name is None:
        return json_resp({}, -100)
    novel = Novel_Get_Author(name)
    if not novel:
        return json_resp({}, -400)
    # results = []
    # for data in novel:
    #     results.append({
    #         "nid": data["nid"],
    #         "title": data["title"],
    #         "author": data["author"],
    #         "cover": data["cover"],
    #         "description": data["description"],
    #         "tags": data["tags"].split(","),
    #         "update_time": data["update_time"].strftime("%Y-%m-%d"),
    #         "subscribed": data["subscribed"],
    #         "viewed": data["viewed"],
    #         "liked": data["liked"],
    #         "disliked": data["disliked"],
    #         "volumes": json.loads(data["volumes"])
    #     })
    return json_resp(data)

# 已完成
@app.route("/api/novel/search/<text>")
@app.route("/api/novel/search/<text>/<inprofile>")
def search(text, inprofile=None):
    # inprofile = request.args.get("inprofile")
    # text = request.args.get("text")
    if text is None:
        return json_resp({}, -100)
    if inprofile != None and int(inprofile) == True:
        T = True
    else:
        T = False

    data = Novel_Search(text, T)
    return json_resp(data)

# 已完成
@app.route("/api/novel/subscribe/<nid>")
@Grant
@AccessControl
def subscribe(User, nid):
    uid = User["uid"]
    # nid = request.args.get("nid")
    if nid is None:
        return json_resp({}, -100)
    try:
        result = Novel_subscribe(uid=uid, nid=int(nid))
    except Exception as e:
        print(e)
        return json_resp([], -100), 500
    return json_resp([])


# @app.route("/api/comment/like")
# @Grant
# @AccessControl
# def like():
#     pass

# 已完成
@app.route("/api/comment/list")
@app.route("/api/comment/list/<int:page>")
@Grant
@AccessControl
def commit_list(User,page=1):
    nid = request.args.get("nid")
    uid = request.args.get("uid")
    if not nid and not uid:
        return json_resp({}, -100)
    else:
        if not uid:
            result = Comment_List(uid=0, nid=int(nid),page=page)
        else:
            result = Comment_List(uid=int(uid), nid=0,page=page)
    return json_resp(result)
# 已完成
@app.route("/api/comment/post",methods=['POST'])
@Grant
@AccessControl
def post_comment(User):
    uid = int(User["uid"])
    nid = int (request.args.get("nid"))
    text = request.form["text"]
    if nid is None or not text or text is None:
        return json_resp({}, -100)
    rep_cid = int(request.args.get("rep_cid"))
    if rep_cid is not None:
        result = Comment_Post(uid=int(uid), nid=int(nid),
                              text=text, rep_cid=int(rep_cid))
    else:
        result = Comment_Post(uid=int(uid), nid=int(nid), text=text)
    return json_resp(result)

# 点赞，已完成
@app.route("/api/comment/like_comment")
@Grant
@AccessControl
def like_comment(User):
    uid = User["uid"]
    cid = int(request.args.get("cid"))
    if not uid or not cid:
        return json_resp([],-100)
    else:
        try:
            return Comment_Like(cid, uid, True)
        except:
            return False

# 取消赞，已完成
@app.route("/api/comment/dislike_comment")
@Grant
@AccessControl
def dislike_comment(User):
    uid = User["uid"]
    cid = request.args.get("cid")
    if not uid or not cid:
        return json_resp([],-100)
    else:
        try:
            return Comment_like(cid, uid, False)
        except:
            return False

@app.route("/api/user/info")
@Grant
@AccessControl
def user_info(User):
    uid = request.args.get("uid")
    if not uid:
        uid = User["uid"]
    data = User_Get(int(uid))
    if data == -1:
        return json_resp({"uid":-1})
    else:
        return json_resp({
            "uid":data.uid,
            "avatar":data.avatar,
            "source_uid":data.source_uid,
            "username":data.username,
            "register_time":data.register_time,
            "qq":data.qq,
            "bio":data.bio,
            "describe":data.describe,
            "downloads":data.downloads,
            "old_driver":data.old_driver,
            "pushmail":data.pushmail})

# 修改签名，已完成
@app.route("/api/user/describe_change",methods=["POST"])
@Grant
@AccessControl
def describe_change(User):
    uid = User["uid"]
    describe = request.form["describe"]
    if not uid:
        return json_resp({}, -100)
    else:
        if User_Describe_Change(uid,describe):
            return json_resp({})
        else:
            return json_resp({},-111)




# 关注列表，已完成
@app.route("/api/user/subscribe_list")
@Grant
@AccessControl
def subscribe_list(User):
    uid = User["uid"]
    if not uid:
        return json_resp({}, -100)
    data = User_Get_Subscribe_List(int(uid))
    if not data:
        return json_resp({}, -400)
    else:
        return json_resp(data)


@app.route("/api/user/link_qq")
@Grant
@AccessControl
def link_qq(User):
    qqnumber = request.form["qqnumber"]
    pass
# 已完成
@app.route("/api/user/setting",methods=["POST"])
@Grant
@AccessControl
def setting(User):
    uid = User["uid"]
    avatar = request.form["avatar"]
    qq = request.form["qq"]
    bio = request.form["bio"]
    describe = request.form["describe"]
    pushmail = request.form["pushmail"]
    hito = request.form["hito"]
    if User_Setting(uid, qq, bio, describe, avatar, pushmail):
        return json_resp([])
    else:
        return json_resp([],-111)

@app.route("/api/user/msg_list")
@app.route("/api/user/msg_list/<int:page>")
@Grant
@AccessControl
def msg_list(User,page=1):
    uid = User["uid"]
    items =  UserMsg_List(uid=uid,page=page,raw=True)
    return items

@app.route("/api/user/msg_get")
@Grant
@AccessControl
def msg_readed(User):
    uid = int(User["uid"])
    mid = int(request.args.get("mid"))
    if UserMsg_Get(mid=mid).uid == uid:
        if UserMsg_Readed(mid=mid):
            return json_resp([])
        else:
            return json_resp([],-111)
    else:
        return json_resp([],-666)





@app.route('/api/upload', methods=['POST'])
@Grant
@AccessControl
def upload(User):
    if request.method == 'POST':
        f = request.files['file']
        upload_path = os.path.join(DL_SAVE_ADDRESS,
                                   secure_filename(f.filename))  # 注意：没有的文件夹一定要先创建，不然会提示没有该路径
        f.save(upload_path)
        return json_resp(secure_filename(f.filename))


##########################################
#@@@@          用户部分结束           @@@@@#
##########################################

##########################################
#@@@@          管理部分开始           @@@@@#
##########################################


# 已完成
@app.route("/api/admin/novel_delete")
@Grant
@AccessControl
def novel_delete(User):
    nid = request.args.get("nid")
    if User["admin"] != 1:
        return json_resp([],-666)
    if Novel_Delete(int(nid)):
        return json_resp([])
    else:
        return json_resp([],-111)

# 已完成
@app.route("/api/admin/comment_delete")
@Grant
@AccessControl
def comment_delete(User):
    nid = request.args.get("cid")
    if User["admin"] != 1:
        return json_resp([],-666)
    if Comment_Delete(int(nid)):
        return json_resp([])
    else:
        return json_resp([],-111)


@app.route('/api/admin/novel_append_volume')
@Grant
@AccessControl
def novel_append_volume(User):
    nid = request.args.get("nid")
    vid = request.args.get("vid")
    if User["admin"] != 1:
        return json_resp([],-666)
    else:
        if Novel_Append_Volume(nid=nid, vid=vid):
            return json_resp([])
        else:
            return json_resp([],-111)

@app.route("/api/admin/novel_info_change",methods=["POST"])
@Grant
@AccessControl
def novel_change_info(User):
    pass

@app.route("/api/admin/novel_create",methods=["POST"])
@Grant
@AccessControl
def novel_create(User):
    title = request.form["title"]
    author = request.form["author"]
    cover = request.form["cover"]
    description = request.form["description"]
    tags = request.form["tags"]
    if User["admin"] != 1:
        return json_resp([],-666)
    else:

        return json_resp({"nid":Novel_Create(title=title,author=author,cover=cover,tags=tags,description=description)})

@app.route("/api/admin/volume_create",methods=["POST"])
@Grant
@AccessControl
def volume_create(User):
    files = request.form["files"]
    chapters = request.form["chapters"]# must be list!!!!must!!!
    title = request.form["title"]
    if User["admin"] != 1:
        return json_resp([],-666)
    else:
        return json_resp({"vid":Volume_Create(title=title,chapters=chapters,files=files)})


@app.route("/api/admin/workist")
@Grant
@AccessControl
def workist(User):
    if User["admin"] != 1:
        return json_resp([],-666)
    page = int(request.args.get("page"))
    status = int(request.args.get("status"))
    
    if not status:
        status = 0
    if not page:
        items = Workist_List(status,raw=False)
    else:
        items = Workist_List(status=status,page=page,raw=False)
    return json_resp(items)



@app.route("/api/admin/workist_accpet")
@Grant
@AccessControl
def workist_accpet(User):
    # 同意代表着立刻执行
    if User["admin"] != 1:
        return json_resp([],-666)
    wid = int(request.args.get("wid"))
    if Workist_Do(wid=wid):
        return json_resp([])
    else:
        return json_resp([],-111)
@app.route("/api/admin/workist_refuse")
@Grant
@AccessControl
def workist_refuse(User):
    if User["admin"] != 1:
        return json_resp([],-666)
    wid = int(request.args.get("wid"))
    if Workist_Refuse(wid=wid):
        return json_resp([])
    else:
        return json_resp([],-111)

@app.route("/api/admin/workist_clean")
@Grant
@AccessControl
def workist_clean(User):
    if User["admin"] != 1:
        return json_resp([],-666)
    clean = int(request.args.get("clean"))
    if Workist_Clean(clean):
        return json_resp([])
    else:
        return json_resp([],-111)

@app.route("/api/admin/tag_create")
@Grant
@AccessControl
def tag_create(User):
    if User["admin"] != 1:
        return json_resp([],-666)
    title = request.form["title"]
    if Tag_Create(title):
        return json_resp([])
    else:
        return json_resp([],-111)

@app.route("/api/admin/tag_list")
@app.route("/api/admin/tag_list/<int:page>")
@Grant
@AccessControl
def tag_list(User,page=1):
    if User["admin"] != 1:
        return json_resp([],-666)
    tid = int(request.args.get("tid"))
    # page = request.args.get("page")
    if Tag_List(title):
        return json_resp([])
    else:
        return json_resp([],-111)

@app.route("/api/admin/tag_append")
@Grant
@AccessControl
def tag_append(User):
    if User["admin"] != 1:
        return json_resp([],-666)
    tid = int(request.args.get("tid"))
    novel = int(request.args.get("nid"))
    if Tag_Novel2Tag(tid=tid,nid=novel):
        return json_resp([])
    else:
        return json_resp([],-111)

# 发送系统私信用
@app.route("/api/admin/msg_send")
@Grant
@AccessControl
def msg_send(User):
    if User["admin"] != 1:
        return json_resp([],-666)
    uid = int(request.args.get("uid"))
    txt = request.form["txt"]
    if UserMsg_New(uid=uid,txt=txt):
        return json_resp([])
    else:
        return json_resp([],-111)




##########################################
#@@@@          管理部分结束           @@@@@#
##########################################
if __name__ == "__main__":
    app.run(debug=True)

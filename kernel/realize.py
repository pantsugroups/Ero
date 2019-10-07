# -*- coding: utf-8 -*-
import json

from playhouse.shortcuts import model_to_dict

from .database import *


def Novel_Create(title:str,author:str,cover:str,tags:str,description:str)->int:
    if not title or not author or not cover or not tags or not description:
        return -1
    else:
        try:
            return Novel.create(author=author,tags=tags,title=title,cover=cover).nid
        except:
            return -1
def Novel_Append_Volume(nid:int,vid:int)->bool:
    if not nid or not vid:
        return False
    else:
        vlist = Novel.get(Novel.nid==nid).volumes
        vlist = json.loads(vlist)
        if type(vlist)!=list:
            return False
        else:
            vlist.append(vid)
            try:
                Novel.update(vid = vlist).where(Novel.nid == nid).execute()
                return True
            except:
                return False

def Novel_Get(nid: int):
        # 获取小说信息
    if not nid or nid < 0:
        return -1
    else:
        return Novel.get(Novel.nid == nid)


def Novel_Index(page=1, raw=True,old=False):
    # 首页渲染的独立接口

    results = []
    try:
        if old == False:
            itmes = Novel.select().order_by(Novel.update_time).paginate(page, 20).where(Novel.hide==0)
        else:
            itmes = Novel.select().order_by(Novel.update_time).paginate(page, 20)
    except Exception as e:
        print (e)
        return False
    if raw:
        return itmes
    for i in itmes:
        results.append(model_to_dict(i))
    return results


def Novel_Get_Author(author: str) -> list:

    data = Novel.select().where(Novel.author == name)
    results = []
    for i in itmes:
        results.append(model_to_dict(i))
    return results


def Novel_ViewUp(nid: int) -> bool:

    try:
        Novel.update(viewed=Novel.get(Novel.nid == nid).viewed +
                     1).where(Novel.nid == nid).execute()
        return True
    except:
        return False


def Novel_Search(text: str, inprofile: bool, raw=True,oldriver=False,page=1):
    # 搜索相关
    # :inprofile: 关键字是否在简介中搜索
    if inprofile:
        itmes = Novel.select().where(Novel.title ** '%' + text + '%' and Novel.description **
                                     '%' + text + '%').paginate(page, 20)  # 这一段也许会有安全问题！
    else:
        itmes = Novel.select().where(Novel.title ** '%' + text + '%').paginate(page, 20)
    if raw:
        return itmes
    results = []
    for i in itmes:
        results.append(model_to_dict(i))
    return results


def Novel_subscribe(uid: int, nid: int) -> bool:

    if not uid or not nid:
        return False
    else:
        try:
            NovelSubscribe.create(novel=nid, user=uid)
            return True
        except:
            return False


def Novel_Get_Volume_Name(nid: int, raw=True):
    # 返回分卷id
    if not nid:
        return []
    try:
        results = Novel.get(Novel.nid == nid).volumes
    except:
        results = []
    if raw:
        return results

def Novel_Delete(nid:int) -> bool:
    if not nid or nid < 0:
        return False
    else:
        try:
            Novel.get(Novel.nid== nid).delete_instance()
            return True
        except:
            return False


def Volume_Create(nid:int,title:str,chapters:str,files:str)->int:
    if not nid or not chapters or not files or not title:
        return -1
    else:
        try:
            return Volume.create(novel=nid,title=title,chapters=chapters,files=files).vid
        except:
            return -1

def Volume_Get(nid: str, raw=True):
    if not nid:
        return []
    else:
        try:
            data = Novel.get(Novel.nid == nid).volume
        except Exception as e:
            print(e)
            return False
    if raw:
        return data
    results = []
    for i in itmes:
        results.append(model_to_dict(i))
    return results


def Volume_Get_Name(vid: int) -> str:
    if not vid:
        return ""
    try:
        return Volume.get(Volume.vid == vid).title
    except:
        return ""


def Volume_Download(vid: int) -> str:
    # 返回下载地址
    if not vid:
        return ""
    try:
        return Volume.get(Volume.vid == vid).files
    except:
        return ""





def User_Init(source_uid: int, new_name: str, qq: str):
        # 初始化这边的用户信息
    pass

# 用户本站uid获取信息
def User_Get(uid: int):
    if not uid or uid < 0:
        return -1
    else:
        return User.get(User.uid == uid)
# 更具用户主站uid获取信息
def User_Get_1(uid: int):
    if not uid or uid < 0:
        return -1
    else:
        return User.get(User.source_uid == uid)

# def User_Get_Token(uid: int) -> str:
#     if not uid:
#         raise ValueError
#     try:
#         return Token.get(Token.tid == User.get(User.uid == uid).default_token).token
#     except:
#         return "GET TOKEN ERROR."


def User_Describe_Change(uid: int,describe:str) ->bool:
    if not uid:
        return False
    try:
        User.update(describe=describe).where(User.uid == uid).execute()
        return True
    except:
        return False

def User_Sub_Downloads(uid: int) -> bool:
    if not uid:
        return False
    if User.get(User.uid == uid).downloads <= 0:
        return False
    try:
        User.update(downloads=User.get(User.uid == uid).downloads -
                    1).where(User.uid == uid).execute()
        return True
    except:
        return False


def User_Get_Subscribe_List(uid: int):
    if not uid:
        return []
    data = NovelSubscribe.select().where(NovelSubscribe.user == uid)
    results = []
    for i in data:
        results.append(model_to_dict(i)["novel"])
    return results

def User_Setting(uid:int,qq:int,bio:str,describe:str,avatar:str,pushmail:str) -> bool:
    try:
        if not pushmail:
            User.update(qq=qq,bio=bio,describe=describe,avatar=avatar,pushmail=pushmail).where(User.uid==uid).execute()
        else:
            User.update(qq=qq,bio=bio,describe=describe,avatar=avatar,pushmail=pushmail,old_driver=1).where(User.uid==uid).execute()
        return True
    except:
        return False


# def User_Set_Default_Token(uid: int, token_id: int) -> bool:
#     try:
#         User.update(default_token=token_id).where(
#             User.uid == uid).execute()
#         return True
#     except:
#         return False

# 有r_cid说明是回复楼层
def Comment_Post(uid: int, nid: int, text: str, r_cid=-1) -> bool:
    if not uid or not text or not nid:
        return False
    try:
        if not r_cid or r_cid < 0:
            Comment.create(novel=nid, user=uid, content=text)
        else:
            Comment.create(novel=nid, user=uid, content=text, rep_cid=r_cid)
        return True
    except:
        return False


def Comment_List(nid: int,uid: int, raw=False,page=1):
    '''
        要查哪个就把不要查的那个设置成0即可
    '''

    if not uid and not nid:
        return []
    results = []
    if not nid:
        data = Comment.select().where(Comment.user == uid).paginate(page, 20)
    else:
        data = Comment.select().where(Comment.novel == nid).paginate(page, 20)
    if raw:
        return data
    for i in data:
        results.append(model_to_dict(i))
    return results


def Comment_Like(cid: int, uid: int, sign: bool) -> bool:
    # :param:sigh T is like ,F is dislike
    pass

def Comment_Delete(cid:int) -> bool:
    if not cid or cid < 0:
        return False
    else:
        try:
            Comment.get(Comment.nid== nid).delete_instance()
            return True
        except:
            return False




def Tag_List(tid: int, page=1, raw=True):
    # 获取tag下的漫画
    try:
        items = NovelTag.select().where(NovelTag.tag==tid).paginate(page,20)
    except:
        return False
    if raw:
        return items
    else:
        results = []
        for i in itmes:
            results.append(model_to_dict(i))
        return results

def Tag_Create(title:str)->int:
    try:
        return Tag.create(name=title).tid
        return True
    except:
        return -1

# 小说录入
def Tag_Novel2Tag(tid:int,nid:int)->bool:
    try:
        NovelTag.create(novel=nid,tag=tid)
        Tag.select().where(Tag.tid==tag).update(count=Tag.get(Tag.tid==tid).count+1).execute()
        return True
    except:
        return False

# 获取清单
def Workist_List(status:int,raw=True,page=1):
    """
    status -1 是全部
    0则是未处理，1是已通过，2是已拒绝
    """
    if not status or not page or not raw:
        return None
    try:
        items = Workist.select().where(Workist.status==status).paginate(page, 20)
    except:
        return None
    if raw:
        return items
    else:
        results = []
        for i in itmes:
            results.append(model_to_dict(i))
        return results

def Workist_Do(wid:int)->bool:
    # 立刻执行哟
    w = Workist.get(Workist.wid == wid)
    if w.status != 0:
        return False
    else:
        action = w.action
        if action == 0:
            if Novel_Append_Volume(w.novel.nid, w.volume.vid):
                w.update(status=1).execute()
                return True
            else:
                return False
        if action ==1:
            pass

def Workist_Refuse(wid:int)->bool:
    try:
        w = Workist.select().where(Workist.wid == wid).update(status=2).execute()
        return True
    except:
        return False

def Workist_Clean(status:int)->bool:
    try:
        w = Workist.select().where(Workist.status==status).delete_instance().execute()
        return True
    except:
        return False

def UserMsg_New(uid,txt)->int:
    try:
        return UserMessage.create(user=uid,content=txt).mid
    except:
        return -1

def UserMsg_Delete(mid:int)->bool:
    try:
        UserMessage.get(UserMessage.mid == mid).delete_instance()
        return True
    except:
        return False

def UserMsg_Readed(mid:int)->bool:
    try:
        UserMessage.select().where(UserMessage.mid == mid).update(readed=1).execute()
        return True
    except:
        return False

def UserMsg_Get(mid:int,raw=False):
    try:
        UserMsg_Readed(mid)
        if raw:
            return model_to_dict(UserMessage.get(UserMessage.mid == mid))
        else:
            return UserMessage.get(UserMessage.mid == mid)
    except:
        return None

def UserMsg_List(uid:int,page=1,raw=True):
    try:
        items = UserMessage.select().where(UserMessage.user==uid).paginate(page,20)
    except:
        return False
    if raw:
        return items
    else:
        results = []
        for i in itmes:
            results.append(model_to_dict(i))
        return results

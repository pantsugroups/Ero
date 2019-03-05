# -*- coding: utf-8 -*-
import sys

sys.path.append('../')
from config import *
from peewee import *
import json
from datetime import datetime
from hashlib import md5
from app.novel import login_manager

db = SqliteDatabase(DB_PATH) if CONFIG_DEBUG else MySQLDatabase(host=MYSQL_HOST, database=MYSQL_DATABASE,
                                                                user=MYSQL_USERNAME, password=MYSQL_PASSWD, port=3306)


class BaseModel(Model):
    class Meta:
        database = db

    def __str__(self):
        r = {}
        for k in self.__data__.keys():
            try:
                r[k] = str(getattr(self, k))
            except:
                r[k] = json.dumps(getattr(self, k))
        # return str(r)
        return json.dumps(r, ensure_ascii=False)
class Novel(BaseModel):
    id = PrimaryKeyField()
    title = TextField()
    author = TextField()
    cover = TextField()
    description = TextField(null=True)
    tags = TextField(null=True)
    update_time = DateTimeField(default=datetime.now)
    subscribed = IntegerField(default=0)
    viewed = IntegerField(default=0)
    liked = IntegerField(default=0)
    ended = IntegerField(default=0)#已完结？
    volumes = TextField(default="[]")
    hide = IntegerField(default=0)# 0 为不隐藏
    # status = IntegerField(defualt=0)# 0 为未通过，1为已通过

class Volume(BaseModel):
    id = PrimaryKeyField()
    novel = ForeignKeyField(Novel, related_name="volume")
    title = TextField()
    update_time = DateTimeField(default=datetime.now)
    chapters = TextField(default="[]")
    files = TextField(default="{}")


class Tag(BaseModel):
    tid = PrimaryKeyField()
    name = TextField()
    count = IntegerField(default=0)


class NovelTag(BaseModel):
    novel = ForeignKeyField(Novel, related_name="noveltag")
    tag = ForeignKeyField(Tag, related_name="noveltag")

    class Meta:
        primary_key = CompositeKey("novel", "tag")


class User(BaseModel):
    id = PrimaryKeyField()
    avatar = TextField(null=True)# remote http address
    # source_uid = IntegerField(unique=True, null=False)
    username = TextField()
    password = TextField(null=False)
    mail = TextField(null=False)
    register_time = DateTimeField(default=datetime.now)
    qq = IntegerField(null=True)
    bio = TextField(null=True)
    # describe = TextField(null=True)
    # default_token = IntegerField(default=-1)
    downloads = IntegerField(default=250)
    pushmail = TextField(null=True)
    lv = IntegerField(default=0)# 0为未验证，1为普通，2为管理员
    hito = TextField(null=True)
    def verify_password(self, raw_password):
        return md5(raw_password) == self.password

class NovelSubscribe(BaseModel):
    novel = ForeignKeyField(Novel, related_name="novelsubscribe")
    user = ForeignKeyField(User, related_name="novelsubscribe")

    class Meta:
        primary_key = CompositeKey("novel", "user")


class Comment(BaseModel):
    cid = PrimaryKeyField()
    novel = ForeignKeyField(Novel, related_name="comment")
    user = ForeignKeyField(User, related_name="comment")
    post_time = DateTimeField(default=datetime.now)
    content = TextField()
    liked = IntegerField(default=0)
    disliked = IntegerField(default=0)
    rep_cid = IntegerField(default=-1)# 有则是楼中楼


class CommentLike(BaseModel):
    comment = ForeignKeyField(Comment, related_name="commentlike")
    user = ForeignKeyField(User, related_name="commentlike")

    class Meta:
        primary_key = CompositeKey("comment", "user")

class Workist(BaseModel):
    # 工单，记录用户的上传动作
    wid = PrimaryKeyField()
    novel = ForeignKeyField(Novel, related_name="workist")
    user = ForeignKeyField(User, related_name="workist")
    action = IntegerField(default=0) #0为上传并添加到novel,2为删除volume，3为删除小说，后面两个基本不怎么用
    volume = ForeignKeyField(Volume, related_name="workist",null=True)
    status = IntegerField(default=0)# 0则是未处理，1是已通过，2是已拒绝

class UserMessage(BaseModel):
    mid = PrimaryKeyField()
    user = ForeignKeyField(User,related_name="messages")
    content = TextField()
    time=DateTimeField(default=datetime.now)
    readed = IntegerField(default=0)#默认未阅读


@login_manager.user_loader
def load_user(user_id):
    return User.get(User.id == int(user_id))


# 建表
def create_table():
    db.connect()
    db.create_tables([User,Novel,Volume,Tag,NovelTag,NovelSubscribe,Comment,CommentLike,UserMessage])
    User.create(username="baka",password="pantsu",lv=2,mail="admin@admin.com")
    Novel.create(title="胖次群的奇妙日常",author="everybody",cover="暂时木有",ended=1,volumes="[1]")
    Volume.create(
        novel=1,
        title="你以为这是开始？其实这是结束daze",
        chapters=__import__("json").dumps(["蚊子的肛与被肛"]),
        )

if __name__ == '__main__':
    create_table()

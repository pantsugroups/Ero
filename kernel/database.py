# -*- coding: utf-8 -*-

import sys
sys.path.append('../')
from config import *

from datetime import datetime

from peewee import *


db = SqliteDatabase(DB_PATH) if CONFIG_DEBUG else MySQLDatabase(host=MYSQL_HOST, database=MYSQL_DATABASE,
                                                                user=MYSQL_USERNAME, password=MYSQL_PASSWD, port=3306)

if CONFIG_DEBUG == False:
    control_db = MySQLDatabase(host=MYSQL_HOST, database="typecho",
                               user=MYSQL_USERNAME, password=MYSQL_PASSWD, port=3306)

    class typecho_users(Model):
        class Meta:
            database = control_db
        uid = PrimaryKeyField()
        name = CharField(max_length=32)
        password = CharField(max_length=64)
        mail = CharField(max_length=200)
        url = CharField(max_length=200)
        screenName = CharField(32)
        created = IntegerField(null=False)
        activated = IntegerField(null=False)
        logged = IntegerField(null=False)
        group = CharField(max_length=16)
        authCode = CharField(max_length=64)


class BaseModel(Model):
    class Meta:
        database = db

class Novel(BaseModel):
    nid = PrimaryKeyField()
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
    vid = PrimaryKeyField()
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
    uid = PrimaryKeyField()
    avatar = TextField(null=True)# remote http address
    source_uid = IntegerField(unique=True, null=False)
    username = TextField()
    register_time = DateTimeField(default=datetime.now)
    qq = IntegerField()
    bio = TextField(null=True)
    describe = TextField(null=True)
    default_token = IntegerField(default=-1)
    downloads = IntegerField(default=250)
    old_driver = IntegerField(default=0)#1 则为是
    pushmail = TextField(null=True)
    admin = IntegerField(default=0)# 并不是用户等级，而是用户权限，
    hito = TextField(null=True)


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
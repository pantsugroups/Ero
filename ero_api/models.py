# -*- coding: utf-8 -*-
import datetime

from peewee import *

db = Proxy()


class BaseModel(Model):
    class Meta:
        database = db


class User(BaseModel):
    uid = PrimaryKeyField()
    username = CharField(50)
    password = CharField(255)
    nickname = CharField(50, null=True)
    avatar = TextField()
    email = CharField(50, null=True)
    register_time = DateTimeField(default=datetime.datetime.now)
    qq = IntegerField(null=True)
    bio = TextField(null=True)
    pushmail = TextField(null=True)
    permission = IntegerField(default=0)


class Novel(BaseModel):
    nid = PrimaryKeyField()
    title = CharField(50)
    author = CharField(10)
    cover = TextField(null=True)
    update_time = DateTimeField(default=datetime.datetime.now)
    subscribed = IntegerField(default=0)
    viewed = IntegerField(default=0)
    liked = IntegerField(default=0)
    ended = BooleanField(default=False)
    hide = BooleanField(default=False)
    subscriber = ManyToManyField(User, backref="subscribe")
UserNovelSubscribe = Novel.subscriber.get_through_model()


class Volume(BaseModel):
    vid = PrimaryKeyField()
    novel = ForeignKeyField(Novel, backref="volumes")
    title = CharField(50)
    update_time = DateTimeField(default=datetime.datetime.now)
    chapters = TextField(default="[]")


class NTag(BaseModel):
    tid = PrimaryKeyField()
    name = CharField(10)
    novels = ManyToManyField(Novel, backref="tags")
NovelTag = NTag.novels.get_through_model()


class NovelComment(BaseModel):
    cid = PrimaryKeyField()
    novel = ForeignKeyField(Novel, backref="comments")
    user = ForeignKeyField(User, backref="comments")
    post_time = DateTimeField(default=datetime.datetime.now)
    content = TextField()
    liked = IntegerField(default=0)
    disliked = IntegerField(default=0)
    to_comment = IntegerField(default=0)


class Game(BaseModel):
    gid = PrimaryKeyField()
    title = CharField(50)
    jp_title = CharField(50, null=True)
    cover = TextField(null=True)
    description = TextField()
    summary = TextField()
    screenshot = TextField(default="[]")
    publish_time = DateTimeField(default=datetime.datetime.now)
    download = TextField(null=True)


class GTag(BaseModel):
    tid = PrimaryKeyField()
    name = CharField(10)
    games = ManyToManyField(Game, backref="tags")
GameTag = GTag.games.get_through_model()

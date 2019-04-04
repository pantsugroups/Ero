# -*- coding: utf-8 -*-
import datetime

from peewee import *

db = Proxy()


class BaseModel(Model):
    class Meta:
        database = db


class Novel(BaseModel):
    nid = PrimaryKeyField()
    title = CharField(50)
    author = CharField(10)
    cover = TextField(null=True)
    tags = TextField(null=True)
    update_time = DateTimeField(default=datetime.datetime.now)
    subscribed = IntegerField(default=0)
    viewed = IntegerField(default=0)
    liked = IntegerField(default=0)
    ended = BooleanField(default=False)
    hide = BooleanField(default=False)


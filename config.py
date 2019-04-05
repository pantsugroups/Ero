# -*- coding: utf-8 -*-
from peewee import SqliteDatabase

HOSTNAME = "0.0.0.0"
PORT = 5000
DEBUG = True
DB = SqliteDatabase("ero.db")
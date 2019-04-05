# -*- coding: utf-8 -*-
from peewee import SqliteDatabase

HOSTNAME = "0.0.0.0"
PORT = 5000
DEBUG = True

DB = SqliteDatabase("ero.db")

VERSION = "0.0.1a"
SWAGGER = {
    "title": "Ero API",
    "uiversion": 2,
    "version": VERSION
}
# -*- coding: utf-8 -*-
from .. import models
from flask import Flask, render_template, request, redirect, url_for

from . import novel
@novel.route("/novel")
@novel.route("/novel/<int:page>")
def index(page=1):
    pass
@novel.route("/search/<text>")
@novel.route("/search/<text>/<inprofile>")
def search(text, inprofile=None):
    pass
@novel.route("/api/novel/detail/<int:nid>")
def detail(nid):
    pass
@novel.route("/volumes/<int:nid>")
def volumes(User,nid):
    pass
@novel.route("/author/<name>")
def author(name):
    pass
@novel.route("/subscribe/<nid>")
def subscribe(User, nid):
    pass
# -*- coding: utf-8 -*
from flask import Flask, request, render_template, jsonify, make_response
import json
from kernel import *
from functions import *


# app = Flask(__name__)


@app.route("/api/old/test")
@Grant
@AccessControl
def old_test(User):
  print(User)
  return "just a test!"


@app.route("/api/old/index")
@app.route("/api/old/index/<int:page>")
@Grant
def old_index(User, page=1):
  data = Novel_Index(page, False)
  return render_template('index.html', username=User['username'],items=data)


@app.route('/api/old/search/<text>')
@app.route('/api/old/search/<text>/<T>')
@Grant
def searchs(text, User, T=False):
  data = Novel_Search(text, T, False)
  return render_template('index.html', username=User["username"], items=data)


@app.route("/api/old/detail/<int:nid>")
@Grant
def old_detail(nid, User):
  data = Novel_Get(nid)
  app.add_template_global(Volume_Get_Name, "Volume_Get_Name")
  return render_template('book.html', nid=data.nid,
                         title=data.title,
                         author=data.author,
                         auto=RandomString(4),
                         username=User["username"],
                         cover=data.cover,
                         description=data.description,
                         tags=data.tags.replace(",", " / "),
                         update_time=data.update_time.strftime("%Y-%m-%d"),
                         subscribed=data.subscribed,
                         viewed=data.viewed,
                         liked=data.liked,
                         ended=data.ended,
                         volumes=json.loads(data.volumes))

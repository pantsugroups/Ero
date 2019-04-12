# -*- coding: utf-8 -*-
from flask import request, jsonify
from flask_restful import Resource

from ..utils import require_permission
from ..models import Game, GTag, GameTag


class GameItem(Resource):
    
    def get(self, gid):
        return jsonify({
            "status": True,
            "data": {}
        })
    
    @require_permission(2)
    def put(self, gid):
        return jsonify({
            "status": True,
            "data": {}
        })
    
    @require_permission(2)
    def delete(self, gid):
        return jsonify({
            "status": True,
            "data": {}
        })


class GameList(Resource):
    
    def get(self):
        return jsonify({
            "status": True,
            "data": {}
        })
    
    def post(self):
        return jsonify({
            "status": True,
            "data": {}
        })

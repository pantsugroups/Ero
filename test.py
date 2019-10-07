#coding:utf-8
import peewee
from config import *
from kernel import *
#Novel.create(title="この素晴らしい世界に祝福を！", author="暁なつめ", cover="/static/img/konosuba1.jpg", description="喜爱游戏的家里蹲少年佐藤和真的人生突然闭幕……但是他的眼前出现自 称女神的美少女。转生到异世界的和真就此为了满足食衣住而努力工作！原本只想安稳度日的和真，却因为带去的女神接二连三引发问题，甚至被魔王军盯上了!?",
#             tags="角川文库,异世界,冒险,搞笑", volumes='[1]')
nid = Novel_Create(title="この素晴らしい世界に祝福を！", author="暁なつめ", cover="/static/img/konosuba1.jpg",description="喜爱游戏的家", tags="角川文库,异世界,冒险,搞笑")
#Volume.create(novel=1, title="第一卷 啊啊，没用的女神大人 ",
#              chapters='["第一章 这个自称女神和异世界转生！","第二章 这个右手中握着的财宝！","第三章 这个湖中自称女神",“第四章 这场毫不轻松的战斗的终结!”,“终章”]', files='http://127.0.0.1:5001/1.txt')
vid = Volume_Create(nid = nid,title="第一卷 啊啊，没用的女神大人 ", chapters='["第一章 这个自称女神和异世界转生！","第二章 这个右手中握着的财宝！","第三章 这个湖中自称女神",“第四章 这场毫不轻松的战斗的终结!”,“终章”]', files='http://127.0.0.1:5001/1.txt')
Novel_Append_Volume(nid, vid)
print(nid,vid)
# User.create(source_uid=1, username="BIE的测试号", qq=123456789)
# Tag.create(name="角川文库", count=1)
# Tag.create(name="异世界", count=1)
# Tag.create(name="冒险", count=1)
# Tag.create(name="搞笑", count=1)
print(Novel_Get(1).title)

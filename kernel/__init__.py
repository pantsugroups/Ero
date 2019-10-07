# -*- coding: utf-8 -*-

from .database import *
from .realize import *

Novel.create_table()
User.create_table()
Volume.create_table()
Tag.create_table()
NovelTag.create_table()
NovelSubscribe.create_table()
Comment.create_table()
CommentLike.create_table()

if CONFIG_DEBUG == False:
        # Do Something...
    pass

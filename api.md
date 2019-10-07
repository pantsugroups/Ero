# API 规范
开发中......BIE的玄学文档

规范就是没有规范

# 响应结构
响应内容为JSON，结构如下

|Name|Type|Description|
|:-:|:-:|-|
|code|Integer|响应码，正常响应为0，若发生错误则不为0，具体响应码见下表|
|msg|String|具体错误信息，正常响应则无此字段|
|data|Dict|接口返回的数据|

响应码如下(待补充)

|Code|Description|
|:-:|-|
|0|正常响应|
|-100|请求参数错误|
|-200|nid错误|
|-999|未登录|
|-500|下载点数不足|
|-111|内部错误，一般实在realize部分出错，和业务层无关|
|-101|没有权限浏览老司机模式|
|-666|没有权限执行，越权执行了一些东西|
响应示例

```javascript
{
    "code": -100,
    "msg": "请求参数错误",
    "data": {...}
}
```


## API列表-用户层

### /api/novel/index/<page>

#### 接口简介
一般用于获取最新的数据，每页20条数据

#### 请求方法
Url传参

必须登陆
#### 请求参数
int:page => 页面page，可空。默认二十条

#### 响应内容:数组之一
|Name|Type|Description|
|:-:|:-:|-|
|author|String|作者|
|cover|String|封面Url|
|description|String|小说简介|
|ended|Int|这是啥我忘了|
|liked|Int|点赞人数，也许会废弃|
|subscribed|Int|订阅人数|
|tags|String|tags列表，已逗号为分割|
|title|Stromg|小说标题|
|update_time|Date|最后更新时间|
|viewed|Int|浏览人数|
|volumes|String[]|小说卷id，使用/volumes查询|

样例

```javascript
{
  "code": 0, 
  "data": {
    "data": [
      {
        "author": "暁なつめ", 
        "cover": "/static/img/konosuba1.jpg", 
        "description": "喜爱游戏的家里蹲少年佐藤和真的人生突然闭幕……但是他的眼前出现自 称女神的美少女。转生到异世界的和真就此为了满足食衣住而努力工作！原本只想安稳度日的和真，却因为带去的女神接二连三引发问题，甚至被魔王军盯上了!?", 
        "ended": 0, 
        "liked": 0, 
        "nid": 1, 
        "subscribed": 0, 
        "tags": "角川文库,异世界,冒险,搞笑", 
        "title": "この素晴らしい世界に祝福を！", 
        "update_time": "Tue, 16 Oct 2018 19:49:50 GMT", 
        "viewed": 0, 
        "volumes": "[1]"
      }
    ]
  }
}
```



### /api/novel/detail/<nid>

#### 接口简介
获取轻小说详细信息

#### 请求方法
URL传参

#### 请求参数
int:pid => 文章pid

#### 响应内容
|Name|Type|Description|
|:-:|:-:|-|
|nid|Integer|轻小说ID|
|title|String|轻小说书名|
|author|String|作者|
|cover|String|封面url|
|description|String|简介|
|tags|List|标签|
|update_time|String|更新时间，格式 %Y-%m-%d|
|subscribed|Integer|订阅数|
|viewed|Integer|浏览数|
|liked|Integer|赞数|
|disliked|Integer|踩数|
|volumes|List|小说分卷，按从旧到新排序|

样例

```javascript
{
    "code": 0,
    "data":{
        "nid": 1,
        "title": "苍海的少女们",
        "author": "白鳥士郎",
        "cover": "cover/abdn2h9snkc9dh3bs83h4j82yx03jc6z.jpg",
        "description": "「从现在开始，你就是女人了。」\n「哎？哎…哎哎哎哎？」\n遭到祖国的追杀而落海漂流的皇子修芬，被亚拉米斯的军舰所救。但是，这艘全部由少女构成的船上有一条铁的纪律——「禁止男性」。\n在这种一旦暴露就会被判处死刑的状况下，船长的提议居然是：用你与生俱来的可爱扮成女孩子吧！在少女们百般的呵护与骚扰之下，修芬能守住自己的秘密吗？\n男女比例1:200的船上所上演的 Boy meets Girls。现在开幕！",
        "tags": [
            "搞笑",
            "后宫",
            "伪娘"
        ],
        "update_time": "2013-2-25",
        "subscribed": 23053,
        "viewed": 297204,
        "liked": 2395,
        "disliked": 58,
        "volumes": [
            1,
            2,
            3,
            4
        ]
    }
}
```

### /api/novel/volumes/<vid>

#### 接口简介
获取每卷的信息,参数从上面detial接口的volumes获取

#### 请求方法
Url传参

#### 请求参数
int:nid => 小说id

#### 响应内容

|Name|Type|Description|
|:-:|:-:|-|
|vid|Int|卷id|
|title|String|卷标题|
|update_time|String|更新时间|

测试样例

```javascript
{
  "code": 0, 
  "data": {
    "data": [
      {
        "vid":1,
        "title": "第一卷 啊啊，没用的女神大人",
        "chapters": [
            "第一章 这个自称女神和异世界转生！",
            "第二章 这个右手中握着的财宝！",
            "第三章 这个湖中自称女神",
            "第四章 这场毫不轻松的战斗的终结!",
            "终章"
            ],
        "update_time": "Tue, 16 Oct 2018 19:49:50 GMT"
      }
    ]
  }
}
```

### /api/novel/download/<vid>/<String:auto>

#### 接口简介

获取下载地址

#### 请求方法

Url传参

#### 请求参数

int:vid  [URL] => 卷id

可选参数:String:auto [URL] => 自动跳转下载地址

可选参数: int:servers  => 下载服务器地址，对应config.py内容

#### 响应内容

|Name|Type|Description|
|:-:|:-:|-|
|token|String|总而言之就是token|
|hash|String|别看写着hash其实是随机字符串|
|name|String|小说保存到本地时的名字|

如果有auto参数，则直接302跳转到小说的下载页面

反之则返回上面的数据

测试样例

```javascript
{
  "code": 0, 
  "data": {
    "data": [
      {
        "token":"balablabalabaskhaksdjaskdkasd",
        "hash": "balabalabalabalabaa",
        "name": "第一卷 啊啊，没用的女神大人.txt"
      }
    ]
  }
}
```

### /api/novel/author/<name>

### 接口简介

根据作者名字搜索漫画

#### 请求方法

Url传参

#### 请求参数
string:name  => 作者名字

#### 响应内容

咕咕咕还没实现



### /api/novel/search/<text>/<inprofile>

### 接口简介

搜索漫画

#### 请求方法

Url传参

#### 请求参数
string:text  => 关键词
可选参数： string:inprofile => 搜索时是否包含简介

#### 响应内容

和Index结果一样

### /api/novel/subscribe/<nid>

### 接口简介

订阅小说

#### 请求方法

Url传参

#### 请求参数
int:nid  => 小说id

#### 响应内容

正常返回

### /api/comment/list/<int:page>

#### 接口简介

获取评论列表

#### 请求方式
url传参

get参数

#### 请求参数

int:nid or int:uid => 二选一，要查看某小说下的评论就传入nid，查看用户的评论就传入uid。

int:page => 翻页，默认20条，可空

### /api/comment/post

#### 接口简介

发送评论

#### 请求方式

get

post

#### 请求参数

int:nid  [GET] =>  小说nid

str:text [POST] => 发送的评论

int:rep_cid [GET] => 回复楼层，cid，可空

#### 返回内容

返回小说cid


### /api/comment/like_comment
### /api/comment/dislike_comment
#### 接口简介

点赞 or 取消赞

#### 请求方式

get

#### 请求参数

int:cid => 评论cid

#### 返回内容

依旧是普通返回

### /api/user/subscribe_list

#### 接口简介

获取关注列表

#### 请求方式

get

#### 请求参数

无参数，必须登陆

#### 返回内容

和index一样

### /api/user/setting

#### 接口信息

修改用户信息

#### 请求方式

POST

#### 请求参数
str:avatar => 头像url

str:qq => qq number

str:bio => 性别

str:describe => 签名

str:pushmail => 推送地址

str:hito => 一句话

#### 返回信息

标准返回

### /api/user/msg_list/<int:page>

#### 接口信息

获取消息盒子

#### 请求方式

get

#### 请求参数
无，必须登陆

#### 返回内容

数组，元素之一

|Name|Type|Description|
|:-:|:-:|-|
|mid|int|message id|
|user|int|uid|
|time|str|时间咯|
|readed|int|是否已读|

### /api/user/msg_get

#### 接口简介

获取内容，自动标记为以读

#### 请求方式
get参数

#### 请求参数
int:mid => 消息id

#### 返回内容

和上边一样

### /api/admin/volume_create
### 接口简介
添加卷信息到数据库
### 请求方式
post请求
#### 请求参数
str:title -> 卷标题
str:chapters -> 卷内章节，一定要是list类型！
str:files -> 小说所在的上传路径，已upload接口返回为准
#### 返回数据
卷vid

## API列表-管理层

### /api/admin/novel_delete 
#### 接口简介
删除小说（其实并不会删除

#### 请求方式
get传参

#### 请求参数
int:nid => 小说id

#### 返回数据
标准返回

### /api/admin/comment_delete
#### 接口简介
删除评论（其实并没有

#### 请求方式
get传参

#### 请求参数
int:cid => 评论cid
#### 返回数据
标准返回

### /api/admin/novel_append_volume
#### 接口简介
添加卷到小说

#### 请求方式
get请求

### 请求参数

int:nid => 小说novel id

int:vid => 卷volume id

### 返回数据

标准返回。。。吧？


### /api/admin/novel_create

#### 接口简介
添加小说

#### 请求方式
POST上传

#### 请求接口
str:title => 小说标题

str:author => 小说作者

str:cover => 小说封面（不带网址）

str:description => 小说简介

str[]:tags => 小说tag（最好为空，不然的话就得list(json格式)）

#### 返回数据

应该也是标准返回？

### /api/admin/workist
#### 接口简介

列出工单列表

#### 请求方式
get方式
#### 请求参数
int:page => 页数

int:status => 列出status的工单（0则是未处理，1是已通过，2是已拒绝）

#### 返回数据
数组的元素之一，内容如下

|Name|Type|Description|
|:-:|:-:|-|
|wid|int|workist id|
|novel|int|nid|
|user|int|uid|
|action|int|0为上传并添加到novel,2为删除volume，3为删除小说，后面两个基本不怎么用|
|volume|int|vid|
|status|int|状态：0则是未处理，1是已通过，2是已拒绝|

### /api/admin/workist_accpet

#### 接口简介
接受工单中的任务，代表着立刻执行workist中action的行为

#### 请求方式
get参数
#### 请求参数
Int:wid => 工单id

#### 返回内容
标准返回

### /api/admin/workist_refuse
#### 接口简介
拒绝工单

#### 请求方式
get参数
#### 请求参数
Int:wid => 工单id

#### 返回内容
标准返回

### /api/admin/workist_clean

#### 接口简介
清理已经处理过的
#### 请求方式
get参数
#### 请求参数
int:clean => workist的status类型

#### 返回内容
标准返回

### /api/admin/msg_send
#### 接口简介
发送用户信息
#### 请求方式
get参数
post请求
#### 请求参数
int:uid => user id

str:txt [POST] => 要发送的消息
#### 返回内容
标准返回

### /api/admin/tag_create
#### 接口简介
创建tag
#### 请求方式
POST参数
#### 请求参数
str:title [POST] => 创建新类型的名称
#### 返回内容
标准返回

### /api/admin/tag_list/<int:page>
#### 接口简介
获取该类型下的小说
#### 请求方式
get请求
#### 请求参数
int:tid => tid

int:page [url参数] => 页面数，每页20，可空
#### 返回数据
同index接口

### /api/admin/tag_append
#### 接口简介
把小说添加到指定类型下
#### 请求方式
get传参
#### 请求参数
int:tid => 要添加到的类型tid

int:novel => 小说nid
#### 返回数据
标准返回

### /uploads
#### 接口简介
上传小说
#### 请求方式
POST请求
#### 请求参数
byte:file => 要上传的小说

#### 返回数据
返回小说网址（自行拼接DL服务器）
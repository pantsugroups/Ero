# Ero使用手册

## EroLight
仅针对APP目录下的Novel包

全局申明：

  - nid:novel identify
  - cid:comment idenfity
  - vid:volume identify

如果无特定说明，返回格式则都是

返回样例：
```json
{"code": 0, "msg": "成功。" }
```
### 管理员类
此类下面必须要求有管理员权限
#### admin/novel_delete/<nid>
请求方式：GET

删除小说，要求nid。


#### admin/comment_delete/<cid>
请求方式：GET

删除评论，要求cid。


#### admin/novel_description_change/<nid>
请求方式：POST

修改小说简介，要求nid


#### admin/novel_append_volume/<nid>
请求方式：POST，参数vid。

给小说添加卷，要求nid。



#### admin/novel_create
请求方式：POST

参数：
  - title 小说标题
  - author 小说作者
  - cover  小说封面
  - description 小说简介
  - tags 小说tag，用/隔开



#### admin/volume_create
请求方式：POST

参数：
  - title 分卷标题
  - chapters  分卷中章节标题
  - files 分卷路径，考虑到多节点，请用绝对路径

添加卷到小说中。
请使用 stream/uplod_volume 方式上传小说
  

### 验证类
登陆和注销在这里

#### auth/login
请求方式：POST

参数：
   - user 用户名
   - passwd  密码

只有普通的md5加密。



#### auth/register
请求方式：POST

参数：
   - user 用户名
   - passwd  密码
   - mail 邮箱

请在app/novel/conf/config.py下配置smtp信息

注册之后将会发送一封验证邮件到mail（还未实现）


#### auth/logout
请求方式：GET

注销登陆


### 评论
小说评论相关
#### comment/list/<cid>/<page>
请求方式:GET

获取小说评论。要求cid,page。

page为页数，每页请求20个。

#### comment/post/<nid>
请求方式：POST
参数：
 - content 评论内容
 - rep_cid **GET参数**，如果有则是要回复楼层的cid 

发送评论，要求nid


#### comment/like_comment/<cid>
请求方式：GET

点赞评论，要求cid

#### comment/dislike_comment/<cid>
请求方式：GET

踩评论，要求cid

### 小说类
观看小说咯
#### novel/<page>
请求方式：GET

小说列表，要求参数page，每页20个。

返回样例：
```json
    {
        "code": 0,
        "msg": "成功。",
        "data": {"novels": {}}
    }
```
#### novel/detail/<nid>
请求方式：GET

查看小说，要求nid。

返回样例：
```json
    {
        "code": 0,
        "msg": "成功。",
        "data": {"novel": {}}
    }
```
#### novel/search/<text>/<page>
请求方式：GET

搜索小说，要求text，page，分别是关键词和页数

返回结果和novel/结果一样
#### novel/volumes/<nid>
请求方式：GET

查看小说分卷，要求nid

返回样例：
```json
  {
    "code": 0,
     "msg": "成功",
     "data":{
        "volumes":
        [{"vid":"xxx","title":"xxx","chapters":"xxx","update_time":"xxx"}]}}
```
#### novel/author/<name>/<page>
请求方式get

搜索作者的书籍，要求name，page。

返回结果和novel一样。

#### novel/subscribe/<nid>
请求方式GET

订阅小说。

### 文件类

### 用户类
#### user/info/<uid>
请求方式：GET

查看自己/他人的信息，有nid则是他人。

返回样例：
```json
{
        "code": 0,
        "msg": "成功。",
        "data": {"user": {"avatar":"xxxx",
                          "username":"xxx",
                          "mail":"xxx",
                          "register_time":"xxx",
                          "qq":"xxx",
                          "bio":"xxx",
                          "downloads":"xxx",
                          "pushmail":"xxx",
                          "lv":"xxx",
                          "hito":"xxx",
                          
                          }}
    }
```

#### user/describe_change
请求参数：POST

参数： BIO  签名

修改签名

#### user/subscribe_list/<uid>
请求方式GET

查看自己或他人的关注列表

返回结果和NOVEL/一样

#### user/change_avatar
请求方式：POST

参数：avatar 头像

修改头像


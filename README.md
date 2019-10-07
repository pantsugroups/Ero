# 日志更新表

# 2019-10-7

第三版api已经完成了，现在回来看看这个api，不忍直视.jpg

## 12-21

### 已完成
全部API基本完成

### 还差
emmmm还差子DL下载点的同步和<del>节点切换</del>

节点切换也已经我完成 

API内还有很多问题，例如防CC，还有上传大小验证等等

最大最大最大的问题就是，还没进行测试！
## 12-17

### 已完成
基本用户级和管理级别的API已经完成，但是还*没测试！！！！！*

### 还差
上传问题

DL下载服务切换

主下载服务和子下载服务同步问题

邮件队列

# ero_novel - BIE

## 全局文档
详情：[点我跳转](https://docs.qq.com/doc/DRlNMT1JYSWRoR1hQ?tdsourcetag=s_pctim_aiomsg&ADUIN=1948158539&ADSESSION=1537277825&ADTAG=CLIENT.QQ.5579_.0&ADPUBNO=26833&autoclear=1)
#### 项目介绍
ero.ink下属的轻小说站点，请使用Python3
BIE的玄学分支，只保证能用and不崩溃，其他随缘
#### 开发语言
Python3

#### 主要依赖

- Flask
- Peewee

#### 注意事项

- 用户系统已经开始编写
- 开发阶段使用SQLite3
- 我才不用ide
- 请务必先看权限管理

#### 开发人员

- 天台
- 9bie

#### 结构

- kernel/          数据库操作接口，包含功能的实现
- *.md             说明文档
- web_api          api接口
- web_view         jinji2渲染接口
- functions        脚手架,包含各种工具函数
- config           配置文件
- manage           入口点 

#### API规范
见 [api.md](api.md)

#### 权限管理（必看）
见 [Access.md](Access.md)

#### TODO
- 评论相关    ---  kernel/realize.py
- 权限验证    ---  function.py:Grant

#### 测试方法

首先先安装python3以及相关依赖

```
apt-get install python3
apt-get install python3-pip
python3 -m pip install flask
python3 -m pip install peewee
```

之后添加测试数据

`python3 test.py`

之后启动即可

`python3 manage.py`

相关API请查看API列表 见 [web_api.py](web_api.py)

下载相关请看: [DL.md](DL.md)

#### CONFIG
详细轻阅读[config.py](config.py)
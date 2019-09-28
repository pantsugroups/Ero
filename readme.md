# Ero-Go
这是一个Ero的后端，使用Go语言重构的版本
```
eroauz
  ├─api         // 外部接口部分
  ├─conf        // 配置文件
  ├─middleware  // 中间件
  ├─models      // 数据库对象
  ├─serializer  // 反序列化
  ├─server      // 整体服务
  ├─service     //  业务层
  │  ├─archive
  │  ├─novel
  │  └─tag
  └─utils      // 工具类
```

## 运行机制
`Server->API->Service->Models->Serializer->API->Server`

## 路由
详细API请查看：[API详细调用说明](api.md)
#### 普通 （无需登陆）
  
  - 用户登录(POST)：/user/login
  - 用户注册(POST)：/user/register
  - 文章浏览(GET)：/archive/:id
  
#### R层 （需要登陆）

  - 用户信息(GET)：/user/:id

#### A层 （需要登陆且所有者是自己或是管理员）

  - 发表文章(POST)：/archive/
  - 删除文章(DELETE): /archive/:id
  - 更新文章(PUT): /archive/:id
  
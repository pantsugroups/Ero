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
  │  ├─tag
  │  └─ ...
  └─utils      // 工具类
```

## 运行机制
`Server->API->Service->Models->Serializer->API->Server`

## 路由
详细API请查看： `/swagger/index.html`

或者可以访问程序目录下面`routes.json`查看

或者使用管理员访问`/routes.json`查看
 
## 注意事项

编译前请先打开`conf/config.go`配置好相关参数再运行。

## 历史

第一版本（群魔乱舞版）：[第一版](https://github.com/pantsugroups/Ero/tree/first)

第二版本（勉强能看且能用）：[第二版](https://github.com/pantsugroups/Ero/tree/master)
# AccessControl and UserInfomaion

## 权限管理接口
权限管理接口位于funtions.py下

#### 接口列表
- Grant
- AccessControl

#### 调用方式
```python
  @app.route("/")
  @Grant
  @AccessControl
  def Index(User):pass
```
如果只需要获取用户信息，不需要进行拦截，直接使用Grant装饰器即可

请先调用flask的路由函数，再调用Grant函数进行用户登录的鉴权

之后要对必须登录的函数或者对一些要求进行条件判断(如管理员等等)的函数使用AccessControl进行判断

#### 函数内用户信息的获取

请务必在调用Grant装饰器的函数参数中加入User参数

Grant将会吧用户信息已传惨的形式发送给函数，类型为dirct

```javascript
{
    "logined": True,
    "username": "xxxx",
    "data": {...}
}
```

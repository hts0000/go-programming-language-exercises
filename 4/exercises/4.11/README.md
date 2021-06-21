### http请求方法
- HEAD：请求HTTPD标头信息
- GET：查询信息
- POST：用于创建资源
- PATCH：用于通过部分json数据更新资源
- PUT：用于替换资源或集合
- DELETE：用于删除资源

### 身份验证

GitHub REST API验证有两种方式。
需要身份验证的请求有时将返回`404 Not Found`，而不是`403 Forbidden`。这是为了防止私有仓库意外泄露给未经授权的用户。
基本验证：
```
$ curl -u "username" https://api.github.com
```

OAuth2令牌验证：（建议使用）
```
$ curl -H "Authorization: token OAUTH-TOKEN" https://api.github.com
```

### ISSUES相关操作
https://docs.github.com/cn/rest/reference/issues
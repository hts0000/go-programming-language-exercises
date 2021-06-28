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

### github API
未经身份验证的客户端每小时可以发出 60 个请求。 要每小时发出更多请求，我们需要进行身份验证。 事实上，使用 GitHub API 做任何有意义的事情需要身份验证。

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
Authorization: token OAUTH-TOKEN
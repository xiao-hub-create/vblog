# Web全栈开发(Vblog)


## 软件设计

### 需求

管理markdown个文字的一个网站，作者后台发布文章，访客前台浏览查看文章


### 流程

![](./docs/flow.drawio)

### 产品原型

https://gitee.com/infraboard/go-course/blob/master/new.md#%E6%9E%B6%E6%9E%84%E8%AE%BE%E8%AE%A1

![](./docs/page.png)

### 架构(BS)和概要设计

![](./docs/arch.png)

### 业务的详细设计

直接使用Go的接口 来定义业务
```go
// 业务域
type Service interface {
    UserService
    InnterService
}

// 1. 外部
type UserService interface {
    // 颁发令牌 登录
    IssueToken(context.Context, *IssueTokenRequest) (*Token, error)
    // 撤销令牌 退出
    RevolkToken(context.Context, *RevolkTokenRequest) (*Token, error)
}

type RevolkTokenRequest struct {
    // 访问令牌
    AccessToken string `json:"access_token"`
    // 刷新令牌, 构成一对，避免AccessToken 泄露，用户可以直接 revolk
    RefreshToken string `json:"refresh_token"`
}

type IssueTokenRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
    // 记住我, Token可能1天过期, 过去时间调整为7天
    RememberMe bool `json:"remember_me"`
}

// 内部
type InnterService interface {
    // 令牌校验
    ValidateToken(context.Context, *ValidateTokenRequest) (*Token, error)
}

type ValidateTokenRequest struct {
    // 访问令牌
    AccessToken string `json:"access_token"`
}
```

数据库的设计伴随接口设计已经完成

1. 如何基于Vscode 构造单元测试的配置

```json
{
    "go.testEnvFile": "${workspaceFolder}/etc/test.env",
}
```

添加工作目录环境变量
```
WORKSPACE_DIR="/Users/xxxx/Projects/go-course/go17/vblog"
```

### 业务模块的实现

TDD (Test Drive Development)

+ 用户模块

```go
// 我要测试的对象是什么?, 这个服务的具体实现
// Service的具体实现？现在还没实现
// $2a$10$yHVSVuyIpTrQxwiuZUwSMuaJFsnd4YBd6hgA.31xNzuyTu4voD/QW
// $2a$10$fe0lsMhM15i4cjHmWudroOOIIBR27Nb7vwrigwK.9PhWdFld44Yze
// $2a$10$RoR0qK37vfc7pddPV0mpU.nN15Lv8745A40MkCJLe47Q00Ag83Qru
// https://gitee.com/infraboard/go-course/blob/master/day09/go-hash.md#bcrypt
func TestRegistry(t *testing.T) {
    req := user.NewRegistryRequest()
    req.Username = "test02"
    req.Password = "123456"
    ins, err := impl.UserService.Registry(ctx, req)
    if err != nil {
        t.Fatal(err)
    }
    t.Log(ins)
}

func TestDescribeUser(t *testing.T) {
    ins, err := impl.UserService.DescribeUser(ctx, &user.DescribeUserRequest{
        user.DESCRIBE_BY_USERNAME, "admin",
    })
    if err != nil {
        t.Fatal(err)
    }
    //
    // if ins.Password = in.Password
    t.Log(ins.CheckPassword("1234567"))
}
```

```go
var UserService user.Service = &UserServiceImpl{}

// 定义一个struct, 用于实现 UserService就是刚才定义的接口
// 怎么才能判断这个结构体没有实现这个接口
type UserServiceImpl struct {
}

// DescribeUser implements user.Service.
func (u *UserServiceImpl) DescribeUser(ctx context.Context, in *user.DescribeUserRequest) (*user.User, error) {
    query := datasource.DBFromCtx(ctx)
    switch in.DescribeBy {
    case user.DESCRIBE_BY_ID:
        query = query.Where("id = ?", in.Value)
    case user.DESCRIBE_BY_USERNAME:
        query = query.Where("username = ?", in.Value)
    }

    ins := &user.User{}
    if err := query.Take(ins).Error; err != nil {
        return nil, err
    }
    return ins, nil
}

// Registry implements user.Service.
func (u *UserServiceImpl) Registry(ctx context.Context, in *user.RegistryRequest) (*user.User, error) {
    ins, err := user.New(in)
    if err != nil {
        return nil, err
    }

    // 明文密码保持到数据库，是不安全
    // 对称加密/非对称， 解密
    // 消息摘要, 无法还原
    // 怎么知道用户的密码 比对hash  123 -> (xxx)
    // md5 sha1/256/512, hmac, ...
    // 结果固定
    hashPass, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }
    ins.Password = string(hashPass)

    if err := datasource.DBFromCtx(ctx).Create(ins).Error; err != nil {
        return nil, err
    }

    // context.WithValue()

    // 无事务的模式
    // datasource.DB().Transaction(func(tx *gorm.DB) error {
    // 	ctx := datasource.WithTransactionCtx(ctx)
    // 	// 1.
    // 	svcA.Call(ctx)
    // 	// 2.
    // 	svcB.Call(ctx)
    // 	// 3.
    // 	svcC.Call(ctx)
    // })

    return ins, nil
}
```

### API接口层

Gin 
```go
package api

import (
    "github.com/gin-gonic/gin"
    "github.com/infraboard/mcube/v2/http/gin/response"
    "gitlab.com/go-course-project/go17/vblog/apps/token"
    "gitlab.com/go-course-project/go17/vblog/apps/token/impl"
)

func NewTokenApiHandler() *TokenApiHandler {
    return &TokenApiHandler{
        token: impl.TokenService,
    }
}

type TokenApiHandler struct {
    // 业务控制器
    token token.UserService
}

// 提供注册功能, 提供一个Group
// book := server.Group("/api/tokens")
func (h *TokenApiHandler) Registry(r *gin.Engine) {
    router := r.Group("/api/tokens")
    router.POST("/issue", h.IssueToken)
    router.POST("/revolk", h.RevolkToken)
}

func (h *TokenApiHandler) IssueToken(ctx *gin.Context) {
    in := token.NewIssueTokenRequest("", "")
    if err := ctx.BindJSON(in); err != nil {
        response.Failed(ctx, err)
        return
    }
    ins, err := h.token.IssueToken(ctx.Request.Context(), in)
    if err != nil {
        response.Failed(ctx, err)
        return
    }
    response.Success(ctx, ins)
}

func (h *TokenApiHandler) RevolkToken(ctx *gin.Context) {
    in := &token.RevolkTokenRequest{}
    if err := ctx.BindJSON(in); err != nil {
        response.Failed(ctx, err)
        return
    }
    ins, err := h.token.RevolkToken(ctx, in)
    if err != nil {
        response.Failed(ctx, err)
        return
    }
    response.Success(ctx, ins)
}
```

### 组装程序(v1)

怎么做开发显得专业

```go
package main

import (
    "log"

    "github.com/gin-gonic/gin"
    "github.com/infraboard/mcube/v2/ioc/config/http"
    blogApi "gitlab.com/go-course-project/go17/vblog/apps/blog/api"
    tokenApi "gitlab.com/go-course-project/go17/vblog/apps/token/api"
    "gitlab.com/go-course-project/go17/vblog/config"
)

func main() {
    config.LoadConfig()

    // gin Engine, 它包装了http server
    server := gin.Default()

    // 注册业务模块的路有
    tokenApi.NewTokenApiHandler().Registry(server)
    blogApi.NewBlogApiHandler().Registry(server)

    // 服务器启动
    if err := server.Run(http.Get().Addr()); err != nil {
        log.Println(err)
    }
}
```

功能分层架构(MVC): Book Api
业务分区架构(DDD): Vblog

### 测试验证

```sh
curl --location 'http://127.0.0.1:8080/vblog/api/v1/blogs' \
--header 'Content-Type: application/json' \
--data '{
    "title": "POSTMAN测试01",
    "author": "will",
    "content": "post man 测试",
    "summary": "Go全栈项目",
    "category": "软件开发"
}'
```

```sh
curl --location 'http://127.0.0.1:8080/vblog/api/v1/tokens' \
--header 'Content-Type: application/json' \
--data '{
    "username": "admin",
    "password": "123456"
}'
```

### 中间件鉴权

![](./docs/middleware.png)

开发一个认证中间件: 用于根据用户携带的Token信息，判断用户身份，并把用户身份信息方到上下文中，传递给后面HandleFunc中使用


## ioc优化

+ 会用: mcube ioc / golang-ioc
+ 掌握原理: 自己造，手写一个简单

### 问题

+ 手动管理: main, 自己组装对象, 业务越复杂，组装难度越高

![](./docs/oop.png)

+ ioc: 引入了一个中间层, 这个中间层负责对象的管理, 自己对象自己去ioc获取依赖，而不是我们开发者 把依赖传递给他，完成 对象的 依赖 由被动 变成主动, ioc, 依赖倒置

![](./docs/ioc.png)

### 基于mcube ioc来改造

![alt text](image.png)

https://www.mcube.top/docs/framework/

1. 对象注册到ioc: 把我们的对象实现了1个IOC对象(符合Ioc接口定义的对象), 可以通过继承基础类，直接实现接口ObjectImpl
    + 2 APIHandler: TokenApiHandler, BlogApiHandler
    + 3 Controller: UserServiceImpl, TokenServiceImpl, BlogServiceImpl

对象注册
```go
func init() {
    ioc.Controller().Registry(&TokenServiceImpl{})
}

// 定义一个struct, 用于实现 UserService就是刚才定义的接口
// 怎么才能判断这个结构体没有实现这个接口
type TokenServiceImpl struct {
    ioc.ObjectImpl

    // user service
    user user.AdminService
}

func (*TokenServiceImpl) Name() string {
    return token.AppName
}

// 他需要自己去获取依赖，通过ioc
func (i *TokenServiceImpl) Init() error {
    i.user = user.GetService()
    return nil
}
```

对象获取
```go
func GetService() Service {
    return ioc.Controller().Get(AppName).(Service)
}
```


启动的时候 只需要启动 HTTP Sever就可以啦
```go
func main() {
    config.LoadConfig()

    // 服务器启动
    if err := server.GinServer.Run(http.Get().Addr()); err != nil {
        log.Println(err)
    }
}
```


修改为使用cli来用
+ cobra 的 OnInitialize 进行ioc配置的读取，与加载
+ start 命令 调用 server.run 来进行 服务启动


### ioc 托管Gin框架

```go
// module_name
func (h *BlogApiHandler) Name() string {
    return "blogs"
}

// router := r.Group("/vblog/api/v1/blogs")
// ioc_gin.ObjectRouter(h)
// 模块的名称, 会作为路径的一部分比如: /mcube_service/api/v1/hello_module/
// 路径构成规则 <service_name>/<path_prefix>/<service_version>/<module_name>
func (h *BlogApiHandler) Init() error {
    h.blog = blog.GetService()
    // 在ioc获取gin server *gin.Engine
    ioc_gin.RootRouter()

    // 获取模块路有: url前缀,
    r := ioc_gin.ObjectRouter(h)
    r.GET("", h.QueryBlog)

    r.Use(middleware.Auth)
    r.POST("", h.CreateBlog)
    return nil
}
```


```go
func main() {
    // load config and run with cobra cli
    // 使用ioc的 gin来进行路由加载
    cmd.Start()
}
```

### 测试验证

```sh
curl --location 'http://127.0.0.1:8080/api/vblog/v1/blogs' \
--header 'Content-Type: application/json' --header 'Authorization: Bearer eaabbe42-9bc0-4a8e-bd86-c92542d062bb' \
--data '{
    "title": "POSTMAN测试01",
    "author": "will",
    "content": "post man 测试",
    "summary": "Go全栈项目",
    "category": "软件开发"
}'
```

```sh
curl --location 'http://127.0.0.1:8080/api/vblog/v1/tokens' \
--header 'Content-Type: application/json' \
--data '{
    "username": "admin",
    "password": "123456"
}'
```

## 其他的优化

### Make

```makefile
PKG := "gitlab.com/go-course-project/go17/vblog"
MOD_DIR := $(shell go env GOPATH)/pkg/mod
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/ | grep -v redis)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)

.PHONY: all dep lint vet test test-coverage build clean

all: build

dep: ## Get the dependencies
    @go mod tidy

lint: ## Lint Golang files
    @golint -set_exit_status ${PKG_LIST}

vet: ## Run go vet
    @go vet ${PKG_LIST}

run: ## Run Server
    @go run main.go start

help: ## Display this help screen
    @grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
```

### Cookie机制

需要手动携带Token, Post, 这会比较安全, 但是有些时候不方便, 如果你想要每次服务端调用 浏览器或者Postman这类工具 可以自动给你携带

颁发Token的时候, 通过客户端(浏览器/Postman), Set-Cookie头

1. 颁发Token是不是要 Set-Cookie
```go
// 设置Cookie
ctx.SetCookie(token.COOKIE_NAME, ins.AccessToken, ins.AccessTokenExpireTTL(), "/", application.Get().Domain(), false, true)
```

2. 获取Token是不是也需要Cookie当作获取Token

```go
// 再尝试从cookie中获取
if accessToken == "" {
    tc, err := c.Cookie(token.COOKIE_NAME)
    if err != nil {
        log.L().Error().Msgf("get cookie error, %s", err)
    } else {
        accessToken = tc
    }
}
```

### 优雅关闭

```go
// 优雅关闭HTTP服务
if err := h.server.Shutdown(ctx); err != nil {
    return fmt.Errorf("http graceful shutdown timeout, force exit")
}
```

### DI: 依赖注入

通过配置 ---> 获取一个对象，直接使用

+ 注入工具：https://www.mcube.top/docs/framework/component/mysql/
+ 为你程序注入配置:

```toml
[token]
  default_expired_ttl = 3600
```



```go
// 定义一个struct, 用于实现 UserService就是刚才定义的接口
// 怎么才能判断这个结构体没有实现这个接口
// unmashal TokenServiceImpl <---> Config文件
type TokenServiceImpl struct {
    ioc.ObjectImpl

    // user service
    user user.AdminService

    // 如果你控制器有一些配置
    DefaultExpiredTTL int `json:"default_expired_ttl" yaml:"default_expired_ttl" toml:"default_expired_ttl" env:"DEFAULT_EXPIRED_TTL"`
}

func (*TokenServiceImpl) Name() string {
    return token.AppName
}
```

## 项目的部署

# API 规范

* https://go-zero.dev/docs/tutorials

## api 语法标记

### syntax 语句

syntax 语句用于标记 api 语言的版本，不同的版本可能语法结构有所不同，随着版本的提升会做不断的优化，当前版本为 v1。

* 语法
```text
SyntaxStmt = "syntax" "=" "v1" .
```

- 写法
```text
syntax = "v1"
```




### info 语句

info 语句是 api 语言的 meta 信息，其仅用于对当前 api 文件进行描述，暂不参与代码生成，其和注释还是有一些区别，注释一般是依附某个 syntax 语句存在，而 info 语句是用于描述整个 api 信息的，当然，不排除在将来会参与到代码生成里面来，info 语句的 EBNF 表示为

- 语法
```text
InfoStmt         = "info" "(" { InfoKeyValueExpr } ")" .
InfoKeyValueExpr = InfoKeyLit [ interpreted_string_lit ] .
InfoKeyLit       = identifier ":" .
```

- 写法
```text
// 不包含 key-value 的 info 块
info ()

// 包含 key-value 的 info 块
info (
    foo: "bar"
    bar:
)
```

### import 语句

- 语法
```text
ImportStmt        = ImportLiteralStmt | ImportGroupStmt .
ImportLiteralStmt = "import" interpreted_string_lit .
ImportGroupStmt   = "import" "(" { interpreted_string_lit } ")" .
```

- 写法
```text
// 单行 import
import "foo"
import "/path/to/file"

// import 组
import ()
import (
    "bar"
    "relative/to/file"
)
```

### 数据类型

- 语法
```text
TypeStmt          = TypeLiteralStmt | TypeGroupStmt .
TypeLiteralStmt   = "type" TypeExpr .
TypeGroupStmt     = "type" "(" { TypeExpr } ")" .
TypeExpr          = identifier [ "=" ] DataType .
DataType          = AnyDataType | ArrayDataType | BaseDataType |
                    InterfaceDataType | MapDataType | PointerDataType |
                    SliceDataType | StructDataType .
AnyDataType       = "any" .
ArrayDataType     = "[" { decimal_digit } "]" DataType .
BaseDataType      = "bool"    | "uint8"     | "uint16"     | "uint32" | "uint64"  |
                    "int8"    | "int16"     | "int32"      | "int64"  | "float32" |
                    "float64" | "complex64" | "complex128" | "string" | "int"     |
                    "uint"    | "uintptr"   | "byte"       | "rune"   | "any"     | .

InterfaceDataType = "interface{}" .
MapDataType       = "map" "[" DataType "]" DataType .
PointerDataType   = "*" DataType .
SliceDataType     = "[" "]" DataType .
StructDataType    = "{" { ElemExpr } "}" .
ElemExpr          = [ ElemNameExpr ]  DataType [ Tag ].
ElemNameExpr      = identifier { "," identifier } .
Tag               = raw_string_lit .
```

- 写法
```text
// 别名类型 [1]
type Int int
type Integer = int

// 空结构体
type Foo {}

// 单个结构体
type Bar {
    Foo int               `json:"foo"`
    Bar bool              `json:"bar"`
    Baz []string          `json:"baz"`
    Qux map[string]string `json:"qux"`
}

type Baz {
    Bar    `json:"baz"`
    // 结构体内嵌 [2]
    Qux {
        Foo string `json:"foo"`
        Bar bool   `json:"bar"`
    } `json:"baz"`
}

// 空结构体组
type ()

// 结构体组
type (
    Int int
    Integer = int
    Bar {
        Foo int               `json:"foo"`
        Bar bool              `json:"bar"`
        Baz []string          `json:"baz"`
        Qux map[string]string `json:"qux"`
    }
)

```

**注意**

[1] 虽然语法上支持别名，但是在语义分析时会对别名进行拦截，这或在将来进行放开。

[2] 虽然语法上支持结构体内嵌，但是在语义分析时会对别名进行拦截，这或在将来进行放开。

除此之外：

目前 api 语法中虽然支持了数组的语法，但是在语义分析时会对数组进行拦截，目前建议用切片替代，这或在将来放开。
不支持 package 设计，如 time.Time。



### service 语句

#### server 语句 

```api
// 空内容
@server()

// 有内容
@server (
    // jwt 声明
    // 如果 key 固定为 “jwt:”，则代表开启 jwt 鉴权声明
    // value 则为配置文件的结构体名称
    jwt: Auth

    // 路由前缀
    // 如果 key 固定为 “prefix:”
    // 则代表路由前缀声明，value 则为具体的路由前缀值，字符串中没让必须以 / 开头
    prefix: /v1

    // 路由分组
    // 如果 key 固定为 “group:”，则代表路由分组声明
    // value 则为具体分组名称，在 goctl生成代码后会根据此值进行文件夹分组
    group: Foo

    // 中间件
    // 如果 key 固定为 middleware:”，则代表中间件声明
    // value 则为具体中间件函数名称，在 goctl生成代码后会根据此值进生成对应的中间件函数
    middleware: AuthInterceptor

    // 超时控制
    // 如果 key 固定为  timeout:”，则代表超时配置
    // value 则为具体中duration，在 goctl生成代码后会根据此值进生成对应的超时配置
    timeout: 3s

    // 其他 key-value，除上述几个内置 key 外，其他 key-value
    // 也可以在作为 annotation 信息传递给 goctl 及其插件，但就
    // 目前来看，goctl 并未使用。
    foo: bar
)
```

#### 服务条目

##### doc 语句

@doc 语句是对单个路由的 meta 信息描述，一般为 key-value 值，可以传递给 goctl 及其插件来进行扩展生成，其 EBNF 表示为：

* 语法

```go

AtDocStmt        = AtDocLiteralStmt | AtDocGroupStmt .
AtDocLiteralStmt = "@doc" interpreted_string_lit .
AtDocGroupStmt   = "@doc" "(" { AtDocKVExpr } ")" .
AtDocKVExpr      = AtServerKeyLit  interpreted_string_lit .
AtServerKeyLit   = identifier ":" .
```
* 写法

```text
// 单行 @doc
@doc "foo"

// 空 @doc 组
@doc ()

// 有内容的 @doc 组
@doc (
foo: "bar"
bar: "baz"
)
```

##### handler 语句

@handler 语句是对单个路由的 handler 信息控制，主要用于生成 golang http.HandleFunc 的实现转换方法，其 EBNF 表示为：

```text
AtHandlerStmt = "@handler" identifier .
```

```api
@handler foo
```

##### 路由语句

* 语法

```text
RouteStmt = Method PathExpr [ BodyStmt ] [ "returns" ] [ BodyStmt ].
Method    = "get"     | "head"    | "post" | "put" | "patch" | "delete" |
            "connect" | "options" | "trace" .
PathExpr  = "/" identifier { ( "-" identifier ) | ( ":" identifier) } .
BodyStmt  = "(" identifier ")" .
```

* 写法

```text
// 没有请求体和响应体的写法
get /ping

// 只有请求体的写法
get /foo (foo)                      // 有 "传参", 没有 "返回"

// 只有响应体的写法
post /foo returns (foo)             // 没有 "传参", 有 "返回"

// 有请求体和响应体的写法
post /foo (foo) returns (bar)       // 有 "传参", 有 "返回"
```

##### api 定义完整示例

```api
syntax = "v1"

info (
    title:   "api 文件完整示例写法"
    desc:    "演示如何编写 api 文件"
    author:  "keson.an"
    date:    "2022 年 12 月 26 日"
    version: "v1"
)

type UpdateReq {
    Arg1 string `json:"arg1"`
}

type ListItem {
    Value1 string `json:"value1"`
}

type LoginReq {
    Username string `json:"username"`
    Password string `json:"password"`
}

type LoginResp {
    Name string `json:"name"`
}

type FormExampleReq {
    Name string `form:"name"`
}

type PathExampleReq {
    // path 标签修饰的 id 必须与请求路由中的片段对应，如
    // id 在 service 语法块的请求路径上一定会有 :id 对应，见下文。
    ID string `path:"id"`
}

type PathExampleResp {
    Name string `json:"name"`
}

@server (
    jwt:        Auth            // 对当前 Foo 语法块下的所有路由，开启 jwt 认证，不需要则请删除此行
    prefix:     /v1             // 对当前 Foo 语法块下的所有路由，新增 /v1 路由前缀，不需要则请删除此行
    group:      g1              // 对当前 Foo 语法块下的所有路由，路由归并到 g1 目录下，不需要则请删除此行
    timeout:    3s              // 对当前 Foo 语法块下的所有路由进行超时配置，不需要则请删除此行
    middleware: AuthInterceptor // 对当前 Foo 语法块下的所有路由添加中间件，不需要则请删除此行
    maxBytes:   1048576         // 对当前 Foo 语法块下的所有路由添加请求体大小控制，单位为 byte,goctl 版本 >= 1.5.0 才支持
)
service Foo {
    // 定义没有请求体和响应体的接口，如 ping
    @handler ping
    get /ping

    // 定义只有请求体的接口，如更新信息
    @handler update
    post /update (UpdateReq)

    // 定义只有响应体的结构，如获取全部信息列表
    @handler list
    get /list returns ([]ListItem)

    // 定义有结构体和响应体的接口，如登录
    @handler login
    post /login (LoginReq) returns (LoginResp)

    // 定义表单请求
    @handler formExample
    post /form/example (FormExampleReq)

    // 定义 path 参数
    @handler pathExample
    get /path/example/:id (PathExampleReq) returns (PathExampleResp)
}


```


---



# 类型声明

* https://go-zero.dev/docs/tutorials/api/types

> 暂不支持 泛型, 弱类型, 比如 any 等


在 API 描述语言中，类型声明需要满足如下规则:

* 类型声明必须以 type 开头
* 不需要声明 struct关键字
* 不支持嵌套结构体声明
* 不支持别名

## 示例1

```api
type Request {
	Name string `path:"name,options=you|me"`
}

type Response {
	Message string `json:"message"`
}

service greet-api {
	@handler GreetHandler
	get /from/:name(Request) returns (Response)
}
```

## 示例2

```api

type StructureExample {
    // 基本数据类型示例
    BaseInt     int     `json:"base_int"`
    BaseBool    bool    `json:"base_bool"`
    BaseString  string  `json:"base_string"`
    BaseByte    byte    `json:"base_byte"`
    BaseFloat32 float32 `json:"base_float32"`
    BaseFloat64 float64 `json:"base_float64"`
    // 切片示例
    BaseIntSlice     []int     `json:"base_int_slice"`
    BaseBoolSlice    []bool    `json:"base_bool_slice"`
    BaseStringSlice  []string  `json:"base_string_slice"`
    BaseByteSlice    []byte    `json:"base_byte_slice"`
    BaseFloat32Slice []float32 `json:"base_float32_slice"`
    BaseFloat64Slice []float64 `json:"base_float64_slice"`
    // map 示例
    BaseMapIntString      map[int]string               `json:"base_map_int_string"`
    BaseMapStringInt      map[string]int               `json:"base_map_string_int"`
    BaseMapStringStruct   map[string]*StructureExample `json:"base_map_string_struct"`
    BaseMapStringIntArray map[string][]int             `json:"base_map_string_int_array"`
    // 匿名示例
    *Base
    // 指针示例
    Base4 *Base `json:"base4"`
    
    // 新的特性（ goctl >= 1.5.1 版本支持 ）
    // 标签忽略示例
    TagOmit string
}
```




# 路由前缀

> 我们通过在 @server 中来通过 prefix 关键字声明了路由前缀，
> 然后通过 @handler 来声明了路由处理函数，
> 这样我们就可以通过路由前缀来区分不同的版本了。

```api
// 有前缀的 api
syntax = "v1"

type UserV1 {
    Name string `json:"name"`
}

type UserV2 {
    Name string `json:"name"`
}

@server (
    prefix: /v1 // 前缀
)
service user-api {
    @handler usersv1    // 路由处理函数
    get /users returns ([]UserV1)
}

// 对服务的一些说明
// 
// 前缀
// 分组
@server (
    prefix: /v2 // 前缀
    // ... 这里可以加上分组的说明
//    group:  user    // 这就是分组的说明
)
service user-api {
    @handler usersv2
    get /users returns ([]UserV2)
}



// 没有前缀的 api
type Request {
    Name string `path:"name,options=you|me"`
}

type Response {
    Message string `json:"message"`
}

// 这里没有前缀
service greet-api {
    @handler GreetHandler
    get /from/:name(Request) returns (Response)
}
```


路由生成代码

> 我们可以看到，我们声明的 prefix 其实在生成代码后通过 rest.WithPrefix 来声明了路由前缀，
> 这样我们就可以通过路由前缀来区分不同的版本了。

```go
// 带有前缀的 路由生成代码
func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
    server.AddRoutes(
        []rest.Route{
            {
                Method:  http.MethodGet,
                Path:    "/users",
                Handler: usersv1Handler(serverCtx),
            },
        },
        rest.WithPrefix("/v1"), // 声明前缀
    )

    server.AddRoutes(
        []rest.Route{
            {
                Method:  http.MethodGet,
                Path:    "/users",
                Handler: usersv2Handler(serverCtx),
            },
        },
        rest.WithPrefix("/v2"),
    )
}


// 普通的路由生成代码
func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
    server.AddRoutes(
        []rest.Route{
            {
                Method:  http.MethodGet,
                Path:    "/from/:name",
                Handler: GreetHandler(serverCtx),
            },
        },
    )
}
```

# 服务分组

```http request

https://example.com/v1/user/login
https://example.com/v1/user/info
https://example.com/v1/user/info/update
https://example.com/v1/user/list

https://example.com/v1/user/role/list
https://example.com/v1/user/role/update
https://example.com/v1/user/role/info
https://example.com/v1/user/role/add
https://example.com/v1/user/role/delete

https://example.com/v1/user/class/list
https://example.com/v1/user/class/update
https://example.com/v1/user/class/info
https://example.com/v1/user/class/add
https://example.com/v1/user/class/delete
```


## 不分组效果

### api

```api
syntax = "v1"

type (
    UserLoginReq{}
    UserInfoReq{}
    UserLoginResp{}
    UserInfoResp{}
    UserInfoUpdateReq{}
    UserInfoUpdateResp{}
)

type (
    UserRoleReq{}
    UserRoleResp{}
    UserRoleUpdateReq{}
    UserRoleUpdateResp{}
    UserRoleAddReq{}
    UserRoleAddResp{}
    UserRoleDeleteReq{}
    UserRoleDeleteResp{}
)

type (
    UserClassReq{}
    UserClassResp{}
    UserClassUpdateReq{}
    UserClassUpdateResp{}
    UserClassAddReq{}
    UserClassAddResp{}
    UserClassDeleteReq{}
    UserClassDeleteResp{}
)
@server(
    prefix: /v1
)
service user-api {
    @handler UserLogin
    post /user/login (UserLoginReq) returns (UserLoginResp)

    @handler UserInfo
    post /user/info (UserInfoReq) returns (UserInfoResp)

    @handler UserInfoUpdate
    post /user/info/update (UserInfoUpdateReq) returns (UserInfoUpdateResp)

    @handler UserList
    get /user/list returns ([]UserInfoResp)

    @handler UserRoleList
    get /user/role/list returns ([]UserRoleResp)

    @handler UserRoleUpdate
    get /user/role/update (UserRoleUpdateReq) returns (UserRoleUpdateResp)

    @handler UserRoleInfo
    get /user/role/info (UserRoleReq) returns (UserRoleResp)

    @handler UserRoleAdd
    get /user/role/add (UserRoleAddReq) returns (UserRoleAddResp)

    @handler UserRoleDelete
    get /user/role/delete (UserRoleDeleteReq) returns (UserRoleDeleteResp)

    @handler UserClassList
    get /user/class/list returns ([]UserClassResp)

    @handler UserClassUpdate
    get /user/class/update (UserClassUpdateReq) returns (UserClassUpdateResp)

    @handler UserClassInfo
    get /user/class/info (UserClassReq) returns (UserClassResp)

    @handler UserClassAdd
    get /user/class/add (UserClassAddReq) returns (UserClassAddResp)

    @handler UserClassDelete
    get /user/class/delete (UserClassDeleteReq) returns (UserClassDeleteResp)
}
```

### 生成的代码

```text
.
├── etc
│ └── user-api.yaml
├── internal
│ ├── config
│ │ └── config.go
│ ├── handler
│ │ ├── routes.go
│ │ ├── userclassaddhandler.go
│ │ ├── userclassdeletehandler.go
│ │ ├── userclassinfohandler.go
│ │ ├── userclasslisthandler.go
│ │ ├── userclassupdatehandler.go
│ │ ├── userinfohandler.go
│ │ ├── userinfoupdatehandler.go
│ │ ├── userlisthandler.go
│ │ ├── userloginhandler.go
│ │ ├── userroleaddhandler.go
│ │ ├── userroledeletehandler.go
│ │ ├── userroleinfohandler.go
│ │ ├── userrolelisthandler.go
│ │ └── userroleupdatehandler.go
│ ├── logic
│ │ ├── userclassaddlogic.go
│ │ ├── userclassdeletelogic.go
│ │ ├── userclassinfologic.go
│ │ ├── serclasslistlogic.go
│ │ ├── userclassupdatelogic.go
│ │ ├── userinfologic.go
│ │ ├── userinfoupdatelogic.go
│ │ ├── userlistlogic.go
│ │ ├── userloginlogic.go
│ │ ├── userroleaddlogic.go
│ │ ├── userroledeletelogic.go
│ │ ├── userroleinfologic.go
│ │ ├── userrolelistlogic.go
│ │ └── userroleupdatelogic.go
│ ├── svc
│ │ └── servicecontext.go
│ └── types
│     └── types.go
├── user.api
└── user.go

7 directories, 35 files

```


## 分组效果

由于我们没有进行分组，所以生成的代码中 **handler** 和 **logic** 目录下的文件是全部揉在一起的，这样的目录结构在项目中不太好管理和阅读，接下来我们按照 user，role，class 来进行分组，在 api 语言中，我们可以通过在 @server 语句块中使用 group 关键字来进行分组，分组的语法如下：

### api
```api
syntax = "v1"

type (
    UserLoginReq  {}
    UserInfoReq  {}
    UserLoginResp  {}
    UserInfoResp  {}
    UserInfoUpdateReq  {}
    UserInfoUpdateResp  {}
)

type (
    UserRoleReq  {}
    UserRoleResp  {}
    UserRoleUpdateReq  {}
    UserRoleUpdateResp  {}
    UserRoleAddReq  {}
    UserRoleAddResp  {}
    UserRoleDeleteReq  {}
    UserRoleDeleteResp  {}
)

type (
    UserClassReq  {}
    UserClassResp  {}
    UserClassUpdateReq  {}
    UserClassUpdateResp  {}
    UserClassAddReq  {}
    UserClassAddResp  {}
    UserClassDeleteReq  {}
    UserClassDeleteResp  {}
)

@server (
    prefix: /v1
    group:  user
)
service user-api {
    @handler UserLogin
    post /user/login (UserLoginReq) returns (UserLoginResp)

    @handler UserInfo
    post /user/info (UserInfoReq) returns (UserInfoResp)

    @handler UserInfoUpdate
    post /user/info/update (UserInfoUpdateReq) returns (UserInfoUpdateResp)

    @handler UserList
    get /user/list returns ([]UserInfoResp)
}

// 定义服务
@server (
    // 声明前缀
    prefix: /v1
    
    // 声明分组
    group:  role
)
service user-api {
    @handler UserRoleList
    get /user/role/list returns ([]UserRoleResp)

    @handler UserRoleUpdate
    get /user/role/update (UserRoleUpdateReq) returns (UserRoleUpdateResp)

    @handler UserRoleInfo
    get /user/role/info (UserRoleReq) returns (UserRoleResp)

    @handler UserRoleAdd
    get /user/role/add (UserRoleAddReq) returns (UserRoleAddResp)

    @handler UserRoleDelete
    get /user/role/delete (UserRoleDeleteReq) returns (UserRoleDeleteResp)
}

@server (
    prefix: /v1
    group:  class
)
service user-api {
    @handler UserClassList
    get /user/class/list returns ([]UserClassResp)

    @handler UserClassUpdate
    get /user/class/update (UserClassUpdateReq) returns (UserClassUpdateResp)

    @handler UserClassInfo
    get /user/class/info (UserClassReq) returns (UserClassResp)

    @handler UserClassAdd
    get /user/class/add (UserClassAddReq) returns (UserClassAddResp)

    @handler UserClassDelete
    get /user/class/delete (UserClassDeleteReq) returns (UserClassDeleteResp)
}


```

### 生成的代码

> **handler** **logic** 里面的代码 根绝 api 分组(**分别存在放对应的层级目录里面**)

```text
.
├── etc
│ └── user-api.yaml
├── internal
│ ├── config
│ │ └── config.go
│ ├── handler
│ │ ├── class
│ │ │ ├── userclassaddhandler.go
│ │ │ ├── userclassdeletehandler.go
│ │ │ ├── userclassinfohandler.go
│ │ │ ├── userclasslisthandler.go
│ │ │ └── userclassupdatehandler.go
│ │ ├── role
│ │ │ ├── userroleaddhandler.go
│ │ │ ├── userroledeletehandler.go
│ │ │ ├── userroleinfohandler.go
│ │ │ ├── userrolelisthandler.go
│ │ │ └── userroleupdatehandler.go
│ │ ├── routes.go
│ │ └── user
│ │     ├── userinfohandler.go
│ │     ├── userinfoupdatehandler.go
│ │     ├── userlisthandler.go
│ │     └── userloginhandler.go
│ ├── logic
│ │ ├── class
│ │ │ ├── userclassaddlogic.go
│ │ │ ├── userclassdeletelogic.go
│ │ │ ├── userclassinfologic.go
│ │ │ ├── userclassupdatelogic.go
│ │ │ └── usersclaslistlogic.go
│ │ ├── role
│ │ │ ├── userroleaddlogic.go
│ │ │ ├── userroledeletelogic.go
│ │ │ ├── userroleinfologic.go
│ │ │ ├── userrolelistlogic.go
│ │ │ └── userroleupdatelogic.go
│ │ └── user
│ │     ├── userinfologic.go
│ │     ├── userinfoupdatelogic.go
│ │     ├── userlistlogic.go
│ │     └── userloginlogic.go
│ ├── svc
│ │ └── servicecontext.go
│ └── types
│     └── types.go
├── user.api
└── user.go

13 directories, 35 files
```

通过分组，我们可以很方便的将不同的业务逻辑分组到不同的目录下，这样可以很方便的管理不同的业务逻辑。


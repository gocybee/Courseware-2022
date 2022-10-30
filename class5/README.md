# gocybee——第五节课

## 前言

可能很多同学在上完前几节课之后还处于一个非常懵逼的状态，不过恭喜大伙，坚持到现在后面的内容基本不会劝退！

## gin

### gin 是什么

我们直接借用官网的话。

> Gin is a web framework written in Go (Golang). It features a martini-like API with performance that is up to 40 times faster thanks to httprouter. If you need performance and good productivity, you will love Gin.

翻译成人话就是， gin 是一个很性能很强且很容易上手的框架，我们不多介绍这些有的没的，直接上手使用！

## 实战

因为这是我们真正意义上的第一个 Web 项目（大佬当我没说），所以我们还是简单给大家介绍一下我们新建项目的必备步骤。

### go mod

现阶段大家只需要掌握两个命令就行了，分别是 `go mod init` 和 `go mod tidy` 。

#### go mod init

这个命令是我们开始一个项目的必须命令，在现在的 Go 开发中，使用 go mod 来进行依赖管理是必不可少的。

首先我们打开终端（直接使用 Goland 那个就行），在项目的根目录中使用 `go mod init 你的项目名` 。

eg:
```go
go mod init gin-demo
```

接下来在 Goland 的设置中打开启用 Go 依赖管理（若是英文也会在同样的位置）。

![](https://picture.lanlance.cn/i/2022/10/30/635e6fc0d9ad0.png)

到现在我们的项目的前置工作就差不多完成了，在其他的项目中都需要重复上述操作

#### go mod tidy

一句话，拉取缺少的依赖，移除不用的依赖

### 单体架构

在以往的学习或者作业中，大多时候用一个 `main.go` 文件就能完成我们的需要，但是在大型的开发项目中，如果仍然只使用一个文件来放入我们所有的代码，我想那样的代码没人会想看。

所以说学会分包是很有必要的，在这里向大家介绍一个最简单也很有逻辑的单体架构方案(在后续的学习中可以接着模仿，但希望大家能思考并设计出自己的、能让自己开发更有效率的架构)

```
├── README.md
├── apis
├── dao
├── go.mod
├── model
└── utils
```

在这里简单解释一下每一层的含义

- README.md：项目的说明文档，大家可以提前学习如何写出一个优秀的说明文档（当然在这个项目中的 README 是你们的课件）。
- apis：接口层，在里面是详细的逻辑实现以及路由。
- dao：全名为 data access object，说人话就是操作数据库的。
- model：模型层，主要放数据库实例的结构体。
- utils：一些常用的工具函数，封装在这里减少代码的重复使用。
- go.mod：依赖管理

### 使用 gin 进行简单的项目开发

#### 第一个简单的 Web 服务

首先我们需要将 gin 的源码从 GitHub 中拉取下来，在终端中输入以下代码

```shell
go get github.com/gin-gonic/gin
```

接着我们在根目录创建一个 `main.go` 文件，先把 gin 的样例代码跑一遍，看有没有问题

```go
package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })
  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
```

在 Postman 或者其他软件进行一下测试，我这里就直接在终端测试了

```sh
curl 127.0.0.1:8080/ping
```

**输出**

```sh
{"message":"pong"}
```

如果输出和我一样的话说明没有问题，咱们的第一个 Web 服务就运行成功了。

#### 解析第一个 Web 服务

很多同学估计还是现在还是很懵逼，对刚才运行的代码完全看不懂，我的评价是看得懂才是有鬼了，所以我一行一行给大家剖析一下这个样例服务。

首先第一行我们将 `gin.Default()` 赋值给了 r，进入到源码中我们可以知道调用 `gin.Default()` 会返回已附加 Logger 和 Recovery 中间件的 Engine 实例。这里的 r 就是我们的 Engine 实例。

```go
// gin 源码
// Default returns an Engine instance with the Logger and Recovery middleware already attached.
func Default() *Engine {
	debugPrintWARNINGDefault()
	engine := New()
	engine.Use(Logger(), Recovery())
	return engine
}
```

其中的 Logger 和 Recovery 中间件在这里细说的话大家可能也挺不太懂，简单说就是一个是用于记录日志的，另一个是用于捕获程序问题并让程序恢复使用的。

当然我们也可以直接不使用 `Default` ，这些就在下面的实战部分中再细讲。

再到下面的五行，是我们的 Engine 实例调用 GET 实现了一个 GET 请求的接口，"/ping"为我们的路径，后面的一坨是我们的详细逻辑。如果我们把它抽离出来会更好理解

```go
func main() {
	r := gin.Default()
	r.GET("/ping", Ping)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
```

这样是不是清晰了很多，我们通过 GET 请求访问 "/ping" 就会获得一个 JSON 格式的返回值。

最后的 `r.Run()` 代表我们服务监听在 localhost:8080 。

> localhost 和 127.0.0.1 是一个意思，都代表本地。

其中我们监听的地址和端口号都可以自定义，在下面我们会细讲。

#### 第一个复杂一点的 Web 服务

让我们回到刚才的架构中，从现在开始我们一步一步搭建我们的 Web 服务。我们先提前规划好我们想得到一个什么样的 Web 服务，我认为新手上手最经典的就是实现一个登录注册的功能。在实现我们的服务之前我们先来预备一点前置知识，也就是 cookie、session、token 和 JWT。

**！！！**

**（以下内容又臭又长，不想看可以直接跳到实战，不过还是建议大家都了解一下）**

**！！！**

##### Cookie、Session、JWT

###### 一、先了解几个基础概念

**什么是认证（Authentication）**

通俗地讲就是验证当前用户的身份。

互联网中的认证：

用户名密码登录、邮箱发送登录链接、手机号接收验证码。
只要你能收到邮箱/验证码，就默认你是账号的主人。

**什么是授权（Authorization）**

用户授予第三方应用访问该用户某些资源的权限。

实现授权的方式有：cookie、session、token、OAuth。

**什么是凭证（Credentials）**

实现认证和授权的前提是需要一种媒介（证书）来标记访问者的身份。

在互联网应用中，一般网站（如掘金）会有两种模式，游客模式和登录模式。游客模式下，可以正常浏览网站上面的文章，一旦想要点赞/收藏/分享文章，就需要登录或者注册账号。当用户登录成功后，服务器会给该用户使用的浏览器颁发一个令牌（token），这个令牌用来表明你的身份，每次浏览器发送请求时会带上这个令牌，就可以使用游客模式下无法使用的功能。

###### 二、Cookie
Cookie 最开始被设计出来是为了弥补HTTP在状态管理上的不足。HTTP 协议是一个无状态协议，客户端向服务器发请求，服务器返回响应，故事就这样结束了，但是下次发请求如何让服务端知道客户端是谁呢？这种背景下，就产生了 Cookie。

- cookie 存储在客户端： cookie 是服务器发送到用户浏览器并保存在本地的一小块数据，它会在浏览器下次向同一服务器再发起请求时被携带并发送到服务器上。因此，服务端脚本就可以读、写存储在客户端的cookie的值。
- cookie 是不可跨域的： 每个 cookie 都会绑定单一的域名（绑定域名下的子域都是有效的），无法在别的域名下获取使用，同域名不同端口也是允许共享使用的。

服务端只需要设置 setCookie 这个 header，之后浏览器会自动把 cookie 写入到我们的浏览器存起来，然后当前域名在发送请求的时候都会自动带上这个 cookie。

###### 三、Session
session 是另一种记录服务器和客户端会话状态的机制。session 是基于 cookie 实现的，session 存储在服务器端，sessionId 会被存储到客户端的cookie 中。


session 认证流程：

1. 用户第一次请求服务器的时候，服务器根据用户提交的相关信息，创建对应的 Session
2. 请求返回时将此 Session 的唯一标识 SessionID 返回给浏览器
3. 浏览器接收到服务器返回的 SessionID 后，会将此信息存入到 Cookie 中，同时 Cookie 记录此 SessionID 属于哪个域名
4. 当用户第二次访问服务器的时候，请求会自动把此域名下的 Cookie 信息也发送给服务端，服务端会从 Cookie 中获取 SessionID，
5. 根据 SessionID 查找对应的 Session 信息，如果没有找到说明用户没有登录或者登录失效，如果找到 Session 证明用户已经登录可执行后面操作。

根据以上流程可知，SessionID 是连接 Cookie 和 Session 的一道桥梁，大部分系统也是根据此原理来验证用户登录状态。

###### 四、Cookie 和 Session 的区别

- 安全性： Session 比 Cookie 安全，Session 是存储在服务器端的，Cookie 是存储在客户端的。
- 存取值的类型不同：Cookie 只支持存字符串数据，Session 可以存任意数据类型。
- 有效期不同： Cookie 可设置为长时间保持，比如我们经常使用的默认登录功能，Session 一般失效时间较短，客户端关闭（默认情况下）或者 Session 超时都会失效。
- 存储大小不同： 单个 Cookie 保存的数据不能超过 4K，Session 可存储数据远高于 Cookie，但是当访问量过多，会占用过多的服务器资源。

###### 五、什么是 Token
Token 是访问接口（API）时所需要的资源凭证。

简单 token 的组成：

uid(用户唯一的身份标识)、time(当前时间的时间戳)、sign（签名，token 的前几位以哈希算法压缩成的一定长度的十六进制字符串）

特点：

- 服务端无状态化、可扩展性好
- 支持移动端设备
- 安全
- token 完全由应用管理，所以它可以避开同源策略

**Access Token**

Access Token 的身份验证流程：

1. 客户端使用用户名跟密码请求登录
2. 服务端收到请求，去验证用户名与密码
3. 验证成功后，服务端会签发一个 token 并把这个 token 发送给客户端
4. 客户端收到 token 以后，会把它存储起来，比如放在 localStorage 里
5. 客户端每次发起请求的时候需要把 token 放到请求的 Header 里传给服务端
6. 服务端收到请求，然后去验证客户端请求里面带着的 token ，如果验证成功，就向客户端返回请求的数据

**Refresh Token**

另外一种 token——refresh token

refresh token 是专用于刷新 access token 的 token。如果没有 refresh token，也可以刷新 access token，但每次刷新都要用户输入登录用户名与密码，会很麻烦。有了 refresh token，可以减少这个麻烦，客户端直接用 refresh token 去更新 access token，无需用户进行额外的操作。

Access Token 的有效期比较短，当 Acesss Token 由于过期而失效时，使用 Refresh Token 就可以获取到新的 Token，如果 Refresh Token 也失效了，用户就只能重新登录了。

Refresh Token 及过期时间是存储在服务器的数据库中，只有在申请新的 Acesss Token 时才会验证，不会对业务接口响应时间造成影响，也不需要向 Session 一样一直保持在内存中以应对大量的请求。

###### 六、Token 和 Session 的区别
Session 是一种记录服务器和客户端会话状态的机制，使服务端有状态化，可以记录会话信息。而 Token 是令牌，访问资源接口（API）时所需要的资源凭证。Token 使服务端无状态化，不会存储会话信息。

Session 和 Token 并不矛盾，作为身份认证 Token 安全性比 Session 好，因为每一个请求都有签名还能防止监听以及重复攻击，而 Session 就必须依赖链路层来保障通讯安全了。如果你需要实现有状态的会话，仍然可以增加 Session 来在服务器端保存一些状态。

如果你的用户数据可能需要和第三方共享，或者允许第三方调用 API 接口，用 Token 。如果永远只是自己的网站，自己的 App，用什么就无所谓了。

###### 七、什么是 JWT

JSON Web Token（简称 JWT）是目前最流行的跨域认证解决方案。是一种认证授权机制。

JWT 是为了在网络应用环境间传递声明而执行的一种基于 JSON 的开放标准。JWT 的声明一般被用来在身份提供者和服务提供者间传递被认证的用户身份信息，以便于从资源服务器获取资源。比如用在用户登录上。
可以使用 HMAC 算法或者是 RSA 的公/私秘钥对 JWT 进行签名。因为数字签名的存在，这些传递的信息是可信的。

1. JWT 的原理

JWT 的原理是，服务器认证以后，生成一个 JSON 对象，返回给用户，就像下面这样。

```json
{
  "姓名": "张三",
  "角色": "管理员",
  "到期时间": "2018年7月1日0点0分"
}
```

以后，用户与服务端通信的时候，都要发回这个 JSON 对象。服务器完全只靠这个对象认定用户身份。为了防止用户篡改数据，服务器在生成这个对象的时候，会加上签名。

2. JWT 认证流程：

1. 用户输入用户名/密码登录，服务端认证成功后，会返回给客户端一个 JWT

2. 客户端将 token 保存到本地（通常使用 localstorage，也可以使用 cookie）

3. 当用户希望访问一个受保护的路由或者资源的时候，需要请求头的 Authorization 字段中使用Bearer 模式添加 JWT，其内容看起来是下面这样

   ```
   Authorization: Bearer <token>
   ```

   - 服务端的保护路由将会检查请求头 Authorization 中的 JWT 信息，如果合法，则允许用户的行为
   - 因为 JWT 是自包含的（内部包含了一些会话信息），因此减少了需要查询数据库的需要
   - 因为 JWT 并不使用 Cookie 的，所以你可以使用任何域名提供你的 API 服务而不需要担心跨域问题
   - 因为用户的状态不再存储在服务端的内存中，所以这是一种无状态的认证机制生成

###### 八、Token 和 JWT 的区别
相同：

- 都是访问资源的令牌
- 都可以记录用户的信息
- 都是使服务端无状态化
- 都是只有验证成功后，客户端才能访问服务端上受保护的资源

区别：

Token：服务端验证客户端发送过来的 Token 时，还需要查询数据库获取用户信息，然后验证 Token 是否有效。

JWT： 将 Token 和 Payload 加密后存储于客户端，服务端只需要使用密钥解密进行校验（校验也是 JWT 自己实现的）即可，不需要查询或者减少查询数据库，因为 JWT 自包含了用户信息和加密的数据。

##### 编写 model

我们在开始我们的项目之前需要想想我们的登录注册需要些什么字段，需要创建一个什么样的结构体，由于这是我们刚开始学习 Web 开发嘛，本着一切从简的原则，我们还是越简单越好。

在 model 文件夹下我们创建一个 `user.go` 文件，内容如下

```go
package model

type User struct {
	Username string
	Password string
}
```

##### 编写 dao

因为我们还没有学习数据库，所以得造一点假的数据，不过数据都是小问题，问题在于我们在数据操作中需要哪些逻辑。

首先是注册，注册我们需要将数据插入数据库中，所以得有一个增加用户数据的操作，除此之外我们还要防止重复用户的出现。然后是登录，登录我们需要看他的密码和数据库中的密码是否匹配，所以我们这个小项目在 dao 层就需要这三个操作。

总结一下

1. 新增用户数据
2. 查找用户（注册时查找是否存在该用户，若存在则注册失败）
3. 查找用户密码（登录时使用）

所以我们可以直接写出我们的代码

```go
```




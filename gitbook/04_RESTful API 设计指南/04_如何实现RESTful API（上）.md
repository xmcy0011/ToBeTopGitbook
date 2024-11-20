# 如何实现RESTful API（上）

## 前言

在前面一系列的文章中，我们介绍了 RESTful API 出现的背景，以及为什么要引入 RESTful API 风格。那么从本章开始，我们将进入实现篇，在本篇中，我们主要关注如何实现 RESTful API 风格。

什么是 RESTful API 风格？在 [RESTful API 设计指南（1）——开篇词](./01_开篇词.md) 中，我们进行了总结：

> 只要你的 HTTP 接口使用 POST/DELETE/PUT/GET 代表增删改查操作，使用 HTTP 状态码代表结果，使用 URI 代表操作对象，你的 API 就是 RESTful API

所以，其背后是一个对 HTTP 协议进行再学习的过程，本质上就是要用好 HTTP 协议：

* **设计好 URL 和 HTTP 主体（path和body）**：充分利用 URL 的 path 和 query 部分，path 中代表要操作的资源（对象），query放查询条件，body中放请求参数和响应结果
* **正确选择 HTTP 方法（method）**：平常你可能只使用 GET 和 POST 方法，但是 HTTP 协议中还提供了 DELETE、PUT 和 OPTION 等方法。所以，REST 风格中就充分利用了 HTTP 应用层协议的特性，把增删改查分别对应到 HTTP 的 POST、DEELTE、PUT和GET方法上。
* **正确选择 HTTP 状态码（status code）**：正确的使用状态码，当因为客户端参数异常而发生错误，应当返回 400 Bad Request，如果是服务端内部原因发生错误，应当返回 500 Internal Server Error，如果修改成功，通常情况下应当返回 204 No Content。
* **充分利用 HTTP 报头（request header)**：一些通用的参数应当优先考虑放在 HTTP 报头中，比如接口认证，每个接口都需要携带 token，可以使用 Authorization 这个报头。

下面，让我们再复习一下 HTTP 协议的这几个关键部分。

## URL

URL 是 URI 的一个子集，特指 HTTP 协议中的资源定位符，URI格式定义如下：

```bash
                    hierarchical part
        ┌───────────────────┴─────────────────────┐
                    authority               path
        ┌───────────────┴───────────────┐┌───┴────┐
  abc://username:password@example.com:123/path/data?key=value&key2=value2#fragid1
  └┬┘   └───────┬───────┘ └────┬────┘ └┬┘           └─────────┬─────────┘ └──┬──┘
scheme  user information     host     port                  query         fragment
```

根据 RFC 3986 文档：

* scheme：协议标识符，用于表示访问资源的具体协议，如下是几个示例：
  * http：http://www.example.com，表示通过 HTTP 协议访问 www.example.com。
  * https：https://secure.example.com，表示通过 HTTPS 协议（HTTP 的安全版本）访问 secure.example.com。
  * ftp：ftp://ftp.example.com/resource.txt，表示通过 FTP 协议访问 ftp.example.com 上的资源 resource.txt。
  * mailto：mailto:someone@example.com，表示一个电子邮件地址 someone@example.com。
  * file：file:///C:/path/to/file.txt，表示本地文件系统中的文件 C:/path/to/file.txt。
* authority：标识资源的访问权限信息，包括用户信息（可选）、主机和端口号（可选），示例如下：
  * 带用户信息的 URI：http://user:password@www.example.com
  * 带端口号的 URI：http://www.example.com:8080
  * 常见的 URI：http://www.example.com
* path：标识资源在主机上的具体位置和层级结构，类似于文件系统中的路径。示例如下：
  * 简单路径：http://www.example.com/index.html
  * 带目录的路径：http://www.example.com/images/photo.jpg
* query：用于在请求资源时传递附加信息，也叫查询参数。位于 path 之后，使用英文问号（?）分割，通常由一个或多个键值对组成，键和值之间用等号（=）分隔，多个键值对之间用和号（&）分隔。

其中，我们需要重点关注 path 和 query 2个部分，正如开头所说，HTTP是应用层协议，要实现 RESTful API 的关键就是，我们要充分利用其应用层协议的特性，不要重复造轮子。

### path

在 RESTful API 中，我们用 path 部分表示我们要操作的资源，资源名通常是名词形式，且用复数形式代表一个集合。如果是单个集合，则使用 {id} 形式表示。

path 具有层级结构，指向的资源拥有从属关系。比如在 http://www.example.com/images/photo1.jpg 的 URL 中，photo.jpg 是属于 images 这个目录的，同样的，images 这个目录下可能还有 photo1.jpg、photo2.jgp、…… 和 photoN.jgp 等文件。

所以，在 RESTful API 中，我们可以利用"父子"关系的这种特性来区分操作的是集合还是单个对象。比如 POST /users 代表要往一个集合中新增元素（用户），DELETE users/1 代表要删除单个用户。其中 users/{id} 就属于父子关系，即 id 指向的单个用户是从属于整体用户的。

除了父子关系之外，还有一种"并列"关系，父子关系使用 "/" 分割，并列关系使用 "," 隔开。

并列关系的设定，通常出现在 PUT 更新操作中，如 PUT /users/{id}/name,age,sex 代表要更新用户的部分属性，这里只更新用户的昵称、年龄和性别，而这3个属性之间就是并列关系，但都从属于单个用户，所以和 {id} 之间用 "/" 隔开。

### query

查询参数，kv形式，多个参数之间通过 & 进行拼接。

以查询用户详细信息为例，在传统的 HTTP API 中，我们会把用户 id 放在 query 中：
GET /queryUserInfo?userId=1

但是在 RESTful API 中，id 是放在 path 中的。如下图，其中 :id 代表了占位符，具体的值在 Path Variables 中指定：

![image](./images/03-path-example.png)

一方面，这里也是父子关系的体现，即用户 33 属于 users 这个集合。另一方面，相比传统 API，RESTful API 中 path 标志了一个唯一的资源，/users/33 永远只会出现一个且始终指向同一个用户。

## HTTP Method

传统 API 只使用 GET 和 POST 2种方法，而 RESTful API 使用 4 种方法。其中 GET 和 POST 的定义首次出现在 HTTP 1.0 （RFC 1945）中，PUT 和 DELETE 出现在间隔几个月后发布的 HTTP 1.1 （RFC 2616）版本中。

其定义如下：

* **GET**：表示检索信息，允许缓存。可以请求任意资源，如 HTML、图像或者视频、JSON文件、XML文件和CSS文件等等，是一种安全操作，意味着它不应该改变服务器上任何资源的状态
* **POST**：将数据发送到服务器进行处理，不允许缓存。数据格式可以是 xml、json等。POST操作是不安全的，因为它有能力更新服务器状态，并且在执行时对服务器的状态造成潜在的副作用。
* **PUT**：用于替换给定URL标识的资源，如果资源已存在，则进行修改，应响应200（OK）或者204（No Content），否则创建资源，返回 201（Created）。PUT操作是不安全的，可以改变服务器资源的状态。但它是幂等的，调用100次 PUT 操作将航班状态设置为准时和调用一次 PUT 操作的效果是一样的，航班最终的状态都是准时。
* **DELETE**：用户删除给定URL标志的资源，如果执行成功且响应中包含说明数据，应响应 200（OK），否则应响应 204（No Content），和PUT操作一样，DELETE方法是幂等但是不安全的。

从 HTTP 协议规范中，我们可以发现 GET、POST、PUT 和 DELETE 刚好对应查、增、改、删操作，也就是我们说的增删改查，所以我们可以利用 HTTP Method 代表 API 的动作。这也是为什么传统 API 需要在 path 中声明动作的原因。而 RESTful API 不需要，还是那句话：充分利用 HTTP 应用层协议的特性！

在 RESTful API 中，Method 通常搭配 path 使用，如果按照面向对象的思想，我们可以把 path 看成对象名，把 HTTP method 看成方法，则他们可能的组合如下：

* `POST /users`：新增一个用户，面向对象表示为 users.add(info)
* `DELETE /users/{id}`：删除一个具体的用户，面向对象表示为 users.deleteById(id)
* `PUT /users/{id}`：修改某个用户的信息，具体的信息放在请求体中，面向对象表示为 users.updateById(id, info)
* `GET /users?offset=10&limit=100`：查询所有用户列表，查询条件放在 query 中，面向对象表示为  users.getAll(offset, limit)
* `GET /users/{id}`：某个某个具体的用户信息，面向对象表示为 users.getByUserId(id)

## HTTP Status Code

在 RFC 2616 中规定了服务器响应HTTP请求的消息格式，由三个主要部分组成：状态行（Status-Line）、响应头（response-header）以及消息体（message-body，对于某些响应可能为空）：

```bash
Response      = Status-Line
                *(( general-header
                 | response-header
                 | entity-header ) CRLF)
                 CRLF
                 [ message-body ]
```

其中 Status Code 代表状态码，出现在 HTTP 响应的起始行（Status-Line）部分，提供了关于响应的基本信息，并总是包含以下三个元素，以空格分隔：

* HTTP版本：例如 HTTP/1.1
* 状态码（Status Code）：一个三位数字，用于描述请求的结果。例如，200表示成功，404表示未找到资源。
* 状态消息（Status Message）：一个简短的文本描述，对应于状态码。例如，对于状态码200，状态消息是"OK"。

如下一个典型的 HTTP 响应起始行中，200 就是 Status Code，代表服务器已成功处理该请求：

```bash
HTTP/1.1 200 OK
```

状态码是 HTTP 协议的核心组成部分，用于指示服务器对请求的处理结果。它们被分为五类，范围从100到599，每类都代表了不同种类的响应。例如，1xx类表示信息性响应，2xx类表示成功，3xx类表示重定向，4xx类表示客户端错误，而5xx类表示服务器错误。

在一些不规范的API设计中，对于 Status Code 的处理，可能会永远返回 200，具体的请求结果被封装在了 Body 中：

```bash
HTTP/1.1 200 OK 
Content-Type: application/json 
Content-Length: 123 
```

```json
{
    "errorCode": 0,
    "errorMsg": "success",
    "data": {
        "userId": 123,
        "username": "john_doe",
        "email": "john.doe@example.com",
        "isAdmin": false,
        "createdAt": "2023-04-01T12:00:00Z"
    }
}
```

客户端通过解析 "errorCode" 来判断请求成功还是失败，通过 "errorMsg" 来给出错误提示，可能会定义如下 errorCode：

* 0：成功
* 4000001：参数错误
* 4000003：权限不足
* 4000004：用户不存在
* ……

随着功能的增加，系统中错误码可能越来越多：

* 如果要判断某一类错误，比如参数异常导致的错误或者属于某个服务返回的错误，则我们必须对错误码进行分组。比如 4000000 - 4100000 代表服务A的错误码范围或者代表参数错误。
* 引入微服务架构后，各个服务之间错误码可能出现重复。所以需要维护一个全局的仓库？每次新增错误码都需要小心翼翼的查阅，应该取多少值合适。
* 错误码太多了，维护成本越来越高，查找也越来越困难，什么情况下会返回这个错误码？是因为什么原因？
* 跨团队协作时，大家习惯不一致。可能有些人对成功的错误码定义成0，有的又定义成一个固定的数字，甚至有的直接使用字符串 "success" 返回，导致服务对接比较痛苦

等等这些问题都要考虑，但是如果我们回顾 RFC 文档中对 Status-Code 的定义，如上的问题其实我们也就解决了大半：

* 2xx：代表成功，不管是任何服务，都按此规范实现API行为。比如 200 OK 通常并表示成功并且有 Body 响应体，201 Created 表示创建成功，204 Not Conent 表示没有响应数据，如修改成功。
* 4xx：表示由客户端导致的错误。如果是参数类异常，就返回 400 Bad Request；是权限不够，就返回 403 Forbidden。要查询的用户不存在，就返回 404 Not Found。
* 5xx：表示由服务端导致的错误。如调用第三方接口超时，数据库异常，内部崩溃错误或者触发了限流等等。

如下是一些在 RESTful API 中常用的 Status-Code：

```bash
| "200"  ; Section 10.2.1: OK
| "201"  ; Section 10.2.2: Created
| "203"  ; Section 10.2.4: Non-Authoritative Information
| "204"  ; Section 10.2.5: No Content
| "400"  ; Section 10.4.1: Bad Request
| "401"  ; Section 10.4.2: Unauthorized
| "402"  ; Section 10.4.3: Payment Required
| "403"  ; Section 10.4.4: Forbidden
| "404"  ; Section 10.4.5: Not Found
| "408"  ; Section 10.4.9: Request Time-out
| "409"  ; Section 10.4.10: Conflict
| "500"  ; Section 10.5.1: Internal Server Error
| "501"  ; Section 10.5.2: Not Implemented
| "502"  ; Section 10.5.3: Bad Gateway
| "503"  ; Section 10.5.4: Service Unavailable
```

在 RFC 7231 中，对这些状态码进行了更详细的描述，有兴趣可以阅读一下。

## 总结

实现 RESTful API 的关键是用好 HTTP 协议：

* 设计好 URL 和 HTTP 主体
* 正确选择 HTTP 方法
* 正确选择 HTTP 状态码
* 充分利用 HTTP 报头

所以本章对 HTTP 协议中的 path、method 和 status code 进行了简单的回顾，然而 HTTP 协议的内容远远不止于此，比如HTTP报头和HTTP主体这里并没有深入介绍，相关内容需要读者更深入的学习。

下一章，我们将使用 go + gin 实现一个简单的 RESTful API 服务，从代码层面，来进一步学习如何实现 RESTful API。

参考：

* HTTP请求方法：<https://www.w3schools.com/tags/ref_httpmethods.asp>
* Hypertext Transfer Protocol -- HTTP/1.0（RFC 1945）：<https://datatracker.ietf.org/doc/html/rfc1945#page-30>
* Hypertext Transfer Protocol -- HTTP/1.1（RFC 2616）：<https://datatracker.ietf.org/doc/html/rfc2616>
* Uniform Resource Identifier (URI): Generic Syntax（RFC 3986）：<https://datatracker.ietf.org/doc/html/rfc3986>
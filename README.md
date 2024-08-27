# Go_Web_Template

一个平平无奇的仅适用于 Java 选手转 Golang 初学者使用的 Golang web 模板



## 技术栈

**第三方框架**：

控制器层：Gin

数据存储层：GORM

日志：Zap



**特性**：

- 三层架构：controller - service - dao

- 自定义业务异常 (`businesserr`)
- Controller 通用返回结果 (`response`)
- 自定义 yml 配置文件
- 简单的 RBAC 权限控制 (`auth`)
- 集成 swag 接口文档
- 简单的日志封装，支持日志分级，磁盘持久化，按日期分割 (`logger`)

注：本模板不包含什么高级的技术，仅适用于学习食用，或者是用于快速开发简单的 web 应用



## config 包说明

模板中简单内置了几个必备的配置项，可根据个人需要进行补充，[参考文档](./config/README.md)

自定义配置项，需要同步维护好 `config` 包下的 `Config` 结构



## swagger 接口文档说明

文档通用信息注释：`internal/web/route.go`

接口代码发生更新时，需要手动在终端中执行一下文档更新命令，检查没有报错后再运行程序，方能查看和调试最新的接口

访问地址：`http://localhost:10086/api/swagger/index.html`

> /api 为示例程序使用的 context-path 配置，此处应根据实际配置值作出修改



文档更新命令（每次接口注释更新时都需要重新执行一遍）：

```shell
swag init -g internal/web/route.go --parseDependency
```

文档格式化命令：

```shell
swag fmt -d internal/controller
```

> 如果没有 swag 命令，需要在终端执行以下命令进行安装：
>
> ```shell
> go install github.com/swaggo/swag/cmd/swag@latest
> ```



通用 API 注释：

| 注释                     | 说明                                                         | 示例                                                         |
| ------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| title                    | **必填** 应用程序的名称。                                    | // @title Swagger Example API                                |
| version                  | **必填** 提供应用程序API的版本。                             | // @version 1.0                                              |
| description              | 应用程序的简短描述。                                         | // @description This is a sample server celler server.       |
| tag.name                 | 标签的名称。                                                 | // @tag.name This is the name of the tag                     |
| tag.description          | 标签的描述。                                                 | // @tag.description Cool Description                         |
| tag.docs.url             | 标签的外部文档的URL。                                        | // @tag.docs.url [https://example.com](https://example.com/) |
| tag.docs.description     | 标签的外部文档说明。                                         | // @tag.docs.description Best example documentation          |
| termsOfService           | API的服务条款。                                              | // @termsOfService http://swagger.io/terms/                  |
| contact.name             | 公开的API的联系信息。                                        | // @contact.name API Support                                 |
| contact.url              | 联系信息的URL。 必须采用网址格式。                           | // @contact.url http://www.swagger.io/support                |
| contact.email            | 联系人/组织的电子邮件地址。 必须采用电子邮件地址的格式。     | // @contact.email [support@swagger.io](mailto:support@swagger.io) |
| license.name             | **必填** 用于API的许可证名称。                               | // @license.name Apache 2.0                                  |
| license.url              | 用于API的许可证的URL。 必须采用网址格式。                    | // @license.url http://www.apache.org/licenses/LICENSE-2.0.html |
| host                     | 运行API的主机（主机名或IP地址）。                            | // @host localhost:8080                                      |
| BasePath                 | 运行API的基本路径。                                          | // @BasePath /api/v1                                         |
| accept                   | API 可以使用的 MIME 类型列表。 请注意，Accept 仅影响具有请求正文的操作，例如 POST、PUT 和 PATCH。 值必须如“[Mime类型](https://github.com/swaggo/swag/blob/master/README_zh-CN.md#mime类型)”中所述。 | // @accept json                                              |
| produce                  | API可以生成的MIME类型的列表。值必须如“[Mime类型](https://github.com/swaggo/swag/blob/master/README_zh-CN.md#mime类型)”中所述。 | // @produce json                                             |
| query.collection.format  | 请求URI query里数组参数的默认格式：csv，multi，pipes，tsv，ssv。 如果未设置，则默认为csv。 | // @query.collection.format multi                            |
| schemes                  | 用空格分隔的请求的传输协议。                                 | // @schemes http https                                       |
| externalDocs.description | Description of the external document.                        | // @externalDocs.description OpenAPI                         |
| externalDocs.url         | URL of the external document.                                | // @externalDocs.url https://swagger.io/resources/open-api/  |
| x-name                   | 扩展的键必须以x-开头，并且只能使用json值                     | // @x-example-key {"key": "value"}                           |

接口注释：

| 注释                 | 描述                                                         |
| -------------------- | ------------------------------------------------------------ |
| description          | 操作行为的详细说明。                                         |
| description.markdown | 应用程序的简短描述。该描述将从名为`endpointname.md`的文件中读取。 |
| id                   | 用于标识操作的唯一字符串。在所有API操作中必须唯一。          |
| tags                 | 每个API操作的标签列表，以逗号分隔。                          |
| summary              | 该操作的简短摘要。                                           |
| accept               | API 可以使用的 MIME 类型列表。 请注意，Accept 仅影响具有请求正文的操作，例如 POST、PUT 和 PATCH。 值必须如“[Mime类型](https://github.com/swaggo/swag/blob/master/README_zh-CN.md#mime类型)”中所述。 |
| produce              | API可以生成的MIME类型的列表。值必须如“[Mime类型](https://github.com/swaggo/swag/blob/master/README_zh-CN.md#mime类型)”中所述。 |
| param                | 用空格分隔的参数。`param name`,`param type`,`data type`,`is mandatory?`,`comment` `attribute(optional)` |
| security             | 每个API操作的[安全性](https://github.com/swaggo/swag/blob/master/README_zh-CN.md#安全性)。 |
| success              | 以空格分隔的成功响应。`return code`,`{param type}`,`data type`,`comment` |
| failure              | 以空格分隔的故障响应。`return code`,`{param type}`,`data type`,`comment` |
| response             | 与success、failure作用相同                                   |
| header               | 以空格分隔的头字段。 `return code`,`{param type}`,`data type`,`comment` |
| router               | 以空格分隔的路径定义。 `path`,`[httpMethod]`                 |
| deprecatedrouter     | 与router相同，但是是deprecated的。                           |
| x-name               | 扩展字段必须以`x-`开头，并且只能使用json值。                 |
| deprecated           | 将当前API操作的所有路径设置为deprecated                      |

MIME 类型：

| Alias                 | MIME Type                         |
| --------------------- | --------------------------------- |
| json                  | application/json                  |
| xml                   | text/xml                          |
| plain                 | text/plain                        |
| html                  | text/html                         |
| mpfd                  | multipart/form-data               |
| x-www-form-urlencoded | application/x-www-form-urlencoded |
| json-api              | application/vnd.api+json          |
| json-stream           | application/x-json-stream         |
| octet-stream          | application/octet-stream          |
| png                   | image/png                         |
| jpeg                  | image/jpeg                        |
| gif                   | image/gif                         |

参数类型：

- query
- path
- header
- body
- formData

数据类型：

- string (string)
- integer (int, uint, uint32, uint64)
- number (float32)
- boolean (bool)
- user defined struct

示例注释：

```go
// Login 用户登录
//	@Summary		用户登录
//	@Description	接收用户账号密码进行登录, 并保存登录态
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			userLoginDTO	body		dto.UserLogin	true	"用户登录参数"
//	@Success		200				{object}	response.R[vo.User]
//	@Failure		400				{object}	response.R[string]
//	@Router			/user/login [post]
func (u *UserController) Login(c *gin.Context) {
	params := new(dto.UserLogin)
	if err := c.ShouldBindJSON(params); err != nil {
		json(c, response.Error(businesserr.EnumParamsError))
		return
	}
	uvo, err := service.User().Login(c, params)
	if err != nil {
		json(c, response.NewByError(err))
		return
	}
	json(c, response.OkWithData(uvo))
}
```



## 待加入

1. Docker 部署配置
2. 更多的通用工具包
3. 定时任务
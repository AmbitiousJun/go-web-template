# 配置文件注释

## 文件分别有哪些？

主配置文件: `config.yml`

可在主配置中, 通过 `application.profiles` 指定若干个运行环境, 每个运行环境对应到一个单独的配置文件, 配置的优先级取决于环境名在 `application.profiles` 中的位置, 越靠右优先级越高

## 各个配置说明

| 一级配置    | 二级配置       | 可选值                                 | 示例值                                                       | 注释                                             |
| ----------- | :------------- | -------------------------------------- | ------------------------------------------------------------ | ------------------------------------------------ |
| application | name           | 任意字符串                             | `go_web_template`                                            | 应用程序名称，没有特殊作用                       |
| application | profiles       | 任意字符串，可配置多个，以“,”进行分隔  | `dev,redis`                                                  | 程序的运行环境，一般分为 `dev`、`prod` 和 `test` |
| server      | port           | 整型数，取值范围：`[1, 65535]`         | `10086`                                                      | 程序启动后监听的端口号                           |
| web         | context-path   | 以 `/` 分隔的字符串路径                | `/api`                                                       | web 上下文路径，即接口请求前缀                   |
| web         | cookie-max-age | 整型数，取值范围：`[0, math.MaxInt32]` | `2592000`                                                    | 保存用户登录态的 session 过期时间, 单位：秒      |
| log         | level          | `debug`、`info`、`error`               | `info`                                                       | 日志输出级别                                     |
| database    | dsn            | 任意字符串                             | `root:123456@tcp(127.0.0.1:3306)/my_db?charset=utf8mb4&parseTime=True&loc=Local` | 数据库连接信息                                   |


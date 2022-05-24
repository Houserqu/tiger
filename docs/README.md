# tiger

基于 golang 的 web 起始工程，封装了常用的基础能力，以及包含一个可配置化的管理后台，便于快速开始项目开发

## 目录结构

```go
.
├── Dockerfile
├── README.md
├── build.sh       // 构建脚本
├── config.yaml    // 系统配置（也可以放业务配置，修改实时生效）
├── core           // 脚手架框架代码
├── dev.sh         // 开发模式服务启动脚手架
├── docs           // 文档目录
├── logs           // 日志目录
├── go.mod
├── go.sum
├── main.go
├── middleware     // 中间件目录
├── module         // 业务代码模块目录
│   ├── router.go  // 路由定义文件
│   ├── user       // 用户模块
│   ├── view       // 视图模块，展示页面
│   └── ...
├── public         // 静态文件和 html 模板目录，访问路径 /public/
└── admin          // 管理后台源代码目录
```

## 开发

```bash
go run main.go # 直接启动服务
````

### 自动重启

使用 [air](https://github.com/cosmtrek/air/blob/master/README.md) 工具可以监听文件改动，自动重启服务，便于开发。

1. 安装 air 工具

```bash
# binary will be $(go env GOPATH)/bin/air
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

# or install it into ./bin/
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s

air -v
```

2. 启动 air

项目自带 air 配置文件（.air.toml），只需要在项目根目录运行 air

```bash
air       # 全局安装启动
./bin/air # 指定目录安装启动
```

## 特性

### 配置

使用 [viper](https://github.com/spf13/viper) 库来管理系统配置，同时开启 watch 模式（修改文件后 viper 能获取到最新的配置），默认加载项目根目录的 config.yaml 文件
config.yaml 配置文件不应该包含在版本库中

### 日志

使用 [logrus](https://github.com/sirupsen/logrus) 日志工具，支持记录请求ID，用于查看日志链路。

```golang
import 	"houserqu.com/tiger/core"

internal.Log(c).Info("123") // 由于无法获取协程上下文，所以需要显式的传递 gin.Context，才能记录 request id
```

### 数据库

使用 [gorm](https://gorm.io) 作为 ORM 库，支持多种类型的数据库，本项目默认使用 Mysql。

```bash
go get -u gorm.io/gorm
```

#### Mysql

```bash
go get -u github.com/go-sql-driver/mysql
```

## 部署

### Docker 部署

1. 构建镜像
```bash
docker build -t tiger:latest .
```

2. 启动容器
```bash
docker run -d --env-file prod.env -p 8090:8088 -v /Users/houserqu/tiger/logs:/app/logs tiger:latest
```

-p: 指定映射端口，容器内服务端口可以通过 config.yaml 文件配置  
-v: 指定日志文件主机挂载目录，容器内日志目录可以通过 config.yaml 文件配置

### 直接部署

```bash
go build main.go # 构建
./main           # 直接启动
pm2 start ./main # pm2 守护进程启动
```
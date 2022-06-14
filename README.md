# [tiger](https://houserqu.com/tiger/)

基于 golang 的 web 起始工程，封装了常用的基础能力，以及包含一个可配置化的管理后台，便于快速开始项目开发

## 特性

- 可动态配置页面
- 配置中心
- 角色-权限管理
- Docker 部署
- 文档工具 [docsify]((https://docsify.js.org/))

[查看文档 https://houserqu.com/tiger/](https://houserqu.com/tiger/)

## 开发

### 开发环境

- golang
- nodejs >= 12.0.0

### 初始化数据库

sql 文件路径 `./docs/init.sql`

### 安装依赖

```bash
# 安装 go mod
go mod download

# 安装 nodejs module
cd admin && yarn # 或者 cd admin && npm i

# 安装 docsify 工具
npm i docsify-cli -g # 或者 yarn global add docsify-cli
```

### 启动服务
```bash
# 启动 go 服务
go run main.go

# 启动管理后台开发模式
cd admin && npm run dev

# 启动文档预览服务
docsify serve docs
```

### 生成环境构建

```bash
go build --tags=production
```

生产环境构建会删除 swagger 相关功能

### 生成 swagger 文档

```bash
 swag init -p snakecase
```

### 规范

接口参数用蛇行命名法
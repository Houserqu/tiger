package main

import (
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/viper"
	"houserqu.com/gin-starter/core"
	_ "houserqu.com/gin-starter/docs"
	"houserqu.com/gin-starter/middleware"
	"houserqu.com/gin-starter/module/config"
	"houserqu.com/gin-starter/module/login"
	"houserqu.com/gin-starter/module/user"
	"houserqu.com/gin-starter/module/view"
)

var swagHandler gin.HandlerFunc

// @title API 文档
// @version 1.0
// https://swaggo.github.io/swaggo.io/

func main() {
	core.InitConfig()
	core.InitLogger()
	core.InitMysql()

	r := gin.New()

	// session
	store := cookie.NewStore([]byte(viper.GetString("session.key")))
	r.Use(sessions.Sessions("session", store))

	// 注册中间件
	r.Use(gin.Recovery())
	r.Use(middleware.Access())
	if viper.GetBool("cors.enable") {
		r.Use(middleware.Cors())
	}

	// 静态文件
	r.Static("/public", "./public")                            // 管理后台静态文件目录
	r.Static("/amis-editor-demo", "./static/amis-editor-demo") //
	// html 模板文件
	r.LoadHTMLGlob("./public/*.html")    // 管理后台 html 模板
	r.LoadHTMLGlob("./static/**/*.html") // 管理后台 html 模板

	// swagger
	if swagHandler != nil {
		r.GET("/swagger/*any", swagHandler)
	}

	// 注册每个模块中定义的路由
	view.Controller(r)
	config.Controller(r) // 系统配置
	user.Controller(r)   // 用户
	login.Controller(r)  // 登录注册

	// 监听端口
	err := r.Run(viper.GetString("server.addr"))
	if err != nil {
		log.Fatal("Error listen")
	}
}

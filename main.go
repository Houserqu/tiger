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
	"houserqu.com/gin-starter/module"
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
	// r.Use(middleware.CheckLogin())

	// 静态文件
	r.Static("/public", "./public")
	// html 模板文件
	r.LoadHTMLGlob("./public/*.html")

	// swagger
	if swagHandler != nil {
		r.GET("/swagger/*any", swagHandler)
	}

	// 注册路由
	module.InitRouter(r)

	// 监听端口
	err := r.Run(viper.GetString("server.addr"))
	if err != nil {
		log.Fatal("Error listen")
	}
}

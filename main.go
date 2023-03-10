package main

import (
	"flag"

	"github.com/gin-gonic/gin"
	"github.com/uerax/goconf"

	"github.com/uerax/chatgpt-prj/handler"
	"github.com/uerax/chatgpt-prj/handler/middleware"
	"github.com/uerax/chatgpt-prj/logger"
)

func Init() {

	var path string
	flag.StringVar(&path, "c", "etc", "config path")
	flag.Parse()

	err := goconf.LoadConfig(path)
	if err != nil {
		panic("配置文件读取失败")
	}

	logger.Init()
	middleware.LogInit()

	middleware.FilterInit()

	// db.Init()
}

func main() {
	
	Init()

	r := gin.New()

	r.Use(middleware.ZapLogger(), middleware.ZapRecovery())

	// r.Use(middleware.LoggerToFile())
	r.Use(middleware.FilterHandler())

	// 获取微信公众号access
	// wechat.Init()

	r.POST("/question", handler.AskHandler)

	// wx公众号
	// 必须结尾加/ 否则公众号无法正确请求
	r.GET("/wechat/", handler.WechatCheckHandler)
	r.POST("/wechat/", handler.WechatMessageHandler)
	
	r.Run(":8080") // 监听并在 0.0.0.0:80 上启动服务
}
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/uerax/goconf"

	"github.com/uerax/chatgpt-prj/handler"
	"github.com/uerax/chatgpt-prj/handler/middleware"
	"github.com/uerax/chatgpt-prj/logger"
)

func Init() {
	err := goconf.LoadConfig("etc")
	if err != nil {
		panic("配置文件读取失败")
	}

	logger.Init()

	middleware.FilterInit()
}


func main() {
	
	Init()

	r := gin.New()

	r.Use(middleware.ZapLogger(), middleware.ZapLoggerRec())

	// r.Use(middleware.LoggerToFile())
	r.Use(middleware.FilterHandler())

	// Recovery 中间件会 recover 任何 panic。如果有 panic 的话，会写入 500。
	r.Use(gin.Recovery())

	// err = wechat.Init()
	// if err != nil {
	// 	// TODO
	// 	os.Exit(1)
	// }

	r.POST("/question", handler.AskHandler)

	// wx公众号
	// 必须结尾加/ 否则公众号无法正确请求
	r.GET("/wechat/", handler.WechatCheckHandler)
	r.POST("/wechat/", handler.WechatMessageHandler)
	
	r.Run(":80") // 监听并在 0.0.0.0:80 上启动服务
}
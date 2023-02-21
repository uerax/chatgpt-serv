package main

import (
	"github.com/gin-gonic/gin"
	"github.com/uerax/goconf"

	"github.com/uerax/chatgpt-prj/global"
	"github.com/uerax/chatgpt-prj/handler"
)


func main() {

	err := goconf.LoadConfig("etc")
	if err != nil {
		panic("配置文件读取失败")
	}

	global.Init()

	r := gin.Default()

	r.Use(handler.MiddlewareHandler())

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
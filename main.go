package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/uerax/chatgpt-prj/handler"
	"github.com/uerax/chatgpt-prj/wechat"
)


func main() {
	r := gin.Default()

	err := wechat.Init()
	if err != nil {
		// TODO
		os.Exit(1)
	}

	r.POST("/question", handler.Question)
	r.GET("/wechat/", handler.WechatCheck)
	r.POST("/wechat/", handler.WechatMessage)
	
	r.Run(":80") // 监听并在 0.0.0.0:80 上启动服务
}
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/uerax/chatgpt-prj/handler"
)


func main() {
	r := gin.Default()
	r.GET("/question", handler.Question)
	r.GET("/wechat", handler.WechatCheck)
	
	r.Run(":80") // 监听并在 0.0.0.0:1919 上启动服务
}
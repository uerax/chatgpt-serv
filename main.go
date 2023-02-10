package main

import (
	"github.com/gin-gonic/gin"
	"github.com/uerax/chatgpt-prj/handler"
)


func main() {
	r := gin.Default()
	r.GET("/question", handler.Question)
	r.Run(":1919") // 监听并在 0.0.0.0:8080 上启动服务
}
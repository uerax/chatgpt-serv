package main

import (
	"github.com/gin-gonic/gin"
	"github.com/uerax/goconf"

	"github.com/uerax/chatgpt-prj/handler"
)


func main() {

	err := goconf.LoadConfig("etc")
	if err != nil {
		panic("配置文件读取失败")
	}

	r := gin.Default()

	// err = wechat.Init()
	// if err != nil {
	// 	// TODO
	// 	os.Exit(1)
	// }

	r.POST("/question", handler.Question)
	r.GET("/wechat/", handler.WechatCheck)
	r.POST("/wechat/", handler.WechatMessage)
	
	r.Run(":80") // 监听并在 0.0.0.0:80 上启动服务
}
package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/uerax/chatgpt-prj/global"
)

func MiddlewareHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		l := global.Irl.GetLimiter(c.ClientIP())
		if !l.Allow() {
			c.JSON(http.StatusTooManyRequests, nil)
			return
		}
		c.Next()
	}
}
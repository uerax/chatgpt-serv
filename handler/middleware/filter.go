package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/uerax/chatgpt-prj/common"
	"github.com/uerax/goconf"
)

var (
	irl *common.IpRateLimiter
)

func FilterInit() {
	r := goconf.VarIntOrDefault(1500, "filter", "ip", "rate")
	size := goconf.VarIntOrDefault(1500, "filter", "ip", "size")
	irl = common.NewIpRateLimiter(r, size)
}

func FilterHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO 可以加上一个全局限流器, 限制时间内最大访问数
		l := irl.GetLimiter(c.ClientIP())
		if !l.Allow() {
			c.JSON(http.StatusTooManyRequests, nil)
			return
		}
		c.Next()
	}
}
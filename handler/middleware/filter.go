package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"

	"github.com/uerax/chatgpt-prj/common"
	"github.com/uerax/goconf"
)

var (
	irl *common.IpRateLimiter
	Reqirl *rate.Limiter

)

func FilterInit() {
	r := goconf.VarIntOrDefault(5, "filter", "ip", "rate")
	size := goconf.VarIntOrDefault(5, "filter", "ip", "size")
	ReqR := goconf.VarIntOrDefault(10, "filter", "request", "rate")
	ReqSize := goconf.VarIntOrDefault(10, "filter", "request", "size")
	irl = common.NewIpRateLimiter(r, size)
	Reqirl = rate.NewLimiter(rate.Limit(ReqR), ReqSize)
}

func FilterHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		
		l := irl.GetLimiter(c.ClientIP())
		if !l.Allow() {
			c.JSON(http.StatusTooManyRequests, nil)
			return
		}

		if !Reqirl.Allow() {
			c.JSON(http.StatusTooManyRequests, nil)
			return
		}

		c.Next()
	}
}
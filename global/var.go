package global

import (
	"github.com/uerax/chatgpt-prj/common"
	"github.com/uerax/goconf"
)

var Token string
var Irl *common.IpRateLimiter

func Init() {
	r := goconf.VarIntOrDefault(1500, "filter", "ip", "rate")
	size := goconf.VarIntOrDefault(1500, "filter", "ip", "size")
	Irl = common.NewIpRateLimiter(r, size)
}

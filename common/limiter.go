package common

import (
	"sync"

	"golang.org/x/time/rate"
)

type IpRateLimiter struct {
	ips map[string]*rate.Limiter
	mx *sync.RWMutex
	r rate.Limit // 每秒填充速率
	size int // 令牌桶大小
}

func NewIpRateLimiter(r int, size int) *IpRateLimiter {
	return &IpRateLimiter{
		ips: make(map[string]*rate.Limiter),
		mx: &sync.RWMutex{},
		r: rate.Limit(r),
		size: size,
	}
}


func (t *IpRateLimiter) addIp(ip string) *rate.Limiter {
	t.mx.Lock()
	defer t.mx.Unlock()

	limiter := rate.NewLimiter(t.r, t.size)

	t.ips[ip] = limiter

	return limiter
}

func (t *IpRateLimiter) GetLimiter(ip string) *rate.Limiter {
	t.mx.Lock()

	if _, ok := t.ips[ip]; !ok {
		t.mx.Unlock()
		return t.addIp(ip)
	}

	defer t.mx.Unlock()

	return t.ips[ip]
}



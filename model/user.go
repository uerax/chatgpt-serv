package model

import "time"

type User struct {
	Id string `json:"id"` // id
	Expired time.Duration `json:"expired"` // 过期时间
	ReqCnt int `json:"request_count"` // 请求次数
	Free bool `json:"free"` // 是否免费
}
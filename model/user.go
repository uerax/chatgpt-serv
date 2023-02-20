package model

type User struct {
	Id string `json:"id"` // id
	AvailableTimes int `json:"available_times"` // 剩余请求次数
	Expired int64 `json:"ttl"` // vip过期时间
	Vip bool `json:"vip"` // 是否vip用户
	RegisterTime int64 `json:"register_time"` // 注册时间
}
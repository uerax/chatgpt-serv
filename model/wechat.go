package model

import "encoding/xml"

type AccessToken struct {
	Token  string `json:"access_token"`
	Expire int    `json:"expires_in"`
	Code   int    `json:"errcode"`
	Msg    string `json:"errmsg"` // 0 请求成功
}

type UserInfo struct {
	ToName     string   `xml:"ToUserName"`
	FromName   string   `xml:"FromUserName"`
	CreateTime int64    `xml:"CreateTime"`
	MsgType    string   `xml:"MsgType"`
	Content    string   `xml:"Content"`
	XMLName    xml.Name `xml:"xml"` // 若不标记XMLName, 则解析后的xml名为该结构体的名称
}

type ServiceMessage struct {
	ToUser  string             `json:"touser"`
	MsgType string             `json:"msgtype"`
	Text    ServiceMessageText `json:"text"`
}

type ServiceMessageText struct {
	Content string `json:"content"`
}

type TemplateMessage struct {
	ToUser string `json:"touser"`
	TmpId  string `json:"template_id"`
}

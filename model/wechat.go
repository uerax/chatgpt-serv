package model

type AccessToken struct {
	Token  string `json:"access_token"`
	Expire int    `json:"expires_in"`
}
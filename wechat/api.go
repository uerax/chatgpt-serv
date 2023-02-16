package wechat

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/uerax/chatgpt-prj/global"
	"github.com/uerax/chatgpt-prj/model"
)

func Init() error {
	token, err := getAccessToken()
	if err != nil {
		return err
	}
	global.Token = token
	return nil
}

func getAccessToken() (string, error) {
	// CONFIG
	appid := ""
	secret := ""
	url := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
	resp, err := http.Get(fmt.Sprintf(url, appid, secret))
	if err != nil {
		// TODO ADD LOG
		return "", fmt.Errorf("获取 Access Token 失败")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// TODO ADD LOG
		return "", fmt.Errorf("获取 Access Token 的 Json 解析失败")
	}

	accTk := &model.AccessToken{}

	err = json.Unmarshal(body, accTk)
	if err != nil {
		// TODO ADD LOG
		return "", fmt.Errorf("获取 Access Token 的 Json 解析失败")
	}

	if accTk.Code != 0 {
		return "", fmt.Errorf(accTk.Msg)
	}

	return accTk.Token, nil
}
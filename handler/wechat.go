package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/uerax/chatgpt-prj/model"
	"github.com/uerax/chatgpt-prj/util"
)

func WechatCheck(c *gin.Context) {
	// signature := c.Query("signature")
	// timestamp := c.Query("timestamp")
	// nonce := c.Query("nonce")

}

func checkSignature(signature, timestamp, nonce string) bool {
	token := ""

	tmpArr := []string{token, timestamp, nonce}

	sort.Strings(tmpArr)

	tmpStr := strings.Join(tmpArr, "")

	tmpStr = util.Sha1(tmpStr)

	if tmpStr == signature {
		return true
	} else {
		return false
	}

}

func getAccessToken() (string, error) {
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

	return accTk.Token, nil
}


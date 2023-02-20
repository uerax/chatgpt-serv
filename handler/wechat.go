package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/uerax/chatgpt-prj/chatgpt"
	"github.com/uerax/chatgpt-prj/global"
	"github.com/uerax/chatgpt-prj/model"
	"github.com/uerax/chatgpt-prj/util"
)

func WechatCheck(c *gin.Context) {
	signature := c.Query("signature")
	timestamp := c.Query("timestamp")
	nonce := c.Query("nonce")
	echostr := c.Query("echostr")
	if checkSignature(signature, timestamp, nonce) {
		c.String(http.StatusOK, echostr)
	} else {
		c.String(http.StatusBadRequest, "ERROR")
	}
}

func WechatMessage(c *gin.Context) {

	userInfo := &model.UserInfo{}

	err := c.BindXML(userInfo)
	if err != nil {
		// TODO ADD LOG
		return
	}
	resp := &model.UserInfo{
		FromName: userInfo.ToName,
		ToName: userInfo.FromName,
		MsgType: "text",
		CreateTime: time.Now().Unix(),
	}

	if userInfo.MsgType != "text" {
		resp.Content = "只支持文字类型提问"
		
		c.XML(http.StatusOK, resp)
		return
	}

	resp.Content = "服务器压力过大,请耐心等待 AI 回答"
	go sendMsg(userInfo.FromName, userInfo.Content)
	c.XML(http.StatusOK, resp)

}

func checkSignature(signature, timestamp, nonce string) bool {
	// CONFIG  url配置的token
	token := ""

	tmpArr := sort.StringSlice{token, timestamp, nonce}

	sort.Strings(tmpArr)

	tmpStr := strings.Join(tmpArr, "")

	tmpStr = util.Sha1(tmpStr)

	if tmpStr == signature {
		return true
	} else {
		return false
	}

}

func addCostmer() {

	type customer struct {
		Account string `json:"kf_account"`
      	Nickname string `json:"nickname"`
	}

	token := global.Token
	url := "https://api.weixin.qq.com/customservice/kfaccount/add?access_token=%s"

	cus := &customer{
		Account: "test@test",
		Nickname: "AI",
	}

	b, _ := json.Marshal(cus)

	http.Post(fmt.Sprintf(url, token), "", bytes.NewBuffer(b))
}

func sendMsg(touser, qst string) {
	token := global.Token
	url := "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=%s"
	
	ans := chatgpt.SendQuestion(qst)
	if ans == "" {
		return 
	}

	msg := model.ServiceMessage{
		ToUser: touser,
		Text: model.ServiceMessageText{
			Content: strings.Replace(ans, "\n\n", "", 1),
		},
	}

	msgReq, err := json.Marshal(msg)
	if err != nil {
		// TODO ADD LOG
		return
	}
	
	http.Post(fmt.Sprintf(url, token), "application/json", bytes.NewBuffer(msgReq))
	
}
package chatgpt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/uerax/chatgpt-prj/logger"
	"github.com/uerax/chatgpt-prj/model"
	"github.com/uerax/goconf"
)

type poll struct {
	mx sync.RWMutex
	cnt int
}

func New() *poll {
	return &poll{
		sync.RWMutex{}, 0,
	}
}

func (t *poll) index(len int) int {
	t.mx.Lock()
	defer t.mx.Unlock()
	t.cnt++
	return t.cnt % len
}

var p *poll

func init() {
	p = New()
}

func SendQuestion(qst string) string {
	log := logger.GetLogger()
	if len(qst) == 0 {
		return ""
	}
	url := "https://api.openai.com/v1/chat/completions"
	openai_key, err := goconf.VarArray("chatgpt", "key") // CONFIG
	if err != nil {
		log.Panic("获取chatgpt的key失败")
	}

	mode := "gpt-3.5-turbo"

	client := &http.Client{}
	reqMsgs := make([]model.GptMessage, 0)
	reqMsg := model.GptMessage{
		Role: "user",
		Content: qst,
	}
	reqMsgs = append(reqMsgs, reqMsg)

	reqBody := &model.GptQuestion{
		Model:       mode,
		Qst:	    reqMsgs,
	}

	b, err := json.Marshal(reqBody)
	if err != nil {
		log.Error(err.Error())
		return ""
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(b))
	if err != nil {
		log.Error(err.Error())
		return ""
	}
	key := fmt.Sprintf("%v", openai_key[p.index(len(openai_key))])
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+key)

	resp, err := client.Do(req)
	if err != nil {
		log.Error(err.Error())
		return ""
	}

	if resp.StatusCode != http.StatusOK {
		log.Error(err.Error())
		return "ChatGPT请求失败"
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err.Error())
		return ""
	}

	defer resp.Body.Close()

	ans := &model.GptAnswer{}
	err = json.Unmarshal(body, ans)
	if err != nil {
		log.Error(err.Error())
		return ""
	}

	return ans.Choices[0].Text.Content
}
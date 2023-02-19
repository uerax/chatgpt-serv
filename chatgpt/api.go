package chatgpt

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/uerax/chatgpt-prj/model"
	"github.com/uerax/goconf"
)

func SendQuestion(qst string) string {
	if len(qst) == 0 {
		return ""
	}
	url := "https://api.openai.com/v1/completions"
	openai_key, err := goconf.VarString("chatgpt", "key") // CONFIG
	if err != nil {
		return ""
	}
	mode := "text-davinci-003"
	max_tokens := 1000
	temperature := 0
	client := &http.Client{}
	reqBody := &model.GptQuestion{
		MaxTokens:   max_tokens,
		Model:       mode,
		Qst:         qst,
		Temperature: temperature,
	}

	b, err := json.Marshal(reqBody)
	if err != nil {
		// TODO ADD LOG
		return ""
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(b))
	if err != nil {
		// TODO ADD LOG
		return ""
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+openai_key)

	resp, err := client.Do(req)
	if err != nil {
		// TODO ADD LOG
		return ""
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// TODO ADD LOG
		return ""
	}

	defer resp.Body.Close()

	ans := &model.GptAnswer{}
	err = json.Unmarshal(body, ans)
	if err != nil {
		// TODO ADD LOG
		return ""
	}

	return ans.Choices[0].Text
}
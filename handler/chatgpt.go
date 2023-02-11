package handler

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/uerax/chatgpt-prj/model"

	"github.com/gin-gonic/gin"
)

func Question(c *gin.Context) {
	qst := c.Query("qst")

	c.JSON(200, gin.H{
		"status": http.StatusOK,
		"answer": sendQuestion(qst),
	})
}

func sendQuestion(qst string) string {
	if len(qst) == 0 {
		return ""
	}
	url := "https://api.openai.com/v1/completions"
	token := ""
	mode := "text-davinci-003"
	max_tokens := 1000
	temperature := 0
	client := &http.Client{}
	reqBody := &model.GptQuestion{
		MaxTokens: max_tokens,
		Model: mode,
		Qst: qst,
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
	req.Header.Add("Authorization", "Bearer " + token)

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

	return strings.TrimSpace(ans.Choices[0].Text)
}

package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

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
	url := ""
	token := ""
	mode := "text-davinci-003"
	max_tokens := "1000"
	temperature := "0"
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		// TODO ADD LOG
		return ""
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer " + token)

	req.PostForm.Set("model", mode)
	req.PostForm.Set("prompt", qst)
	req.PostForm.Set("max_tokens", max_tokens)
	req.PostForm.Set("temperature", temperature)

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

	ans := model.GptAnswer{}
	err = json.Unmarshal(body, ans)
	if err != nil {
		// TODO ADD LOG
		return ""
	}

	return ans.Ans
	

}

package handler

import (
	"net/http"
	"strings"

	"github.com/uerax/chatgpt-prj/chatgpt"
	"github.com/uerax/chatgpt-prj/model"

	"github.com/gin-gonic/gin"
)

func Question(c *gin.Context) {
	// qst := c.Query("qst")

	req := &model.GptReq{}

	err := c.BindJSON(req)
	if err != nil {
		// TODO ADD LOG
		c.JSON(500, gin.H{
			"status": http.StatusInternalServerError,
			"answer": "",
		})
		return
	}

	ans := "无访问权限, 请充值"

	if canQustion(req.Id) {
		ans = strings.Replace(chatgpt.SendQuestion(req.Qst), "\n\n", "", 1)
		
	} 

	c.JSON(200, gin.H{
		"status": http.StatusOK,
		"answer": ans,
	})
}



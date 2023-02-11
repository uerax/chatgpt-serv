package handler

import (
	"fmt"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestQuestion(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Question(tt.args.c)
		})
	}
}

func Test_sendQuestion(t *testing.T) {
	got := sendQuestion("这是一个测试")
	fmt.Println(got)
}

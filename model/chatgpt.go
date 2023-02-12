package model

type GptReq struct {
	Id string `json:"id"`
	Qst string `json:"question"`
}

type GptQuestion struct {
	Qst string `json:"prompt"`
	Model string `json:"model"`
	Temperature int `json:"temperature"`
	MaxTokens int `json:"max_tokens"`
}

type GptAnswer struct {
	ID      string   `json:"id"`     
	Object  string   `json:"object"` 
	Created int64    `json:"created"`
	Model   string   `json:"model"`  
	Choices []Choice `json:"choices"`
	Usage   Usage    `json:"usage"`  
}

type Choice struct {
	Text         string      `json:"text"`         
	Index        int64       `json:"index"`        
	Logprobs     interface{} `json:"logprobs"`     
	FinishReason string      `json:"finish_reason"`
}

type Usage struct {
	PromptTokens     int64 `json:"prompt_tokens"`    
	CompletionTokens int64 `json:"completion_tokens"`
	TotalTokens      int64 `json:"total_tokens"`     
}
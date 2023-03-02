package model

type GptReq struct {
	Id  string `json:"id"`
	Qst string `json:"question"`
}

type GptQuestion struct {
	Qst   []GptMessage `json:"messages"`
	Model string       `json:"model"`
}

type GptMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
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
	Text         Message `json:"message"`
	Index        int64   `json:"index"`
	FinishReason string  `json:"finish_reason"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Usage struct {
	PromptTokens     int64 `json:"prompt_tokens"`
	CompletionTokens int64 `json:"completion_tokens"`
	TotalTokens      int64 `json:"total_tokens"`
}

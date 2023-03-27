package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Temperature float64   `json:"temperature"`
}

type ChatResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Usage   struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Choices []struct {
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		FinishReason string `json:"finish_reason"`
		Index        int    `json:"index"`
	} `json:"choices"`
}

func main() {
	// 构造请求体
	chatReq := ChatRequest{
		Model:       "gpt-3.5-turbo",
		Temperature: 0.7,
		Messages: []Message{{
			Role:    "user",
			Content: "Hello, how are you?",
		}},
	}
	reqBody, err := json.Marshal(chatReq)
	if err != nil {
		panic(err)
	}

	// 发送POST请求
	url := "http://127.0.0.1:8080"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer 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")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 解析响应体
	var chatResp ChatResponse
	if err := json.NewDecoder(resp.Body).Decode(&chatResp); err != nil {
		panic(err)
	}

	// 输出结果
	for _, choice := range chatResp.Choices {
		fmt.Println(choice.Message.Content)
	}
}

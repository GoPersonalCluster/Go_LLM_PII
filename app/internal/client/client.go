package llm

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/GoPersonalCluster/go_llm_pii/app/internal/config"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
	Stream   bool      `json:"stream"`
}

type ChatResponse struct {
	Model   string `json:"model"`
	Message struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"message"`
}

type Client struct {
	BaseURL    string
	Model      string
	HTTPClient *http.Client
}

func NewClient() *Client {
	envConfig := config.NewEnvironmentConfig()

	return &Client{
		BaseURL: envConfig.LLMPIIHost,
		Model:   envConfig.LLMModel,
		HTTPClient: &http.Client{
			Timeout: 5 * time.Minute,
		},
	}
}

func (c *Client) Chat(ctx context.Context, prompt string) (string, error) {

	reqBody := ChatRequest{
		Model:  c.Model,
		Stream: false,
		Messages: []Message{
			{
				Role: "system",
				Content: `You are a PII detection engine.

Analyze the user's text and extract every Personally Identifiable Information (PII).
Rules:
- Respond ONLY with key=value pairs.
- One pair per line.
- Each Pair must contain a key identifier and its value.
- Do not include explanations.
- Do not include markdown.
- Do not include JSON.
- If multiple values exist for the same key, repeat the key.`,
			},
			{
				Role:    "user",
				Content: prompt,
			},
		},
	}

	body, err := json.Marshal(reqBody)

	if err != nil {
		println("Error marshalling request body:", err)
		return "error 100", err
	}
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		c.BaseURL+"/api/chat",
		bytes.NewBuffer(body),
	)

	if err != nil {
		println("Error creating request:", err)
		return "error 111", err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return "error 122", err
	}
	//defer resp.Body.Close()
	println("llm 85")
	if resp.StatusCode != http.StatusOK {

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}

		fmt.Println(string(data))

		var result ChatResponse
		if err := json.Unmarshal(data, &result); err != nil {
			return "", err
		}
	}
	data, err := io.ReadAll(resp.Body)
	println(json.Marshal(data))
	if err != nil {
		println("error")
		return "", err
	}

	fmt.Println(string(data))

	var result ChatResponse
	err = json.Unmarshal(data, &result)
	println(json.Marshal(data))

	// var result ChatResponse
	// if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {

	// 	return result.Message.Content, err
	// }
	println("llm data")
	return result.Message.Content, nil
}

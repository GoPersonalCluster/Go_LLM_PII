package llm

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
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

func NewClient(baseURL string) *Client {
	return &Client{
		BaseURL: baseURL,
		Model:   "gemma3:270m",
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
				Role:    "user",
				Content: prompt,
			},
		},
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		c.BaseURL+"/api/chat",
		bytes.NewBuffer(body),
	)

	if err != nil {
		println("Error creating request:", err)
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)

	if err != nil {
		return "", err
	}
	//defer resp.Body.Close()

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
	time.Sleep(2 * time.Second)
	println("llm data 100")
	data, err := io.ReadAll(resp.Body)
	println("llm data")
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

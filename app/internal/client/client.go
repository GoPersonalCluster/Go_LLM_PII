package llm

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
)

type Request struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}
type Response struct {
	Response string `json:"response"`
	Error    string `json:"error,omitempty"`
}

func Ask(prompt string) (string, error) {
	client := resty.New()

	reqBody := Request{
		Model:  "granite4:350m",
		Prompt: prompt,
		Stream: false,
	}

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(reqBody).
		Post("http://localhost:11434/api/generate")

	if err != nil {
		return "", err
	}

	var result Response
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return "", err
	}

	if result.Error != "" {
		return "", fmt.Errorf(result.Error)
	}

	return result.Response, nil
}

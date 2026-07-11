package llm

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClientChatServerError(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "internal error", http.StatusInternalServerError)
	}))

	defer server.Close()

	client := NewClient("http://127.0.0.1:11434/")

	answer, err := client.Chat(context.Background(), "My name is John Michael Smith. I was born on March 14, 1992. My email address is john.smith@example.com and my personal phone number is +1 (555) 123-4567. My home address is 1234 Maple Avenue, Springfield, IL 62704. My Social Security Number is 123-45-6789, and my driver's license number is D12345678. My passport number is X1234567. My credit card number is 4111 1111 1111 1111 with expiration date 12/29 and CVV 123. My bank account number is 9876543210 and routing number is 021000021. Please send all correspondence to john.smith@example.com.")
	println("llm answer :" + answer)
	if err == nil {
		t.Fatal("expected error")
	}
}
func TestClientChat(t *testing.T) {
	// Arrange
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			t.Fatalf("expected POST, got %s", r.Method)
		}

		if r.URL.Path != "/api/chat" {
			t.Fatalf("expected /api/chat, got %s", r.URL.Path)
		}

		var req ChatRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("invalid request: %v", err)
		}

		if req.Model != "gemma3:270m" {
			t.Fatalf("expected model gemma3:270m, got %s", req.Model)
		}

		if req.Stream {
			t.Fatal("expected stream=false")
		}

		if len(req.Messages) != 1 {
			t.Fatalf("expected 1 message, got %d", len(req.Messages))
		}

		if req.Messages[0].Role != "user" {
			t.Fatalf("expected role=user")
		}

		if req.Messages[0].Content != "Hello" {
			t.Fatalf("unexpected prompt")
		}

		resp := ChatResponse{
			Model: "gemma3:270m",
		}

		resp.Message.Role = "assistant"
		resp.Message.Content = "Hello back!"

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(resp); err != nil {
			t.Fatal(err)
		}
	}))

	defer server.Close()

	client := NewClient(server.URL)

	// Act
	answer, err := client.Chat(context.Background(), "Hello")

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if answer != "Hello back!" {
		t.Fatalf("expected 'Hello back!', got '%s'", answer)
	}
}

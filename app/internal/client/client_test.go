package llm

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClientChatServerError(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "internal error", http.StatusInternalServerError)
	}))

	defer server.Close()

	client := NewClient()

	answer, err := client.Chat(context.Background(), "My name is John Michael Smith. I was born on March 14, 1992. My email address is john.smith@example.com and my personal phone number is +1 (555) 123-4567. My home address is 1234 Maple Avenue, Springfield, IL 62704. My Social Security Number is 123-45-6789, and my driver's license number is D12345678. My passport number is X1234567. My credit card number is 4111 1111 1111 1111 with expiration date 12/29 and CVV 123. My bank account number is 9876543210 and routing number is 021000021. Please send all correspondence to john.smith@example.com.")

	println("llm answer :" + answer)
	if err == nil {
		t.Fatal("expected error")
	}
}

package openai_client

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChatGPTClient_PromptCompletions_Success(t *testing.T) {
	// Create a fake server to mock the OpenAI API response
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedURL := "/v1/chat/completions"
		assert.Equal(t, expectedURL, r.URL.Path)

		// Create a mock response
		responseRaw := `{
  "id": "chatcmpl-8fOq5PXHLG695gkTRoPoTEGUHTOwx",
  "object": "chat.completion",
  "created": 1704877621,
  "model": "gpt-3.5-turbo-0613",
  "choices": [
    {
      "index": 0,
      "message": {
        "role": "assistant",
        "content": "find . -type f -iname \"*.jpg\" -o -iname \"*.jpeg\" -o -iname \"*.png\" -o -iname \"*.gif\" -o -iname \"*.bmp\""
      },
      "logprobs": null,
      "finish_reason": "stop"
    }
  ],
  "usage": {
    "prompt_tokens": 65,
    "completion_tokens": 38,
    "total_tokens": 103
  },
  "system_fingerprint": null
}`
		// Convert mock response to JSON
		responseJSON := []byte(responseRaw)

		// Send the response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(responseJSON)
	}))

	defer server.Close()

	client := NewChatGPTClient(server.URL, "dummy_api_key")

	// Make a request
	message := "Find all images"
	result, err := client.PromptCompletions(message)

	// Verify the result
	assert.Nil(t, err)
	assert.Equal(
		t,
		"find . -type f -iname \"*.jpg\" -o -iname \"*.jpeg\" -o -iname \"*.png\" -o -iname \"*.gif\" -o -iname \"*.bmp\"",
		result,
	)
}

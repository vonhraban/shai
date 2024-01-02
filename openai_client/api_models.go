package openai_client

// ChatCompletion represents the chat completion response structure.
type ChatCompletion struct {
	ID                string      `json:"id"`
	Object            string      `json:"object"`
	Created           int64       `json:"created"`
	Model             string      `json:"model"`
	Choices           []Choice    `json:"choices"`
	Usage             Usage       `json:"usage"`
	SystemFingerprint interface{} `json:"system_fingerprint"`
}

// Choice represents a choice in the completion.
type Choice struct {
	Index        int      `json:"index"`
	Message      Message  `json:"message"`
	Logprobs     []string `json:"logprobs"`
	FinishReason string   `json:"finish_reason"`
}

// Message represents the role and content of a message.
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// Usage represents apiKey usage information.
type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// APIError represents the structure of the API error response.
type APIError struct {
	Error APIErrorDetails `json:"error"`
}

// APIErrorDetails represents the details of the API error.
type APIErrorDetails struct {
	Message string `json:"message"`
	Type    string `json:"type"`
	Param   string `json:"param"`
	Code    string `json:"code"`
}

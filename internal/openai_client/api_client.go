package openai_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ChatGPTClient struct {
	apiHost string
	apiKey  string
	debug   bool
}

const completionsUrl = "v1/chat/completions"
const model = "gpt-3.5-turbo"

func NewChatGPTClient(apiHost, apiKey string) *ChatGPTClient {
	return &ChatGPTClient{
		apiHost: apiHost,
		apiKey:  apiKey,
	}
}

func (c *ChatGPTClient) SetDebug(debug bool) {
	c.debug = debug
}

func (c *ChatGPTClient) PromptCompletions(message string) (string, error) {
	// Make a request to the ChatGPT3.5 API
	data := map[string]interface{}{
		"model": model,
		"messages": []interface{}{
			map[string]interface{}{"role": "system", "content": "" +
				"You are an unix command generation assitant. Output only code. Do not ouput" +
				"additional remarks about the code. Dot not output anything that is not a terminal command." +
				"Do not add any markdown formatting. The commands need to be compatible with zsh."},
			map[string]interface{}{"role": "user", "content": message}},
		"max_tokens": 150,
	}

	payload, _ := json.Marshal(data)
	body, err := c.post(payload)
	if err != nil {
		return "", err
	}

	if c.debug {
		fmt.Println("Response from ChatGPT 3.5 API:")
		fmt.Println(string(body))
	}

	var response ChatCompletion
	if err := json.Unmarshal(body, &response); err != nil {
		return "", fmt.Errorf("could not parse OpenAI completions api response: %w", err)
	}

	// if we unmarshalled into empty object, assume it is an error
	if response.Choices == nil {
		var errResponse APIError
		if err := json.Unmarshal(body, &errResponse); err != nil {
			return "", fmt.Errorf("could not parse OpenAI completions error api response: %w", err)
		}

		return "", fmt.Errorf("OpenAPI error %s: %s", errResponse.Error.Code, errResponse.Error.Message)
	}

	return response.Choices[0].Message.Content, nil
}

func (c *ChatGPTClient) post(payload []byte) ([]byte, error) {
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/%s", c.apiHost, completionsUrl), bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	defer resp.Body.Close()

	// Parse and print the response from the API
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	return body, nil
}

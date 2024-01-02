package main

import (
	"fmt"
	"github.com/vonhraban/shai/internal/openai_client"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	apiKey := os.Getenv("SHAI_OPENAI_API_KEY")
	// Set your conversation prompt here
	prompt := "What is square root of 49"

	client := openai_client.NewChatGPTClient(apiKey)
	res, err := client.PromptCompletions(prompt)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}

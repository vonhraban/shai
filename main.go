package main

import (
	"fmt"
	"github.com/vonhraban/shai/internal/openai_client"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	args := os.Args
	if len(args) == 1 {
		fmt.Printf("Usage: %s <plain text query | setup> \n", args[0])
		os.Exit(1)
	}

	input := strings.Join(args[1:], " ")

	if input == "setup" {
		fmt.Println("Setup wizard")
		return
	}

	apiKey := os.Getenv("SHAI_OPENAI_API_KEY")
	// Set your conversation prompt here
	//prompt := "What is square root of 49"
	prompt := input

	client := openai_client.NewChatGPTClient(apiKey)
	res, err := client.PromptCompletions(prompt)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}

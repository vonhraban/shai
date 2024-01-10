package main

import (
	"fmt"
	"github.com/vonhraban/shai/internal/openai_client"
	"os"
	"strings"
)

var config *Config

func init() {
	config = initConfig()
}

func main() {
	input := getInputFromArgs(os.Args)

	for input == "" {
		input = askPrompt("")
	}

	if input == "setup" {
		fmt.Println("Setup wizard")
		return
	}

	var client openai_client.Client
	if !config.DummyAPI {
		client = openai_client.NewChatGPTClient(config.APIHost, config.APIKey)
	} else {
		client = &openai_client.DummyClient{}
	}

	client.SetDebug(config.DebugMode)

	command := promptForInputInteractive(client, input)

	if command == "" {
		exitWithMessage("Cancelled")
	}

	printMessagef("Executing")

	executeCommand(command)
}

func getInputFromArgs(args []string) string {
	return strings.Join(args[1:], " ")
}

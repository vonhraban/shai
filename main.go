package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/vonhraban/shai/internal/openai_client"
	"os"
	"strings"
)

var debug bool
var apiKey string

func init() {
	godotenv.Load()

	apiKey = os.Getenv("SHAI_OPENAI_API_KEY")

	debugEnvVar := os.Getenv("SHAI_DEBUG")
	debug = debugEnvVar != "" && debugEnvVar != "0"
}

func main() {
	apiKey = os.Getenv("SHAI_OPENAI_API_KEY")
	debugEnvVar := os.Getenv("SHAI_DEBUG")
	if debugEnvVar != "" && debugEnvVar != "0" {
		debug = true
	}

	input := getInputFromArgs(os.Args)

	for input == "" {
		input = askPrompt("")
	}

	if input == "setup" {
		fmt.Println("Setup wizard")
		return
	}

	client := openai_client.NewChatGPTClient(apiKey)
	client.SetDebug(debug)

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

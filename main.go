package main

import (
	"bufio"
	"fmt"
	"github.com/vonhraban/shai/internal/openai_client"
	"os"
	"os/exec"
	"strings"

	"github.com/joho/godotenv"
	"github.com/manifoldco/promptui"
)

var debug bool
var apiKey string

func main() {
	godotenv.Load()
	apiKey = os.Getenv("SHAI_OPENAI_API_KEY")
	_, debug = os.LookupEnv("SHAI_DEBUG")

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

	client := openai_client.NewChatGPTClient(apiKey)
	client.SetDebug(debug)

	res, err := client.PromptCompletions(input)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
	prompt := promptui.Prompt{
		Label:     "Run command?",
		IsConfirm: true,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	if result == "y" {
		shell := os.Getenv("SHELL")

		cmd := exec.Command(shell, "-c", res)

		//Get a pipe to read from standard out
		r, _ := cmd.StdoutPipe()

		// Use the same pipe for standard error
		cmd.Stderr = cmd.Stdout

		// Make a new channel which will be used to ensure we get all output
		done := make(chan struct{})

		// Create a scanner which scans r in a line-by-line fashion
		scanner := bufio.NewScanner(r)

		// Use the scanner to scan the output line by line and log it
		// It's running in a goroutine so that it doesn't block
		go func() {

			// Read line by line and process it
			for scanner.Scan() {
				line := scanner.Text()
				fmt.Println(line)
			}

			// We're all done, unblock the channel
			done <- struct{}{}

		}()

		// Start the command and check for errors
		err := cmd.Start()
		if err != nil {
			fmt.Println("can not run command: %w", err)
		}

		// Wait for all output to be processed
		<-done

		// Wait for the command to finish
		err = cmd.Wait()
	}
}

package main

import (
	"bufio"
	"fmt"
	"github.com/vonhraban/shai/internal/openai_client"
	"os"
	"os/exec"
	"strings"

	"github.com/joho/godotenv"
	"github.com/pterm/pterm"
)

var debug bool
var apiKey string

func main() {
	godotenv.Load()
	apiKey = os.Getenv("SHAI_OPENAI_API_KEY")
	debugEnvVar := os.Getenv("SHAI_DEBUG")
	if debugEnvVar != "" && debugEnvVar != "0" {
		debug = true
	}

	args := os.Args

	input := strings.Join(args[1:], " ")

	for input == "" {
		input = askPrompt("")
	}

	if input == "setup" {
		fmt.Println("Setup wizard")
		return
	}

	client := openai_client.NewChatGPTClient(apiKey)
	client.SetDebug(debug)

	command := promptForCommand(client, input)

	if command == "" {
		pterm.Warning.Println("Cancelled")
		os.Exit(0)
	}

	pterm.Info.Printfln("Executing")

	executeCommand(command)
}

func promptForCommand(client *openai_client.ChatGPTClient, input string) string {
	// Create a multi printer. This allows multiple spinners to print simultaneously.
	multi := pterm.DefaultMultiPrinter
	writer := multi.NewWriter()
	// Create and start spinner 1 with a new writer from the multi printer.
	// The spinner will display the message "Spinner 1".
	spinner1, _ := pterm.DefaultSpinner.WithWriter(writer).Start("Loading")
	multi.Start()

	res, err := client.PromptCompletions(input)
	if err != nil {
		panic(err)
	}

	spinner1.Success("Completed")

	pterm.DefaultHeader.WithWriter(writer).WithMargin(15).WithBackgroundStyle(pterm.NewStyle(pterm.BgCyan)).WithTextStyle(pterm.NewStyle(pterm.FgBlack)).Println(res)

	// Stop the multi printer. This will stop printing all the spinners.
	multi.Stop()

	// Show an interactive confirmation dialog and get the result.
	//result, _ := pterm.DefaultInteractiveConfirm.Show()

	// Print a blank line for better readability.
	pterm.Println()

	options := []string{"Run", "Change Query", "Cancel"}

	selectedOption, _ := pterm.DefaultInteractiveSelect.
		WithDefaultText("Run the command?").
		WithDefaultOption("Yes").
		WithFilter(false).
		WithOptions(options).
		Show()

	switch selectedOption {
	case "Run":
		return res
	case "Change Query":
		input = askPrompt(input)
		return promptForCommand(client, input)
	}

	// Everything else - terminate
	return ""
}

func askPrompt(defaultValue string) string {
	textInput := pterm.DefaultInteractiveTextInput.WithMultiLine(false)

	// Show the text input and get the result
	input, _ := textInput.WithDefaultValue(defaultValue).Show()

	return input
}

func executeCommand(command string) {
	shell := os.Getenv("SHELL")

	cmd := exec.Command(shell, "-c", command)

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

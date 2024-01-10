package main

import (
	"github.com/pterm/pterm"
	"github.com/vonhraban/shai/internal/openai_client"
	"os"
)

func askPrompt(defaultValue string) string {
	textInput := pterm.DefaultInteractiveTextInput.WithMultiLine(false)

	// Show the text input and get the result
	input, _ := textInput.WithDefaultValue(defaultValue).Show()

	return input
}

func promptForInputInteractive(client openai_client.Client, input string) string {
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
	printNewLine()

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
		return promptForInputInteractive(client, input)
	}

	// Everything else - terminate
	return ""
}

func exitWithMessage(message string) {
	pterm.Warning.Println("Cancelled")
	os.Exit(0)
}

func printMessagef(format string, a ...interface{}) {
	pterm.Info.Printfln(format)
}

func printNewLine() {
	pterm.Println()
}

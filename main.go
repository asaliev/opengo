package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/asaliev/opengo/openai"
	"github.com/briandowns/spinner"
)

const (
	apiKeyName     string = "OPENAI_TOKEN"
	exitCommand    string = "exit"
	promptMessage  string = "Query (exit to quit): "
	goodbyeMessage string = "Goodbye..."
)

var queryPtr = flag.String("q", "", "Query sent to OpenAI")

func main() {
	flag.Parse()
	openai := openai.NewOpenaiProvider(apiKeyName)
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)

	// Get user input
	exit := false

	for !exit {
		if !isFlagSet("q") {
			fmt.Println("Query (`exit` to quit):")
		}

		input := getUserInput()
		if isExitCommand(input) {
			fmt.Println(goodbyeMessage)
			os.Exit(1)
		}

		// Show the spinner only in interactive mode
		if !isFlagSet("q") {
			s.Start()
		}

		// Contact OpenAI
		response, err := openai.Ask(input)
		if err != nil {
			panic(err.Error())
		}

		// Hide the spinner only in interactive mode
		if !isFlagSet("q") {
			s.Stop()
		}

		fmt.Println(response)

		// Exit when in non-interactive mode
		if isFlagSet("q") {
			break
		}
	}
}

func isExitCommand(input string) bool {
	return input == exitCommand
}

func getUserInput() string {
	if isFlagSet("q") {
		if *queryPtr == "" {
			panic("Error: empty query parameter")
		}
		return *queryPtr
	}

	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		if input.Text() != "" {
			break
		}
	}
	return input.Text()
}

func isFlagSet(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

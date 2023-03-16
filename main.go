package main

import (
	"bufio"
	"errors"
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
	for {
		query := ""
		if isFlagSet("q") {
			queryFromArgs, err := getUserQueryFromArgs()
			if err != nil {
				panic(err.Error())
			}
			query = queryFromArgs
		} else {
			fmt.Println("Query (`exit` to quit):")
			scanner := bufio.NewScanner(os.Stdin)
			query = getUserInput(scanner)
			if isExitCommand(query) {
				fmt.Println(goodbyeMessage)
				break
			}
		}

		if isExitCommand(query) {
			fmt.Println(goodbyeMessage)
			os.Exit(1)
		}

		// Show the spinner only in interactive mode
		if !isFlagSet("q") {
			s.Start()
		}

		// Contact OpenAI
		response, err := openai.Ask(query)
		if err != nil {
			if isFlagSet("q") {
				panic(err.Error())
			} else {
				s.Stop()
				fmt.Printf("ChatGPT error: %s", err.Error())
				continue
			}
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

func getUserQueryFromArgs() (string, error) {
	if *queryPtr == "" {
		return "", errors.New("empty -q parameter")
	}
	return *queryPtr, nil
}

func getUserInput(input *bufio.Scanner) string {
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

package main

import (
	"bufio"
	"flag"
	"os"

	"github.com/asaliev/opengo/command"
	"github.com/asaliev/opengo/openai"
)

const apiKeyName string = "OPENAI_TOKEN"

var queryPtr = flag.String("q", "", "Query sent to OpenAI")

func main() {
	flag.Parse()

	var provider command.Command = nil
	switch isFlagSet("q") {
	case false:
		scanner := bufio.NewScanner(os.Stdin)
		provider = command.NewInteractiveCommand(scanner)

	default:
		provider = command.NewSimpleCommand(queryPtr)
	}

	openai := openai.NewOpenaiChatgptProvider(apiKeyName)
	service := command.NewCommandService(provider, openai)
	service.Execute()
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

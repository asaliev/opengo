package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/asaliev/opengo/openai"
	"github.com/briandowns/spinner"
)

const apiKeyName string = "OPENAI_TOKEN"

func main() {
	// Get the users query via args or stdin
	openaiQueryPtr := flag.String("q", "", "Query sent to OpenAI")
	flag.Parse()
	if len(strings.TrimSpace(*openaiQueryPtr)) == 0 {
		fmt.Println("Please input a query for OpenAI or type `exit` to quit:")
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			if input.Text() == "exit" {
				fmt.Println("Goodbye...")
				break
			}
			if input.Text() == "" {
				continue
			}
			*openaiQueryPtr = input.Text()
			break
		}
	}

	// If we don't have a query at this stage we should just exit
	if *openaiQueryPtr == "" {
		os.Exit(0)
	}

	// Contact OpenAI
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Start()
	openai := openai.NewOpenaiProvider(apiKeyName)
	response, err := openai.Ask(openaiQueryPtr)
	if err != nil {
		fmt.Printf("\n%s\n", err.Error())
		os.Exit(1)
	}
	s.Stop()
	fmt.Println(response)
}

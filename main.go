package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/asaliev/opengo/config"
	"github.com/briandowns/spinner"
	openai "github.com/sashabaranov/go-openai"
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
	response, err := queryOpenAi(openaiQueryPtr)
	if err != nil {
		fmt.Printf("\n%s\n", err.Error())
		os.Exit(1)
	}
	s.Stop()
	fmt.Println(response)
}

func queryOpenAi(question *string) (string, error) {
	config := config.NewConfigProvider()
	client := openai.NewClient(config.ReadString(apiKeyName))
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: *question,
				},
			},
		},
	)

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}

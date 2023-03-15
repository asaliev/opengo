package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
)

var openAiToken string

func init() {
	openAiToken = getEnvVar("OPENAI_TOKEN")
}

func getEnvVar(key string) string {
	token := os.Getenv(key)
	if token == "" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error loading .env file")
		}
		token = os.Getenv(key)
	}

	return token
}

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
	client := openai.NewClient(openAiToken)
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
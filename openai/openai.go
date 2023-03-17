package openai

import (
	"context"

	"github.com/asaliev/opengo/config"
	"github.com/sashabaranov/go-openai"
)

type OpenaiProvider struct {
	client   *openai.Client
	messages []openai.ChatCompletionMessage
}

// How many context messages to send to ChatGPT
const contextNumberOfMessages int = 10

func NewOpenaiProvider(apiKeyName string) *OpenaiProvider {
	config := config.NewConfigProvider()
	client := openai.NewClient(config.ReadString(apiKeyName))
	return &OpenaiProvider{
		client: client,
		messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "You are a helpful assistant.",
			},
		},
	}
}

func (p *OpenaiProvider) Ask(question string) (string, error) {
	p.messages = append(p.messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: question,
	})

	// Probably need to look into garbage collection in the code below at some point
	length := len(p.messages)
	cut := contextNumberOfMessages
	if length < contextNumberOfMessages {
		cut = length
	}
	p.messages = p.messages[length-cut:]

	resp, err := p.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: p.messages,
		},
	)

	if err != nil {
		// Remove the last question from context since there was an error
		p.messages = p.messages[:len(p.messages)-1]

		return "", err
	}

	p.messages = append(p.messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: resp.Choices[0].Message.Content,
	})

	return resp.Choices[0].Message.Content, nil
}

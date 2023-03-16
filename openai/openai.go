package openai

import (
	"context"

	"github.com/asaliev/opengo/config"
	"github.com/sashabaranov/go-openai"
)

type openaiProvider struct {
	client *openai.Client
}

func NewOpenaiProvider(apiKeyName string) *openaiProvider {
	config := config.NewConfigProvider()
	client := openai.NewClient(config.ReadString(apiKeyName))
	return &openaiProvider{
		client: client,
	}
}

func (p *openaiProvider) Ask(question string) (string, error) {
	resp, err := p.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: question,
				},
			},
		},
	)

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}

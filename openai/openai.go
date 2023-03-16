package openai

import (
	"context"

	"github.com/asaliev/opengo/config"
	"github.com/sashabaranov/go-openai"
)

type openaiProvider struct {
	client   *openai.Client
	messages []openai.ChatCompletionMessage
}

// How many context messages to send to ChatGPT
const contextNumberOfMessages int = 5

func NewOpenaiProvider(apiKeyName string) *openaiProvider {
	config := config.NewConfigProvider()
	client := openai.NewClient(config.ReadString(apiKeyName))
	return &openaiProvider{
		client: client,
		messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "You are a helpful assistant.",
			},
		},
	}
}

func (p *openaiProvider) Ask(question string) (string, error) {
	p.messages = append(p.messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: question,
	})

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
		return "", err
	}

	p.messages = append(p.messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: resp.Choices[0].Message.Content,
	})

	return resp.Choices[0].Message.Content, nil
}

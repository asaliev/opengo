package command

import "github.com/asaliev/opengo/openai"

type Command interface {
	Run(openai.OpenaiProvider)
}

type CommandService struct {
	command Command
	openai  openai.OpenaiProvider
}

func NewCommandService(c Command, o openai.OpenaiProvider) *CommandService {
	return &CommandService{
		command: c,
		openai:  o,
	}
}

func (cs *CommandService) Execute() {
	cs.command.Run(cs.openai)
}

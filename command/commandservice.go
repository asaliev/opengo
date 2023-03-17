package command

import "github.com/asaliev/opengo/openai"

type Command interface {
	Run(openai.OpenaiApi)
}

type CommandService struct {
	command Command
	openai  openai.OpenaiApi
}

func NewCommandService(c Command, o openai.OpenaiApi) *CommandService {
	return &CommandService{
		command: c,
		openai:  o,
	}
}

func (cs *CommandService) Execute() {
	cs.command.Run(cs.openai)
}

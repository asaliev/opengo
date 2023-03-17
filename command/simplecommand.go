package command

import (
	"errors"
	"fmt"

	"github.com/asaliev/opengo/openai"
)

type SimpleCommand struct {
	query *string
}

func NewSimpleCommand(query *string) *SimpleCommand {
	return &SimpleCommand{
		query: query,
	}
}

func (p *SimpleCommand) Run(openai openai.OpenaiApi) {
	err := p.validateQuery()
	if err != nil {
		panic(err.Error())
	}

	response, err := openai.Ask(*p.query)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(response)
}

func (sc *SimpleCommand) validateQuery() error {
	if *sc.query == "" {
		return errors.New("empty -q parameter")
	}
	return nil
}

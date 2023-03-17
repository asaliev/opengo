package command

import (
	"bufio"
	"errors"
	"fmt"
	"time"

	"github.com/asaliev/opengo/openai"
	"github.com/briandowns/spinner"
)

const (
	exitCommand    string = "exit"
	promptMessage  string = "Query (exit to quit): "
	goodbyeMessage string = "Goodbye..."
)

type InteractiveCommand struct {
	scanner *bufio.Scanner
}

func NewInteractiveCommand(scanner *bufio.Scanner) *InteractiveCommand {
	return &InteractiveCommand{
		scanner: scanner,
	}
}

func (c *InteractiveCommand) Run(openai openai.OpenaiApi) {
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)

	for {
		// Get user input
		fmt.Println(promptMessage)
		query, err := c.getUserInput()
		if err != nil {
			continue // empty input
		}

		if c.isExitCommand(query) {
			fmt.Println(goodbyeMessage)
			break
		}

		s.Start()
		response, err := openai.Ask(query)
		s.Stop()
		if err != nil {
			fmt.Printf("ChatGPT error: %s", err.Error())
			continue
		}

		fmt.Println(response)
	}
}

func (c *InteractiveCommand) isExitCommand(input string) bool {
	return input == exitCommand
}

func (c *InteractiveCommand) getUserInput() (string, error) {
	c.scanner.Scan()
	if c.scanner.Text() == "" {
		return "", errors.New("empty input from user")
	}
	return c.scanner.Text(), nil
}

package command

import (
	"testing"

	"github.com/asaliev/opengo/openai"
)

type MockFooCommand struct{}

func NewMockFooCommand() *MockFooCommand {
	return &MockFooCommand{}
}

var testFlag = false

func (p *MockFooCommand) Run(api openai.OpenaiApi) {
	testFlag = true
}

type MockOpenaiApi struct{}

func NewMockOpenaiApi() *MockOpenaiApi {
	return &MockOpenaiApi{}
}

func (p *MockOpenaiApi) Ask(question string) (string, error) {
	return "foobar", nil
}

func TestCommandServiceRunsOpenaiProvider(t *testing.T) {
	mockCommand := NewMockFooCommand()
	mockApi := NewMockOpenaiApi()

	service := NewCommandService(mockCommand, mockApi)

	if testFlag {
		t.Fatal("testFlag = true; want: false")
	}
	service.Execute()
	if !testFlag {
		t.Fatal("testFlag = false; want: true")
	}
}

package command

import (
	"testing"
)

func TestGetUserQueryFromArgs_Success(t *testing.T) {
	expected := "To be or not to be?"
	command := NewSimpleCommand(&expected)
	err := command.validateQuery()

	if err != nil {
		t.Fatal("err != nil; want: nil")
	}
}

func TestGetUserQueryFromArgs_WhenPassingEmptyArgument(t *testing.T) {
	expected := ""
	command := NewSimpleCommand(&expected)
	err := command.validateQuery()

	if err == nil {
		t.Fatal("err = nil; want: not nil")
	}

	if err.Error() != "empty -q parameter" {
		t.Fatalf("err = %s, want: empty -q parameter", err.Error())
	}
}

package command

import (
	"bufio"
	"strings"
	"testing"
)

func TestIsExitCommand_Success(t *testing.T) {
	scanner := bufio.NewScanner(strings.NewReader(""))
	provider := NewInteractiveCommand(scanner)

	got := provider.isExitCommand("exit")
	if !got {
		t.Fatalf("isExitCommand(\"exit\") = false; want: true")
	}
}

func TestIsExitCommand_Failure(t *testing.T) {
	scanner := bufio.NewScanner(strings.NewReader(""))
	provider := NewInteractiveCommand(scanner)

	got := provider.isExitCommand("bad")
	if got {
		t.Fatalf("isExitCommand(\"bad\") = true; want: false")
	}
}

func TestGetUserInput_Success(t *testing.T) {
	scanner := bufio.NewScanner(strings.NewReader("foobar"))
	provider := NewInteractiveCommand(scanner)

	got, err := provider.getUserInput()
	if err != nil {
		t.Fatalf("err: %s; want: nil", err.Error())
	}

	if got != "foobar" {
		t.Fatalf("got: %s; want: foobar", got)
	}
}

func TestGetUserInput_EmptyInputFailure(t *testing.T) {
	scanner := bufio.NewScanner(strings.NewReader(""))
	provider := NewInteractiveCommand(scanner)

	got, err := provider.getUserInput()
	if err == nil {
		t.Fatal("err: nil; want: not nil")
	}
	if err.Error() != "empty input from user" {
		t.Fatalf("err: %s; want: empty input from user", err.Error())
	}

	if got != "" {
		t.Fatalf("got: %s; want: [empty]", got)
	}
}

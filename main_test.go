package main

import (
	"flag"
	"os"
	"testing"
)

func TestIsFlagSet_Missing(t *testing.T) {
	got := isFlagSet("foobar")
	if got {
		t.Fatal("isFlagSet(\"foobar\") = true; want: false")
	}
}

func TestIsFlagSet_Exists(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	flag.String("foobar", "", "Foo Bar")
	os.Args = []string{oldArgs[0], "-foobar", "baz"}
	flag.Parse()

	got := isFlagSet("foobar")
	if !got {
		t.Fatal("isFlagSet(\"foobar\") = false; want: true")
	}
}

func TestIsExitCommand_True(t *testing.T) {
	got := isExitCommand("exit")
	if !got {
		t.Fatalf("isExitCommand(\"exit\") = false; want: true")
	}
}

func TestIsExitCommand_False(t *testing.T) {
	got := isExitCommand("hello")
	if got {
		t.Fatalf("isExitCommand(\"exit\") = true; want: false")
	}
}

func TestGetUserQueryFromArgs_Success(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	expected := "To be or not to be?"
	os.Args = []string{oldArgs[0], "-q", expected}
	flag.Parse()

	got, err := getUserQueryFromArgs()
	if err != nil {
		t.Fatalf("err != nil; want: nil")
	}
	if got != expected {
		t.Fatalf("got = %s; want: %s", got, expected)
	}
}

func TestGetUserQueryFromArgs_WhenPassingEmptyArgument(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{oldArgs[0], "-q", ""} // pass empty arg
	flag.Parse()

	got, err := getUserQueryFromArgs()
	if got != "" {
		t.Fatal("got is not empty string; want: empty")
	}
	if err == nil {
		t.Fatalf("err is nil; want: not nil")
	}
}

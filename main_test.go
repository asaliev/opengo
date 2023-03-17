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

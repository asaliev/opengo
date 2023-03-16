package config

import (
	"os"
	"testing"
)

func TestReadStringFromEnvVar(t *testing.T) {
	envValue := "foobar"
	os.Setenv("OPENGO_OPENAI_TOKEN", envValue)

	config := NewConfigProvider()

	got := config.ReadString("OPENAI_TOKEN")
	if got != envValue {
		t.Fatalf("config.ReadString(\"OPENAI_TOKEN\") = %s; want: foobar", got)
	}
}

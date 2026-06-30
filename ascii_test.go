package main

import (
	"os"
	"testing"
)

func TestValidate(t *testing.T) {
	os.Args = []string{"cmd", "Hello", "standard"}
	input, banner := validate()

	if input != "Hello" {
		t.Errorf("expected Hello, got %s", input)
	}
	if banner != "standard" {
		t.Errorf("expected standard, got %s", banner)
	}
}

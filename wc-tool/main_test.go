package main

import (
	"fmt"
	"testing"
)

func TestCountBytes(t *testing.T) {
	input := []byte("Hello, World!")
	expected := 13
	result := CountBytes(input)

	if result != expected {
		t.Errorf("CountBytes(%q) = %d, want %d", input, result, expected)
	} else {
		t.Logf("CountBytes(%q) = %d, want %d", input, result, expected)
	}
}

func TestCountLines(t *testing.T) {
	inputLine1 := []byte("Hello, World!This is a new line")
	inputLine2 := []byte("This is a new line")

	input := []byte(fmt.Sprintf("%s\n %s\n", inputLine1, inputLine2))

	expected := 2
	result := CountLines(input)

	if result != expected {
		t.Errorf("CountLines(%q) = %d, want %d", input, result, expected)
	} else {
		t.Logf("CountLines(%q) = %d, want %d", input, result, expected)
	}
}

func TestCountWords(t *testing.T) {
	input := []byte("Hello, World!")
	expected := 2
	result := CountWords(input)

	if result != expected {
		t.Errorf("CountWords(%q) = %d, want %d", input, result, expected)
	} else {
		t.Logf("CountWords(%q) = %d, want %d", input, result, expected)
	}
}

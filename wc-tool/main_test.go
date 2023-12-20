package main

import (
	"os"
	"testing"
)

// Expected outputs are taken from https://codingchallenges.fyi/challenges/challenge-wc/

var fileName string = "test.txt"

func TestCountBytes(t *testing.T) {
	input, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	expected := 342190
	result := CountBytes(input)

	if result != expected {
		t.Errorf("CountBytes(%q) = %d, exexted %d", input[:10], result, expected)
	} else {
		t.Logf("CountBytes(%q) = %d, exexted %d", input[:10], result, expected)
	}
}

func TestCountLines(t *testing.T) {
	input, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	expected := 7145
	result := CountLines(input)

	if result != expected {
		t.Errorf("CountLines(%q) = %d, exexted %d", input[:10], result, expected)
	} else {
		t.Logf("CountLines(%q) = %d, exexted %d", input[:10], result, expected)
	}
}

func TestCountWords(t *testing.T) {
	input, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	expected := 58164
	result := CountWords(input)

	if result != expected {
		t.Errorf("CountWords(%q) = %d, exexted %d", input[:10], result, expected)
	} else {
		t.Logf("CountWords(%q) = %d, exexted %d", input[:10], result, expected)
	}
}

	} else {
		t.Logf("CountWords(%q) = %d, want %d", input, result, expected)
	}
}

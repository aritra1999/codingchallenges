package main

import "testing"

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


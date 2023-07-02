package main

import "testing"

func TestHelloWorld(t *testing.T) {
	expected := "Hello, World!"
	if got := hello_world(); got != expected {
		t.Errorf("hello_world() = %q, want %q", got, expected)
	}
}


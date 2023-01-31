package main

import "testing"

func TestGreet(t *testing.T) {
	greeting, greetingLength := greet("Alice")

	// Test assertion to check the returned greeting string.
	if greeting != "Hello Alice" {
		t.Errorf("want: %s; got: %s", "Hello Alice", greeting)
	}

	// Test assertion to check the returned greeting length.
	if greetingLength != 11 {
		t.Errorf("want: %d; got: %d", 11, greetingLength)
	}
}

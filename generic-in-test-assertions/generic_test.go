package main

import (
	"./assert"
	"testing"
)

func TestGreetGeneric(t *testing.T) {
	greeting, greetingLength := greet("Alice")

	assert.Equal(t, "Hello Alice", greeting)
	assert.Equal(t, 11, greetingLength)
}

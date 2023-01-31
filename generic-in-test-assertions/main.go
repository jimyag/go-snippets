package main

import "fmt"

func greet(name string) (string, int) {
	greeting := fmt.Sprintf("Hello %s", name)

	// Return the greeting and its length (in bytes).
	return greeting, len(greeting)
}

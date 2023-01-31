package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	urlA, err := url.Parse("https://example.com/foo")
	if err != nil {
		log.Fatal(err)
	}

	urlB := *urlA

	urlB.Path = "/bar"

	fmt.Printf("urlA: %s\n", urlA.String())
	fmt.Printf("urlB: %s\n", urlB.String())
}

package main

import (
	"fmt"
	"log"
	"net/url"
	"strconv"
)

func main() {
	// Use url.Parse() to parse a string into a *url.URL type. If your URL is
	// already a url.URL type you can skip this step.
	urlA, err := url.Parse("https://example.com?name=alice&age=28&gender=female")
	if err != nil {
		log.Fatal(err)
	}

	// Make a copy of the original url.URL.
	urlB := *urlA

	// Make the param changes to the new url.URL type...
	values := urlB.Query()

	values.Add("occupation", "carpenter")
	values.Del("gender")
	values.Set("age", strconv.Itoa(29))

	urlB.RawQuery = values.Encode()

	fmt.Printf("urlA: %s\n", urlA.String())
	fmt.Printf("urlB: %s\n", urlB.String())
}

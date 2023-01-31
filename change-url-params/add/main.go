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

	// Use the Query() method to get the query string params as a url.Values map.
	values := urlA.Query()

	// Make the changes that you want using the Add(), Set() and Del() methods. If
	// you want to retrieve or check for a specific parameter you can use the Get()
	// and Has() methods respectively.
	values.Add("occupation", "carpenter")
	values.Del("gender")
	values.Set("age", strconv.Itoa(29))

	// Use the Encode() method to transform the url.Values map into a URL-encoded
	// string (like "age=29&name=alice...") and assign it back to the URL. Note
	// that the encoded values will be sorted alphabetically based on the parameter
	// name
	urlA.RawQuery = values.Encode()

	fmt.Printf("urlA: %s", urlA.String())
}

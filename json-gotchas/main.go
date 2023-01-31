package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"strings"
	"time"
)

func integerExample() {
	m := map[int]string{
		123:     "foo",
		456_000: "bar",
	}

	data, _ := json.Marshal(m)
	fmt.Println(string(data))
}

func timeExample() {
	t1 := time.Now()
	t2 := t1.Add(24 * time.Hour)

	m := map[time.Time]string{
		t1: "foo",
		t2: "bar",
	}

	data, _ := json.Marshal(m)
	fmt.Println(string(data))
}

func escapeExample() {
	s := []string{
		"<foo>",
		"bar & baz",
	}

	data, _ := json.Marshal(s)
	fmt.Println(string(data))
}

func floatExample() {
	s := []float64{
		123.0,
		456.100,
		789.990,
	}

	data, _ := json.Marshal(s)
	fmt.Println(string(data))
}

func omitemptyExample1() {
	s := struct {
		Foo struct {
			Bar string `json:",omitempty"`
		} `json:",omitempty"`
	}{}

	data, _ := json.Marshal(s)
	fmt.Println(string(data))
}

func omitemptyExample2() {
	s := struct {
		Foo *struct {
			Bar string `json:",omitempty"`
		} `json:",omitempty"`
	}{}

	data, _ := json.Marshal(s)
	fmt.Println(string(data))
}

func zeroValueTimeExample() {
	s := struct {
		Foo time.Time `json:",omitempty"`
	}{}

	data, _ := json.Marshal(s)
	fmt.Println(string(data))
}

func stringStructTagExample() {
	s := struct {
		Foo int `json",string"`
	}{
		Foo: 123,
	}

	data, _ := json.Marshal(s)
	fmt.Println(string(data))
}

func nonASCIIExample() {
	s := struct {
		CostUSD string `json:"cost $"` // OK
		CostEUR string `json:"cost €"` // Contains the non-ASCII punctuation character €. Will be ignored.
	}{
		CostUSD: "100.00",
		CostEUR: "100.00",
	}

	data, _ := json.Marshal(s)
	fmt.Println(string(data))
}

func decodeNonASCIIExample1() {
	js := []byte(`{"cost $":"100.00","cost €":"100.00"}`)

	s := struct {
		CostUSD string `json:"cost $"` // OK
		CostEUR string `json:"cost €"` // Contains the non-ASCII punctuation character €. Will be ignored.
	}{}

	err := json.Unmarshal(js, &s)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", s)
}

func decodeNonASCIIExample2() {
	js := []byte(`{"cost $":"100.00","cost €":"100.00"}`)

	var aux map[string]string

	err := json.Unmarshal(js, &aux)
	if err != nil {
		log.Fatal(err)
	}

	s := struct {
		CostUSD string `json:"cost $"`
		CostEUR string `json:"cost €"`
	}{
		CostUSD: aux["cost $"],
		CostEUR: aux["cost €"],
	}

	fmt.Printf("%+v\n", s)
}

func jsonNumberExample() {
	js := `{"foo":123,"bar":true}`

	var m map[string]interface{}

	dec := json.NewDecoder(strings.NewReader(js))
	dec.UseNumber()

	err := dec.Decode(&m)
	if err != nil {
		log.Fatal(err)
	}

	i, err := m["foo"].(json.Number).Int64()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("foo: %d\n", i)
}

func moreExample1() {
	js := `{"name":"alice"}{"name":"bob"}]`

	dec := json.NewDecoder(strings.NewReader(js))
	for {
		var user map[string]string

		err := dec.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v\n", user)

		// Don't do this!
		if !dec.More() {
			break
		}
	}
}

func moreExample2() {
	js := `{"name":"alice"}{"name":"bob"}]`

	dec := json.NewDecoder(strings.NewReader(js))
	for {
		var user map[string]string

		err := dec.Decode(&user)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			log.Fatal(err)
		}

		fmt.Printf("%v\n", user)
	}
}

type Age int

func (age Age) MarshalJSON() ([]byte, error) {
	encodedAge := fmt.Sprintf("%d years", age)
	//encodedAge = strconv.Quote(encodedAge)
	return []byte(encodedAge), nil
}

func marshalJSONExample() {
	users := map[string]Age{
		"alice": 21,
		"bob":   84,
	}

	js, err := json.Marshal(users)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", js)
}

func main() {
	integerExample()
	timeExample()
	escapeExample()
	floatExample()
	omitemptyExample1()
	omitemptyExample2()
	zeroValueTimeExample()
	stringStructTagExample()
	nonASCIIExample()
	decodeNonASCIIExample1()
	decodeNonASCIIExample2()
	jsonNumberExample()
	moreExample1()
	//moreExample2()
	marshalJSONExample()
}

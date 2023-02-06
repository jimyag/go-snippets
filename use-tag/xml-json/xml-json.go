// Go offers built-in support for XML and XML-like
// formats with the `encoding.xml` package.

package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"

	"github.com/jimyag/use-tag/models"
)

var (
	coffee           *models.Plant
	coffeeWithoutXML *models.PlantWithoutXMLName
	coffeeWitXML     *models.PlantWithXMLName
)

func init() {
	coffee = &models.Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}
	coffeeWithoutXML = &models.PlantWithoutXMLName{Id: 27, Name: "Coffee"}
	coffeeWithoutXML.Origin = []string{"Ethiopia", "Brazil"}
	coffeeWitXML = &models.PlantWithXMLName{Id: 27, Name1: "Coffee"}
	coffeeWitXML.Origin = []string{"Ethiopia", "Brazil"}
}

func xmlExample() {
	// Emit XML representing our plant; using
	// `MarshalIndent` to produce a more
	// human-readable output.
	out, _ := xml.MarshalIndent(coffee, " ", "  ")
	fmt.Println(string(out))

	// To add a generic XML header to the output, append
	// it explicitly.
	fmt.Println(xml.Header + string(out))

	// Use `Unmarshal` to parse a stream of bytes with XML
	// into a data structure. If the XML is malformed or
	// cannot be mapped onto Plant, a descriptive error
	// will be returned.
	var p models.Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)

	tomato := &models.Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}

	// The `parent>child>plant` field tag tells the encoder
	// to nest all `plant`s under `<parent><child>...`
	type Nesting struct {
		XMLName xml.Name        `xml:"nesting"`
		Plants  []*models.Plant `xml:"parent>child>plant"`
	}

	nesting := &Nesting{}
	nesting.Plants = []*models.Plant{coffee, tomato}

	out, _ = xml.MarshalIndent(nesting, " ", "  ")
	fmt.Println(string(out))
}

func jsonExample() {
	out, _ := json.Marshal(coffeeWithoutXML)
	fmt.Println(string(out))
	out, _ = json.Marshal(coffeeWitXML)
	fmt.Println(string(out))
	// {"id":27,"name":"Coffee","origin":["Ethiopia","Brazil"]}
	// {"id":27,"name":"Coffee","origin":["Ethiopia","Brazil"]}
}

func main() {
	// xmlExample()
	jsonExample()
}

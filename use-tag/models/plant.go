package models

import (
	"encoding/xml"
	"fmt"
)

// Plant will be mapped to XML. Similarly to the
// JSON examples, field tags contain directives for the
// encoder and decoder. Here we use some special features
// of the XML package: the `XMLName` field name dictates
// the name of the XML element representing this struct;
// `id,attr` means that the `Id` field is an XML
// _attribute_ rather than a nested element.
type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id      int      `xml:"id" json:"id"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v",
		p.Id, p.Name, p.Origin)
}

type PlantWithoutXMLName struct {
	Id     int      `xml:"id" json:"id"`
	Name   string   `xml:"name" json:"name"`
	Origin []string `xml:"origin" json:"origin"`
}

type PlantWithXMLName struct {
	xml.Name `json:"-"`
	Id       int      `xml:"id" json:"id"`
	Name1    string   `xml:"name" json:"name"`
	Origin   []string `xml:"origin" json:"origin"`
}

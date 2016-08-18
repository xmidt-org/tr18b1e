package tr18b1e

import (
	"encoding/xml"
)

type Email struct {
	Where string `xml:"where,attr"`
	Addr  string
}

type Address struct {
	City, State string
}

type Result struct {
	XMLName xml.Name `xml:"Person"`
	Name    string   `xml:"FullName"`
	Phone   string
	Email   []Email
	Groups  []string `xml:"Group>Value"`
	Address
}

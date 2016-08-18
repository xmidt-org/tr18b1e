package main

import (
	"encoding/xml"
	"fmt"
	"github.com/comcast/tr18b1e"
	"io/ioutil"
	"os"
)

// this is an adaptation of the example provided by golang devs here:
// https://golang.org/pkg/encoding/xml/#example_Unmarshal
func main() {
	fmt.Println("\nHello world!")

	v := tr18b1e.Result{Name: "none", Phone: "none"}

	xmlFile, err := os.Open("data.xml")
	if err != nil {
		fmt.Println("Error opening file: ", err)
		os.Exit(1)
	}
	defer xmlFile.Close()

	data, err := ioutil.ReadAll(xmlFile)

	if err := xml.Unmarshal(data, &v); err != nil {
		fmt.Println("error: ", err)
		return
	}

	fmt.Println("XMLName: ", v.XMLName)
	fmt.Println("Name: ", v.Name)
	fmt.Println("Phone: ", v.Phone)
	fmt.Println("Email: ", v.Email)
	fmt.Println("Groups: ", v.Groups)
	fmt.Println("Address: ", v.Address)

	fmt.Println()
}

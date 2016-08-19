package main

import (
	"github.com/comcast/tr18b1e"
)

func main() {
	myLibrary, _ := tr18b1e.New()
	myLibrary.Get()
	myLibrary.Set()
}

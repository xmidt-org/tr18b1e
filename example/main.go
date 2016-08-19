package main

import (
	"github.com/comcast/tr18b1e"
)

func main() {
	myLibrary, _ := tr18b1e.New()
	myLibrary.Get("random")
	myLibrary.Set("random", "stuff")
}

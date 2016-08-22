package main

import (
	"fmt"
	"github.com/comcast/tr18b1e"
)

func main() {
	myLibrary, _ := tr18b1e.New()

	myData, _ := myLibrary.Get("random")
	fmt.Println(myData)

	myLibrary.Put("hello", "stuff")
	myData, _ = myLibrary.Get("hello")
	fmt.Println(myData)

	myData, _ = myLibrary.Get("hasas;ldfkjasd;f.")
	fmt.Println(myData)
}

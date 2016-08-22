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

	myLibrary.Put("goodbye", "stuff")
	myData, _ = myLibrary.Get(".")
	fmt.Println(myData)

	myLibrary.Update("goodbye", "ffuts")
	myData, _ = myLibrary.Get(".")
	fmt.Println(myData)

	myLibrary.Delete("hasas;ldfkjasd;f.")
	myData, _ = myLibrary.Get(".")
	fmt.Println(myData)
}

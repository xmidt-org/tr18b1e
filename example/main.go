package main

import (
	"fmt"
	"github.com/comcast/tr18b1e"
)

func main() {
	myLibrary, _ := tr18b1e.New()

	myLibrary.Put("hello", "hello")
	myData, _ := myLibrary.Get("hello")
	fmt.Println(myData[0])

	myLibrary.Put("good.bye", "bye")
	myLibrary.Put("good.buy", "buy")
	myData, _ = myLibrary.Get("good.")
	fmt.Println(myData[0])

	myLibrary.Delete("good.")
	myData, _ = myLibrary.Get("hello")
	fmt.Println(myData[0])
}

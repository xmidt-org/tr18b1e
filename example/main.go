package main

import (
	"fmt"
	"github.com/comcast/tr18b1e"
)

func main() {
	myLibrary, _ := tr18b1e.New()

	myData, _ := myLibrary.Get("random")
	fmt.Println(myData)

	myLibrary.Put("hello", "hello")
	myData, _ = myLibrary.Get("hello")
	fmt.Println(myData)

	myLibrary.Put("good.bye", "bye")
	myLibrary.Put("good.buy", "buy")
	myData, _ = myLibrary.Get("good.")
	fmt.Println(myData)

	myLibrary.Update("goodbye", "ffuts")
	myData, _ = myLibrary.Get(".")
	fmt.Println(myData)

	myLibrary.Delete("adfasdfasdfasfsfas.")
	myData, _ = myLibrary.Get("good.")
	fmt.Println(myData)
}

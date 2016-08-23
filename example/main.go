package main

import (
	"fmt"
	"github.com/comcast/tr18b1e"
)

func main() {
	// testing the pure CRUD interface
	myLibrary, _ := tr18b1e.New()

	myLibrary.Put("hello", "hello")
	myData, _ := myLibrary.Get("hello")
	fmt.Println(myData[0])

	myLibrary.Put("good.bye", "bye")
	myLibrary.Put("good.buy", "buy")
	myData, _ = myLibrary.Get("good.")
	fmt.Println(myData[0])

	myLibrary.Delete("good.")
	myLibrary.Update("hello", 42)
	myData, _ = myLibrary.Get("hello")
	fmt.Println(myData[0])
	fmt.Println()

	// testing the mock interface
	fakeLib, _ := tr18b1e.NewFake()
	fakeLib.Update("hello", "world")

	myData, _ = fakeLib.Get("hello")
	fmt.Println(myData[0])

	fakeLib.Update("hello", "goodbye")
	myData, _ = fakeLib.Get("hello")
	fmt.Println(myData[0])

	fmt.Println()
}

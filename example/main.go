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

	// can remove if a number
	myLibrary.Put("1.bye", "bye")
	myLibrary.Put("1.buy", "buy")
	myData, _ = myLibrary.Get("1.")
	fmt.Println("added", myData[0])
	myLibrary.Delete("1.")
	myData, _ = myLibrary.Get("1.")
	fmt.Println("could remove", len(myData))

	// cannot remove if not a number
	myLibrary.Put("good.bye", "bye")
	myLibrary.Put("good.buy", "buy")
	myData, _ = myLibrary.Get("good.")
	fmt.Println("added", myData[0])
	myLibrary.Delete("good.")
	myData, _ = myLibrary.Get("good.")
	fmt.Println("couldn't remove", len(myData))

	// testing for deep copy
	myLibrary.Update("hello", 42)
	myData, _ = myLibrary.Get("hello")
	fmt.Println("here", myData[0])
	myData[0].Value = 64
	myData, _ = myLibrary.Get("hello")
	fmt.Println("there", myData[0])
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

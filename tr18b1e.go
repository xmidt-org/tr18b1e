package tr18b1e

import (
	// "os"
	"fmt"
	// "encoding/xml"
)

type Library interface {
	Get() error
	Set() error
}

type library struct {

}

func New() (Library, error) {
	newLibrary := &library{}



	return newLibrary, nil
}

func (l *library) Get() error {
	fmt.Println("Stubbed out `Get` function")
	return nil
}

func (l *library) Set() error {
	fmt.Println("Stubbed out `Set` function")
	return nil
}

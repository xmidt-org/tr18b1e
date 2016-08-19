// name regexp
// value interface{}
// varType type

// support get, set, put, post

package tr18b1e

import (
	"fmt"
)

type trData struct {
	name string
	value string
}

type trDatabase map[string]trData

type Library interface {
	Get(key string) error
	Set(key string, data string) error
}

type library struct {
	libraryData trDatabase
}

func New() (Library, error) {
	newLibrary := &library{}

	return newLibrary, nil
}

func (l *library) Get(key string) error {
	fmt.Println("Stubbed out `Get` function")
	fmt.Println("Soon we'll get", key)
	return nil
}

func (l *library) Set(key string, data string) error {
	fmt.Println("Stubbed out `Set` function")
	fmt.Println("Soon we'll set", key, "to", data)
	return nil
}

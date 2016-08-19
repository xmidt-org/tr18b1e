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

type Library interface {
	Get(key string) (string, error)
	Set(key string, data string) error
}

type library struct {
	libraryData map[string]trData
}

func New() (Library, error) {
	newLibrary := &library{}

	newLibrary.libraryData = make(map[string]trData)

	return newLibrary, nil
}

func (l *library) Get(key string) (string, error) {
	if _, ok := l.libraryData[key]; ok {
		return l.libraryData[key].name, nil
	}

	l.libraryData[key] = trData{
		name:  key,
		value: key,
	}

	return l.libraryData[key].value, nil
}

func (l *library) Set(key string, data string) error {
	fmt.Println("Stubbed out `Set` function")
	fmt.Println("Soon we'll set", key, "to", data)
	return nil
}

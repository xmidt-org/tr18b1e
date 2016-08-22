// name regexp
// value interface{}
// varType type

// support get, put, update, delete
// only get and delete should work with the wild card
// make get return the whole array (name, value, and value type)
// create a "fake" version of this that utilizes calls to this to mimic
// 	intended behavior

package tr18b1e

import (
	"strings"
	"errors"
)

type trData struct {
	name  string
	value string
}

type Library interface {
	Get(key string) ([]string, error)
	Put(key string, data string) error
	Update(key string, data string) error
	Delete(key string) error
}

type library struct {
	libraryData map[string]*trData
}

func New() (Library, error) {
	newLibrary := &library{}

	newLibrary.libraryData = make(map[string]*trData)

	return newLibrary, nil
}

// get all the data specified by `key`
func (l *library) Get(key string) ([]string, error) {
	trValues := make([]string, 0)

	// if we're dealing with a wildcard...
	if (strings.LastIndex(key, ".") + 1) == len(key) {
		for i := range l.libraryData {
			if len(key) < len(l.libraryData[i].name) {
				if key == l.libraryData[i].name[:len(key)] {
					trValues = append(trValues, l.libraryData[i].value)
				}
			}
		}

		return trValues, nil
	}

	// if we're dealing with a single instance and it exists
	if _, ok := l.libraryData[key]; ok {
		trValues = append(trValues, l.libraryData[key].value)
		return trValues, nil
	}

	// if the key supplied does not actually represent a field that exists
	return nil, errors.New("No instance of key in data.")
}

// Create or replace `data` at the supplied `key`
func (l *library) Put(key string, data string) error {
	l.libraryData[key] = &trData{
		name:  key,
		value: data,
	}

	return nil
}

// update the data located at `key` with the passed in `data`
func (l *library) Update(key string, data string) error {
	if _, ok := l.libraryData[key]; ok {
		l.libraryData[key].value = data
		return nil
	}

	return errors.New("No instance of key in data.")
}

// deletes the data located at key
func (l *library) Delete(key string) error {
	// if we're dealing with a wildcard...
	if (strings.LastIndex(key, ".") + 1) == len(key) {
		for i := range l.libraryData {
			if len(key) < len(l.libraryData[i].name) {
				if key == l.libraryData[i].name[:len(key)] {
					delete(l.libraryData, l.libraryData[i].name)
				}
			}
		}

		return nil
	}

	delete(l.libraryData, key)

	return nil
}

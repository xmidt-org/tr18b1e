// Name string
// Value interface{}
// ValueType type

// support get, put, update, delete
// only get and delete should work with the wild card
// make get return the whole array (Name, Value, and ValueType)
// create a "fake" version of this that utilizes calls to this to mimic
// 	intended behavior

package tr18b1e

import (
	"errors"
	"reflect"
	"strings"
)

type TRData struct {
	Name  string
	Value interface{}
	ValueType reflect.Type
}

type Library interface {
	Get(key string) ([]string, error)
	Put(key string, data string) error
	Update(key string, data string) error
	Delete(key string) error
}

type library struct {
	libraryData map[string]*TRData
}

func New() (Library, error) {
	newLibrary := &library{}

	newLibrary.libraryData = make(map[string]*TRData)

	return newLibrary, nil
}

// get all the data specified by `key`
func (l *library) Get(key string) ([]string, error) {
	trValues := make([]string, 0)

	// if we're dealing with a wildcard...
	if (strings.LastIndex(key, ".") + 1) == len(key) {
		for i := range l.libraryData {
			if len(key) < len(l.libraryData[i].Name) {
				if key == l.libraryData[i].Name[:len(key)] {
					trValues = append(trValues, l.libraryData[i].Value.(string))
				}
			}
		}

		return trValues, nil
	}

	// if we're dealing with a single instance and it exists
	if _, ok := l.libraryData[key]; ok {
		trValues = append(trValues, l.libraryData[key].Value.(string))
		return trValues, nil
	}

	// if the key supplied does not actually represent a field that exists
	return nil, errors.New("No instance of key in data.")
}

// Create or replace `data` at the supplied `key`
func (l *library) Put(key string, data string) error {
	l.libraryData[key] = &TRData{
		Name:  key,
		Value: data,
	}

	return nil
}

// update the data located at `key` with the passed in `data`
func (l *library) Update(key string, data string) error {
	if _, ok := l.libraryData[key]; ok {
		l.libraryData[key].Value = data
		return nil
	}

	return errors.New("No instance of key in data.")
}

// deletes the data located at key
func (l *library) Delete(key string) error {
	// if we're dealing with a wildcard...
	if (strings.LastIndex(key, ".") + 1) == len(key) {
		for i := range l.libraryData {
			if len(key) < len(l.libraryData[i].Name) {
				if key == l.libraryData[i].Name[:len(key)] {
					delete(l.libraryData, l.libraryData[i].Name)
				}
			}
		}

		return nil
	}

	delete(l.libraryData, key)

	return nil
}

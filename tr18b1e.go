// name regexp
// value interface{}
// varType type

// support get, put, update, delete

package tr18b1e

import (
	"strings"
)

type trData struct {
	name string
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

func (l *library) Get(key string) ([]string, error) {
	trValues := make([]string, 0)

	if (strings.LastIndex(key, ".") + 1) == len(key) {
		for i := range l.libraryData {
			trValues = append(trValues, l.libraryData[i].value)
		}

		return trValues, nil
	}

	if _, ok := l.libraryData[key]; ok {
		trValues = append(trValues, l.libraryData[key].value)
		return trValues, nil
	}

	l.libraryData[key] = &trData{
		name:  key,
		value: key,
	}

	trValues = append(trValues, l.libraryData[key].value)

	return trValues, nil
}

func (l *library) Put(key string, data string) error {
	l.libraryData[key] = &trData{
		name:  key,
		value: data,
	}

	return nil
}

func (l *library) Update(key string, data string) error {
	if _, ok := l.libraryData[key]; ok {
		l.libraryData[key].value = data
		return nil
	}

	l.libraryData[key] = &trData{
		name:  key,
		value: data,
	}

	return nil
}

func (l *library) Delete(key string) error {
	delete(l.libraryData, key)

	return nil
}

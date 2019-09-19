/**
 * Copyright 2019 Comcast Cable Communications Management, LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */
// Name string
// Value interface{}
// ValueType type

// support get, put, update, delete
// only get and delete should work with the wild card
// make get return the whole array (Name, Value, and ValueType)
// create a "fake" version of this that utilizes calls to this to mimic
// 	intended behavior

// TODO: populate currently appends the data... is that what we want?
// 			 I made populate part of the interface... is that ok?

package tr18b1e

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

// TRData is a struct representing the that we want to provide to the user
// we store the value type since the value itself is an interface{}
type TRData struct {
	Name      string
	Value     interface{}
	ValueType reflect.Type
}

// Library is the CRUD interface that we provide to the user, implemented below
// by the library struct, which uses these functions to adjust its internal map
type Library interface {
	Get(key string) ([]TRData, error)
	Put(key string, data interface{}) error
	Update(key string, data interface{}) error
	Delete(key string) error
	Populate(inData []TRData) error
}

type library struct {
	libraryData map[string]*TRData
}

// New is used to generate a new instance of a library (and hence,
// a new map) to be used to make calls to the CRUD interface
func New() (Library, error) {
	newLibrary := &library{}

	newLibrary.libraryData = make(map[string]*TRData)

	return newLibrary, nil
}

func (l *library) Populate(inData []TRData) error {
	for i := range inData {
		l.libraryData[inData[i].Name] = &TRData{
			Name:      inData[i].Name,
			Value:     inData[i].Value,
			ValueType: inData[i].ValueType,
		}
	}

	return nil
}

// get all the data specified by `key`
func (l *library) Get(key string) ([]TRData, error) {
	trValues := make([]TRData, 0)

	// if we're dealing with a wildcard...
	if (strings.LastIndex(key, ".") + 1) == len(key) {
		for i := range l.libraryData {
			if len(key) < len(l.libraryData[i].Name) {
				if key == l.libraryData[i].Name[:len(key)] {
					trValues = append(trValues, *l.libraryData[i])
				}
			}
		}

		if len(trValues) == 0 {
			return nil, errors.New("No instance of key in data.")
		}

		return trValues, nil
	}

	// if we're dealing with a single instance and it exists
	if _, ok := l.libraryData[key]; ok {
		trValues = append(trValues, *l.libraryData[key])
		return trValues, nil
	}

	// if the key supplied does not actually represent a field that exists
	return nil, errors.New("No instance of key in data.")
}

// Create or replace `data` at the supplied `key`
func (l *library) Put(key string, data interface{}) error {
	// going to need to check if the provided key exists within TR-181
	l.libraryData[key] = &TRData{
		Name:      key,
		Value:     data,
		ValueType: reflect.TypeOf(data),
	}

	return nil
}

// update the data located at `key` with the passed in `data`
func (l *library) Update(key string, data interface{}) error {
	if _, ok := l.libraryData[key]; ok {
		l.libraryData[key].Value = data
		l.libraryData[key].ValueType = reflect.TypeOf(data)
		return nil
	}

	return errors.New("No instance of key in data.")
}

// deletes the data located at key
func (l *library) Delete(key string) error {
	// if we're dealing with a wildcard...
	if (strings.LastIndex(key, ".")+1) == len(key) && len(key) > 1 {
		keySplit := strings.Split(key, ".")

		if _, err := strconv.Atoi(keySplit[len(keySplit)-2]); err == nil {
			for i := range l.libraryData {
				if len(key) < len(l.libraryData[i].Name) {
					if key == l.libraryData[i].Name[:len(key)] {
						delete(l.libraryData, l.libraryData[i].Name)
					}
				}
			}

			return nil
		} else {
			return err
		}
	}

	delete(l.libraryData, key)

	return nil
}

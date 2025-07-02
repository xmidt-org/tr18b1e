// SPDX-FileCopyrightText: 2025 Comcast Cable Communications Management, LLC
// SPDX-License-Identifier: Apache-2.0
package tr18b1e

import (
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestNew(t *testing.T) {
	assert := assert.New(t)

	_, err := New()

	assert.Nil(err)
}

func TestPopulate(t *testing.T) {
	assert := assert.New(t)

	newLib, _ := New()

	inData := []TRData{
		{
			Name:      "hello.world",
			Value:     "world",
			ValueType: reflect.TypeOf("world"),
		},
		{
			Name:      "hello.42",
			Value:     42,
			ValueType: reflect.TypeOf(42),
		},
		{
			Name:      "hello.stuff",
			Value:     "stuff",
			ValueType: reflect.TypeOf("stuff"),
		},
	}

	err := newLib.Populate(inData)
	assert.Nil(err)

	outData, err := newLib.Get("hello.")
	assert.Nil(err)
	assert.Equal(3, len(outData))
}

func TestPut(t *testing.T) {
	assert := assert.New(t)
	newLib, _ := New()

	err := newLib.Put("hello", "world")
	myData, _ := newLib.Get("hello")

	assert.Nil(err)
	assert.Equal("world", myData[0].Value.(string))
}

func TestUpdate(t *testing.T) {
	assert := assert.New(t)
	newLib, _ := New()

	newLib.Put("hello", "world")
	err := newLib.Update("hello", 42)
	myData, _ := newLib.Get("hello")

	assert.Nil(err)
	assert.Equal(42, myData[0].Value.(int))
}

func TestGet(t *testing.T) {
	assert := assert.New(t)
	newLib, _ := New()

	newLib.Put("hello", "world")

	newLib.Put("hello.world", "testing")
	newLib.Put("hello.planet", "functions")

	// test getting just one entry
	myData, err := newLib.Get("hello")
	assert.Nil(err)
	assert.Equal("world", myData[0].Value.(string))

	// test getting entries with the wild card
	myData, err = newLib.Get("hello.")
	assert.Nil(err)
	assert.Equal(len(myData), 2)
}

func TestDelete(t *testing.T) {
	assert := assert.New(t)
	newLib, _ := New()

	newLib.Put("hello", "world")

	newLib.Put("hello.world", "testing")
	newLib.Put("hello.planet", "functions")

	newLib.Put("101.hello", "numbers")
	newLib.Put("101.world", "work")

	// test deleting one entry with the wild card
	err := newLib.Delete("hello")
	myData, _ := newLib.Get("hello")
	assert.Nil(err)
	assert.Equal(0, len(myData))

	// test deleting multiple entries with the wild card
	// this shouldn't work because we're not giving it a number
	err = newLib.Delete("hello.")
	myData, _ = newLib.Get("hello.")
	assert.NotNil(err)
	assert.Equal(2, len(myData))

	// test deleting multiple entries with the wild card
	// this should work because we are giving it a number
	err = newLib.Delete("101.")
	myData, _ = newLib.Get("101.")
	assert.Nil(err)
	assert.Equal(0, len(myData))
}

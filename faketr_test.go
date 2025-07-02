// SPDX-FileCopyrightText: 2025 Comcast Cable Communications Management, LLC
// SPDX-License-Identifier: Apache-2.0
package tr18b1e

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFake(t *testing.T) {
	assert := assert.New(t)

	_, err := NewFake()

	assert.Nil(err)
}

func TestPopulateFake(t *testing.T) {
	assert := assert.New(t)

	testFakeLib, _ := NewFake()

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

	err := testFakeLib.Populate(inData)
	assert.Nil(err)

	// test if we get what we expect when we call get on data we provided
	outData, err := testFakeLib.Get("hello.")
	assert.Nil(err)
	assert.Equal(3, len(outData))

	// test if we get an error when asking for a wildcard that does not exist
	outData, err = testFakeLib.Get("broken.")
	assert.NotNil(err)
	assert.Empty(outData)
}

func TestGetFake(t *testing.T) {
	assert := assert.New(t)

	testFakeLib, _ := NewFake()

	// first test determines if it will put the missing data into the library
	testData, err := testFakeLib.Get("hello")
	assert.Equal("hello", testData[0].Value.(string))
	assert.Nil(err)

	// second test determines if it remains the same between function calls
	testData, err = testFakeLib.Get("hello")
	assert.Equal("hello", testData[0].Value.(string))
	assert.Nil(err)
}

func TestPutFake(t *testing.T) {
	assert := assert.New(t)

	testFakeLib, _ := NewFake()

	err := testFakeLib.Put("hello", "world")
	assert.Nil(err)

	testData, _ := testFakeLib.Get("hello")
	assert.Equal("world", testData[0].Value.(string))
}

func TestUpdateFake(t *testing.T) {
	assert := assert.New(t)

	testFakeLib, _ := NewFake()

	// test what happens when the data is not in there at all
	err := testFakeLib.Update("hello", "world")
	assert.Nil(err)

	testData, _ := testFakeLib.Get("hello")
	assert.Equal("world", testData[0].Value.(string))

	// test what happens when the data is already in there
	err = testFakeLib.Update("hello", "goodbye")
	assert.Nil(err)

	testData, _ = testFakeLib.Get("hello")
	assert.Equal("goodbye", testData[0].Value.(string))
}

func TestDeleteFake(t *testing.T) {
	assert := assert.New(t)

	testFakeLib, _ := NewFake()

	err := testFakeLib.Put("hello", "world")
	assert.Nil(err)
	err = testFakeLib.Delete("hello")
	assert.Nil(err)
}

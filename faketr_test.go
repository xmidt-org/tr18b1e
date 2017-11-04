/**
 * Copyright 2016 Comcast Cable Communications Management, LLC
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
package tr18b1e

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
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

	// test if we get what we expect when we call get on data we provided
	outData, err := testFakeLib.Get("hello.")
	assert.Nil(err)
	assert.Equal(3, len(outData))

	// test if we get an error when asking for a wildcard that does not exist
	outData, err = testFakeLib.Get("broken.")
	assert.NotNil(err)
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
	err = testFakeLib.Delete("hello")
	assert.Nil(err)
}

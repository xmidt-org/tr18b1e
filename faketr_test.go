package tr18b1e

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMock(t *testing.T) {
	assert := assert.New(t)

	_, err := NewMock()

	assert.Nil(err)
}

func TestGetMock(t *testing.T) {
	assert := assert.New(t)

	testMockLib, _ := NewMock()

	// first test determines if it will put the missing data into the library
	testData, err := testMockLib.GetMock("hello")
	assert.Equal("hello", testData[0].Value.(string))
	assert.Nil(err)

	// second test determines if it remains the same between function calls
	testData, err = testMockLib.GetMock("hello")
	assert.Equal("hello", testData[0].Value.(string))
	assert.Nil(err)
}

func TestPutMock(t *testing.T) {
  assert := assert.New(t)

  testMockLib, _ := NewMock()

  err := testMockLib.PutMock("hello", "world")
  assert.Nil(err)

  testData, _ := testMockLib.GetMock("hello")
  assert.Equal("world", testData[0].Value.(string))
}

func TestUpdateMock(t *testing.T) {
  assert := assert.New(t)

  testMockLib, _ := NewMock()

  // test what happens when the data is not in there at all
  err := testMockLib.UpdateMock("hello", "world")
  assert.Nil(err)

  testData, _ := testMockLib.GetMock("hello")
  assert.Equal("world", testData[0].Value.(string))

  // test what happens when the data is already in there
  err = testMockLib.UpdateMock("hello", "goodbye")
  assert.Nil(err)

  testData, _ = testMockLib.GetMock("hello")
  assert.Equal("goodbye", testData[0].Value.(string))
}

func TestDeleteMock(t *testing.T) {
  assert := assert.New(t)

  testMockLib, _ := NewMock()

  err := testMockLib.PutMock("hello", "world")
  err = testMockLib.DeleteMock("hello")
  assert.Nil(err)
}

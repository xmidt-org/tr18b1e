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
	assert.Equal(testData[0].Value.(string), "hello")
	assert.Nil(err)

	// second test determines if it remains the same between function calls
	testData, err = testMockLib.GetMock("hello")
	assert.Equal(testData[0].Value.(string), "hello")
	assert.Nil(err)
}

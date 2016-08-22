package tr18b1e

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMain(m *testing.M) {

	os.Exit(m.Run())
}

func TestNew(t *testing.T) {
	assert := assert.New(t)

	_, err := New()

	assert.Nil(err)
}

func TestGet(t *testing.T) {
	assert := assert.New(t)
	newLib, _ := New()

	// test to see if the function works when making up the value
	_, err := newLib.Get("Device")
	assert.Nil(err)

	// test to see if it works when the value already exists
	_, err = newLib.Get("Device")
	assert.Nil(err)

	// test to see if it works when we have two values
	_, err = newLib.Get("ID")
	_, err = newLib.Get(".")
	assert.Nil(err)
}

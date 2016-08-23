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

	// test deleting one entry with the wild card
	err := newLib.Delete("hello")
	myData, _ := newLib.Get("hello")
	assert.Nil(err)
	assert.Equal(0, len(myData))

	// test deleting multiple entries with the wild card
	err = newLib.Delete("hello.")
	myData, _ = newLib.Get("hello.")
	assert.Nil(err)
	assert.Equal(0, len(myData))
}

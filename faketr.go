package tr18b1e

import (
	"errors"
	"strings"
)

type fakeLibrary struct {
	lib *library
}

// NewFake creates a fakeLibrary that implements the same interface but
// has behavior that delivers fake information to satisfy tests
func NewFake() (Library, error) {
	fakeLib := &fakeLibrary{}
	fakeLib.lib = &library{}

	fakeLib.lib.libraryData = make(map[string]*TRData)

	return fakeLib, nil
}

// Populate just makes a call to the Populate implemented by library
func (m *fakeLibrary) Populate(inData []TRData) error {
	return m.lib.Populate(inData)
}

// Get will return an error if given a "wildcard" it has no data for,
// it will return the value if it exists already,
// and will create a single value and return a single value if it does not yet exist
func (m *fakeLibrary) Get(key string) ([]TRData, error) {
	myData, err := m.lib.Get(key)

	if err != nil || len(myData) == 0 {
		if (strings.LastIndex(key, ".") + 1) == len(key) {
			return nil, errors.New("Using a wildcard in the fake tr18b1e w/no data supplied.")
		}

		m.lib.Put(key, key)
		return m.lib.Get(key)
	}

	return myData, nil
}

// Put just does the same thing that the library put function does
func (m *fakeLibrary) Put(key string, data interface{}) error {
	return m.lib.Put(key, data)
}

// Update will update the entry if it exists, and will make a new one if
// it doesn't yet exist
func (m *fakeLibrary) Update(key string, data interface{}) error {
	if err := m.lib.Update(key, data); err != nil {
		return m.lib.Put(key, data)
	}

	return m.lib.Update(key, data)
}

// Delete will delete the entry if it exists and do nothing if it does not
func (m *fakeLibrary) Delete(key string) error {
	return m.lib.Delete(key)
}

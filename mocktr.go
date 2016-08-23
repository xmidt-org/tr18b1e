package tr18b1e

type MockLibrary interface {
	GetMock(key string) ([]*TRData, error)
	PutMock(key string, data interface{}) error
	UpdateMock(key string, data interface{}) error
	DeleteMock(key string) error
}

type mockLibrary struct {
	lib *library
}

func NewMock() (MockLibrary, error) {
	mockLib := &mockLibrary{}
	mockLib.lib = &library{}

	mockLib.lib.libraryData = make(map[string]*TRData)

	return mockLib, nil
}

func (m *mockLibrary) GetMock(key string) ([]*TRData, error) {
	return nil, nil
}

func (m *mockLibrary) PutMock(key string, data interface{}) error {
	return nil
}

func (m *mockLibrary) UpdateMock(key string, data interface{}) error {
	return nil
}

func (m *mockLibrary) DeleteMock(key string) error {
	return nil
}

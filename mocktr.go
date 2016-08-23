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
  trValues := make([]*TRData, 0)

  if _, ok := m.lib.libraryData[key]; ok {
    trValues = append(trValues, m.lib.libraryData[key])
    return trValues, nil
  }

  m.lib.Put(key, key)

  return m.lib.Get(key)
}

func (m *mockLibrary) PutMock(key string, data interface{}) error {
  m.lib.Put(key, data)

	return nil
}

func (m *mockLibrary) UpdateMock(key string, data interface{}) error {
  if err := m.lib.Update(key, data); err != nil {
    m.lib.Put(key, data)
    return nil
  }

	return nil
}

func (m *mockLibrary) DeleteMock(key string) error {
  m.lib.Delete(key)
  
	return nil
}

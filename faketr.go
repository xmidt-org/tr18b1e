package tr18b1e

type fakeLibrary struct {
	lib *library
}

func NewFake() (Library, error) {
	fakeLib := &fakeLibrary{}
	fakeLib.lib = &library{}

	fakeLib.lib.libraryData = make(map[string]*TRData)

	return fakeLib, nil
}

func (m *fakeLibrary) Get(key string) ([]TRData, error) {
	trValues := make([]TRData, 0)

	if _, ok := m.lib.libraryData[key]; ok {
		trValues = append(trValues, *m.lib.libraryData[key])
		return trValues, nil
	}

	m.lib.Put(key, key)

	return m.lib.Get(key)
}

func (m *fakeLibrary) Put(key string, data interface{}) error {
	return m.lib.Put(key, data)
}

func (m *fakeLibrary) Update(key string, data interface{}) error {
	if err := m.lib.Update(key, data); err != nil {
		return m.lib.Put(key, data)
	}

	return m.lib.Update(key, data)
}

func (m *fakeLibrary) Delete(key string) error {
	return m.lib.Delete(key)
}

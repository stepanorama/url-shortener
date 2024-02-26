package storage

// Concrete type that's gonna be used by methods to store/retrieve URLs.
type MapStorage struct {
	urlMap map[string]string
}

// Func that returns an instance of the MapStorage type. Accepting interface here instead of MapStorage (???)
func NewMapStorage() *MapStorage {
	return &MapStorage{urlMap: make(map[string]string)}
}

// Method for storing URL
func (m *MapStorage) StoreURL(shortURL, fullURL string) error {
	m.urlMap[shortURL] = fullURL
	return nil // TODO Handle possible errors later
}

// Method for getting URL
func (m *MapStorage) RetrieveURL(shortURL string) (string, bool) {
	fullURL, ok := m.urlMap[shortURL]
	return fullURL, ok
}

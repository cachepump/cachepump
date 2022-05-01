package cache

import (
	"errors"
	"sync"
)

var data = make(map[string][]byte)
var dataMtx sync.RWMutex

// ErrKeyNotFound will return if selected key has been not found into inmemory storage.
var ErrKeyNotFound = errors.New("key is not found")

// Get returns value by selected key from inmemory storage.
func Get(key string) (value []byte, err error) {
	dataMtx.RLock()
	defer dataMtx.RUnlock()

	var ok bool
	value, ok = data[key]
	if !ok {
		return []byte{}, ErrKeyNotFound
	}
	return value, err
}

// Set override value by selected key into inmemory storage.
// If selected key is not definition in storage he will created.
func Set(key string, value []byte) (err error) {
	dataMtx.Lock()
	defer dataMtx.Unlock()
	data[key] = value
	return nil
}

// Del removes selected key from inmemory storage.
func Del(key string) {
	dataMtx.Lock()
	defer dataMtx.Unlock()
	delete(data, key)
}

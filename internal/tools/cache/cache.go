package cache

import (
	"sync"
)

type Cache struct {
	mu   sync.RWMutex
	data map[string]string
}

func New() *Cache {
	return &Cache{mu: sync.RWMutex{}, data: map[string]string{}}
}

func (ref *Cache) Get(key string) (string, bool) {
	ref.mu.RLock()
	defer ref.mu.RUnlock()

	value, ok := ref.data[key]

	return value, ok
}

func (ref *Cache) Save(key, value string) {
	ref.mu.Lock()
	defer ref.mu.Unlock()

	ref.data[key] = value
}

package cache

import "sync"

type cache struct {
	mu   sync.RWMutex
	data map[string]string
}

func New() *cache {
	return &cache{data: map[string]string{}}
}

func (ref *cache) Get(key string) (string, bool) {

	ref.mu.RLock()
	defer ref.mu.RUnlock()

	value, ok := ref.data[key]

	return value, ok
}

func (ref *cache) Save(key, value string) {
	ref.mu.Lock()
	defer ref.mu.Unlock()

	ref.data[key] = value
}

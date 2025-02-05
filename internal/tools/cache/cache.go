package cache

import (
	"sync"
	"time"
)

type Cache[T any] struct {
	mu   sync.RWMutex
	data map[string]element[T]
	ttl  *time.Duration
}

type element[T any] struct {
	value      T
	insertTime time.Time
}

func New[T any](ttl *time.Duration) *Cache[T] {
	return &Cache[T]{mu: sync.RWMutex{}, data: map[string]element[T]{}, ttl: ttl}
}

func (ref *Cache[T]) Get(key string) (*T, bool) {
	ref.mu.RLock()
	defer ref.mu.RUnlock()

	value, found := ref.data[key]

	if ref.ttl != nil && value.insertTime.Add(*ref.ttl).Before(time.Now()) {
		delete(ref.data, key)

		return nil, false
	}

	if found {
		return &value.value, true
	}

	return nil, false
}

func (ref *Cache[T]) Save(key string, value T) {
	ref.mu.Lock()
	defer ref.mu.Unlock()

	ref.data[key] = element[T]{value: value, insertTime: time.Now()}
}

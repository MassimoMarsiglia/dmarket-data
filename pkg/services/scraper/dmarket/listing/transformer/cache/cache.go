package cache

import (
	"sync"
	"time"
)

type Item[T any] struct {
	value  T
	expiry time.Time
}

type Cache[T any] struct {
	mu    sync.Mutex
	items map[string]Item[T]
}

func NewCache[T any]() Cache[T] {
	return Cache[T]{
		items: make(map[string]Item[T]),
	}
}

func (c *Cache[T]) Get(key string) (T, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	item, ok := c.items[key]
	if !ok || time.Now().After(item.expiry) {
		delete(c.items, key)
		var zero T
		return zero, false
	}
	return item.value, true
}

func (c *Cache[T]) Set(key string, value T, duration time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items[key] = Item[T]{
		value:  value,
		expiry: time.Now().Add(duration),
	}
}

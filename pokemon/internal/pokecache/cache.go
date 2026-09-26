package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu   sync.Mutex
	data map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]cacheEntry),
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.data[key]

	if !ok {
		return nil, false
	}
	return entry.val, true
}

func reapLoop() {

}

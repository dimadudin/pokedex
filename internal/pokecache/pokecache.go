package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	val       []byte
	createdAt time.Time
}

type Cache struct {
	data map[string]cacheEntry
	mu   *sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	new := Cache{
		data: make(map[string]cacheEntry),
		mu:   &sync.Mutex{},
	}
	go new.reapLoop(interval)
	return new
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = cacheEntry{val: val, createdAt: time.Now()}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.data[key]
	return entry.val, ok
}

func (c *Cache) reap(t time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	timeAgo := time.Now().Add(-t)
	for k, v := range c.data {
		if v.createdAt.Before(timeAgo) {
			delete(c.data, k)
		}
	}
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for range ticker.C {
		c.reap(interval)
	}
}

package cache

import (
	"time"
	"sync"
)

type Cache struct {
	cachedData map[string]cacheEntry
	mu sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

func NewCache(interval time.Duration) *Cache {
	return &Cache {
		cachedData: make(map[string]cacheEntry),
		mu: sync.Mutex{},
	}
}

func (c *Cache) Add(key string, value []byte) {
	
}
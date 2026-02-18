package cache

import (
	"time"
	"sync"
)
// ======= structs =======
type Cache struct {
	cachedData map[string]cacheEntry
	mu sync.RWMutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	value []byte
}

//======= constructor function ======
func NewCache(interval time.Duration) *Cache {
	c := &Cache {
		cachedData: make(map[string]cacheEntry),
		mu: sync.RWMutex{},
		interval: interval,
	}
	go c.reapLoop()
	return c
}


//======== methods ==========
func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cachedData[key] = cacheEntry{
		createdAt: time.Now(),
		value: value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	cachedData, ok := c.cachedData[key]
	if !ok {
		return nil, false
	} 
	return cachedData.value, true
}
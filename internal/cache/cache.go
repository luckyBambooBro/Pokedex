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
	}
	go c.reapLoop(interval)
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

// ======== Methods to delete caches ========
func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for range ticker.C {
		c.reap()
	}
}

func (c *Cache) reap() {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, v := range c.cachedData {
		elapsedTime := time.Since(v.createdAt) //lessons solution passed in time.Time from reaploop() and subtracted interval from it, then compared to v.createdAt
		if elapsedTime >= c.interval {
			delete(c.cachedData,k)
		}
	}
}
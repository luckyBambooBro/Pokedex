package cache

import "time"

// ======== Methods to delete caches ========
func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	for range ticker.C {
		c.reap()
	}
}

func (c *Cache) reap() {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, v := range c.cachedData {
		elapsedTime := time.Since(v.createdAt)
		if elapsedTime >= c.interval {
			delete(c.cachedData,k)
		}
	}
}
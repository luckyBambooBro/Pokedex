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


package pokecache

import (
	"sync"
	"time"
)

// ---------------------------------------------------------
// CACHE INITIALIZATION
// ---------------------------------------------------------

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	msc      map[string]cacheEntry
	mu       sync.Mutex
	interval time.Duration
}

func NewCache(interval0 time.Duration) *Cache {
	c := &Cache{
		msc:      make(map[string]cacheEntry),
		interval: interval0,
	}

	go c.reapLoop()

	return c
}

// ---------------------------------------------------------
// FUNCTIONS
// ---------------------------------------------------------

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.msc[key] = cacheEntry{time.Now(), val}
	// this is a cacheentry so we need the time right now AND the val as the val
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	value, ok := c.msc[key]
	return value.val, ok
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		c.mu.Lock()
		for key, value := range c.msc {
			elapsed := time.Since(value.createdAt)
			if elapsed > c.interval {
				delete(c.msc, key)
			}
		}
		c.mu.Unlock()
	}
}

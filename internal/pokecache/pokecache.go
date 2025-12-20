package Pokecache

import (
	"fmt"
	"sync"
	"time"
)

const cacheInterval = 15 * time.Second

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cache map[string]cacheEntry
	mu    sync.Mutex
}

func NewCache() *Cache {
	cache := Cache{
		cache: map[string]cacheEntry{},
	}
	go cache.reapLoop(cacheInterval)
	return &cache
}

func (c *Cache) Add(key string, val []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, ok := c.cache[key]
	if ok {
		return fmt.Errorf("%v already exists", c.cache[key])
	}
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	return nil
}

func (c *Cache) Get(key string) (val []byte, entryExists bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, ok := c.cache[key]
	if !ok {
		return []byte{}, false
	}
	return c.cache[key].val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		for k, v := range c.cache {
			if time.Since(v.createdAt) > interval {
				c.mu.Lock()
				delete(c.cache, k)
				c.mu.Unlock()
			}
		}
	}
}

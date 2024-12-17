package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	pokecache  map[string]cacheEntry
	cacheMutex sync.Mutex
}

var (
	cacheInstance *Cache
	mu            sync.Mutex
)

func InitCache(interval time.Duration) {
	cacheInstance = NewCache(interval)
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		pokecache:  make(map[string]cacheEntry),
		cacheMutex: sync.Mutex{},
	}

	go c.reapLoop(interval)

	return c
}

func (c *Cache) Add(key string, val []byte) error {
	if key == "" {
		return fmt.Errorf("key cannot be empty")
	}
	c.cacheMutex.Lock()
	c.pokecache[key] = cacheEntry{createdAt: time.Now(), val: val}
	defer c.cacheMutex.Unlock()

	return nil
}

func (c *Cache) Get(key string) ([]byte, bool) {
	if key == "" {
		return []byte{}, false
	}
	c.cacheMutex.Lock()
	val, ok := c.pokecache[key]
	defer c.cacheMutex.Unlock()
	if !ok {
		return []byte{}, false
	}
	return val.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		c.cacheMutex.Lock()
		for key, value := range c.pokecache {
			if time.Now().Sub(value.createdAt) > interval {
				delete(c.pokecache, key)
			}
		}
		c.cacheMutex.Unlock()
	}
}

func AddToCache(key string, val []byte) {
	mu.Lock()
	defer mu.Unlock()
	cacheInstance.Add(key, val)
}

func GetFromCache(key string) ([]byte, bool) {
	mu.Lock()
	defer mu.Unlock()
	return cacheInstance.Get(key)
}

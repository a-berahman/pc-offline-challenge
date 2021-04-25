package cache

import "sync"

// Cache represents basic property for Cache repository
type Cache struct {
	lock *sync.RWMutex
	DB   map[string]string
}

// New returns new instance of cache repository
func New() *Cache {
	return &Cache{lock: &sync.RWMutex{}, DB: make(map[string]string)}
}

//Get returns value of cache with key argument
func (c *Cache) Get(key string) string {
	c.lock.RLock()
	defer c.lock.RUnlock()
	if _, exist := c.DB[key]; exist {
		return c.DB[key]
	}
	return ""
}

//Write inerts new key value in cache db
func (c *Cache) Write(key, value string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.DB[key] = value
}

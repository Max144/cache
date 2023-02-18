package cache

import (
	"errors"
	"time"
)

type Cache struct {
	cache            map[string]interface{}
	keyDeleteTimeMap map[string]time.Time
}

func New() Cache {
	cache := Cache{make(map[string]interface{}), make(map[string]time.Time)}
	go cache.deletingExpiredCacheWorker()
	return cache
}

func (c *Cache) Set(key string, value interface{}, duration time.Duration) {
	c.cache[key] = value
	c.keyDeleteTimeMap[key] = time.Now().Add(duration)
}

func (c *Cache) Get(key string) (interface{}, error) {
	value, ok := c.cache[key]
	if !ok {
		return nil, errors.New("key doesn't exists")
	}
	return value, nil
}

func (c *Cache) deletingExpiredCacheWorker() {
	for {
		for key, expiryTime := range c.keyDeleteTimeMap {
			if expiryTime.Before(time.Now()) {
				delete(c.cache, key)
				delete(c.keyDeleteTimeMap, key)
			}
		}
	}
}

func (c *Cache) Delete(key string) error {
	_, ok := c.cache[key]
	if !ok {
		return errors.New("key doesn't exists")
	}

	delete(c.cache, key)
	return nil
}

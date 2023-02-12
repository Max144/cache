package cache

import "errors"

type Cache struct {
	cache map[string]interface{}
}

func New() Cache {
	return Cache{make(map[string]interface{})}
}

func (c *Cache) Set(key string, value interface{}) {
	c.cache[key] = value
}

func (c *Cache) Get(key string) (interface{}, error) {
	value, ok := c.cache[key]
	if !ok {
		return nil, errors.New("key doesn't exists")
	}
	return value, nil
}

func (c *Cache) Delete(key string) error {
	_, ok := c.cache[key]
	if !ok {
		return errors.New("key doesn't exists")
	}

	delete(c.cache, key)
	return nil
}

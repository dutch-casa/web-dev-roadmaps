// Option B: The deep approach
//
// Three methods. Expiration is automatic. The caller never sees
// entries, timestamps, or cleanup logic.

package option_b

import "time"

type entry struct {
	value     string
	expiresAt time.Time
}

type Cache struct {
	entries map[string]entry
}

func New(cleanupInterval time.Duration) *Cache {
	c := &Cache{entries: make(map[string]entry)}
	go func() {
		for range time.Tick(cleanupInterval) {
			c.evictExpired()
		}
	}()
	return c
}

func (c *Cache) Set(key, value string, ttl time.Duration) {
	c.entries[key] = entry{value: value, expiresAt: time.Now().Add(ttl)}
}

func (c *Cache) Get(key string) (string, bool) {
	e, ok := c.entries[key]
	if !ok {
		return "", false
	}
	if time.Now().After(e.expiresAt) {
		delete(c.entries, key)
		return "", false
	}
	return e.value, true
}

func (c *Cache) Delete(key string) {
	delete(c.entries, key)
}

func (c *Cache) evictExpired() {
	for key, e := range c.entries {
		if time.Now().After(e.expiresAt) {
			delete(c.entries, key)
		}
	}
}

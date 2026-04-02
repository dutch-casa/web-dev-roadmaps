// Option A: The shallow approach
//
// Every internal detail is exposed. The caller manages expiration,
// creates entries, checks timestamps, and drives the cleanup.

package option_a

import "time"

type CacheEntry struct {
	Key       string
	Value     string
	CreatedAt time.Time
	ExpiresAt time.Time
	IsExpired bool
}

type Cache struct {
	Entries map[string]CacheEntry
}

func NewCache() *Cache {
	return &Cache{Entries: make(map[string]CacheEntry)}
}

func CreateEntry(key, value string, ttl time.Duration) CacheEntry {
	now := time.Now()
	return CacheEntry{
		Key:       key,
		Value:     value,
		CreatedAt: now,
		ExpiresAt: now.Add(ttl),
		IsExpired: false,
	}
}

func (c *Cache) PutEntry(entry CacheEntry) {
	c.Entries[entry.Key] = entry
}

func (c *Cache) GetEntry(key string) (CacheEntry, bool) {
	entry, ok := c.Entries[key]
	return entry, ok
}

func (c *Cache) DeleteEntry(key string) {
	delete(c.Entries, key)
}

func CheckExpired(entry CacheEntry) bool {
	return time.Now().After(entry.ExpiresAt)
}

func (c *Cache) CleanupExpired() int {
	removed := 0
	for key, entry := range c.Entries {
		if CheckExpired(entry) {
			delete(c.Entries, key)
			removed++
		}
	}
	return removed
}

func (c *Cache) Size() int {
	return len(c.Entries)
}

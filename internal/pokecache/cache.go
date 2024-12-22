package pokecache

import (
	"time"
	"sync"

)

//Cache

type Cache struct {
	Timeout time.Duration
	data map[string]cacheEntry
	lock sync.RWMutex

}

type cacheEntry struct {
	createdAt time.Time 
	val []byte
}


func NewCache(timeinterval time.Duration) Cache {
	  cache := Cache{
			Timeout: timeinterval,
			data: make(map[string]cacheEntry),
		}

		cache.reapLoop()


		return cache 
}


func (c *Cache) Add(key string, v []byte) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.data[key] = cacheEntry{
		           createdAt: time.Now(),
							 val: v,
	           }
} 


func (c *Cache) Get(key string) ([]byte, bool) {
	 c.lock.RLock()

	 defer c.lock.RUnlock()

	 entry, exists := c.data[key]

	 if !exists {
		 return nil, false
	 }

	 return entry.val, true

}


func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.Timeout)
	go func() {
		for range ticker.C {
			c.lock.Lock()
			

			now := time.Now()

			for key, cachedata := range c.data {
				timeDiff := now.Sub(cachedata.createdAt)
				if timeDiff > c.Timeout {

					delete(c.data, key)

				}
     }

			c.lock.Unlock()
		}
	}()
}

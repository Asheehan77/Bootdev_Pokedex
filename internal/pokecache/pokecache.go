package pokecache

import(
	"time"
	"sync"
)

type Cache struct {
	urlcache	map[string]cacheEntry
	interval	time.Duration
	mutex		*sync.Mutex
}

type cacheEntry struct {
	createdAt 	time.Time
	val			[]byte
}

func NewCache(interv time.Duration)Cache{
	c := Cache{
		urlcache: make(map[string]cacheEntry),
		interval: interv,
		mutex:   &sync.Mutex{},
	}
	go c.reapLoop()
	return c
	
}

func (c *Cache)Add(url string, entry []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.urlcache[url] = cacheEntry{
		createdAt: time.Now(),
		val: entry,
	}
}


func (c * Cache)Get(url string)([]byte, bool){
	c.mutex.Lock()
	defer c.mutex.Unlock()
	val,exists := c.urlcache[url]
	return val.val,exists
}

func (c *Cache)reapLoop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		c.reap(time.Now(),c.interval)
	}
}

func (c *Cache)reap(now time.Time, prev time.Duration){
	c.mutex.Lock()
	defer c.mutex.Unlock()
	for e,entry := range c.urlcache {
		if entry.createdAt.Before(now.Add(-prev)){
			delete(c.urlcache,e)
		}
	}
}
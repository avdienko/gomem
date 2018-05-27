package gomem

import (
	"sync"
	"time"
)

type CacheContainer struct {
	defaultTimeExpiration time.Duration
	gcInterval            time.Duration
	mutex                 sync.RWMutex
	items                 map[string]cacheItem
}

type cacheItem struct {
	Value      string
	Created    time.Time
	Expiration int64
}

func New(defaultTimeExpiration, gcInterval time.Duration) *CacheContainer {
	item := make(map[string]cacheItem)
	cache := CacheContainer{
		defaultTimeExpiration: defaultTimeExpiration,
		gcInterval:            gcInterval,
		items:                 item,
	}

	if gcInterval > 0 {
		go cache.startGC()
	}

	return &cache
}

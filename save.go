package gomem

import "time"

func (cc *CacheContainer) Save(key, value string, ttl time.Duration) {
	var expiration int64

	if ttl == 0 {
		ttl = cc.defaultTimeExpiration
	}

	if ttl > 0 {
		expiration = time.Now().Add(ttl).UnixNano()
	}

	cc.mutex.Lock()

	cc.items[key] = cacheItem{
		Value:      value,
		Created:    time.Now(),
		Expiration: expiration,
	}

	cc.mutex.Unlock()
}

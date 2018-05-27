package gomem

import "time"

func (cc *CacheContainer) Get(key string) (string, bool) {
	cc.mutex.RLock()
	item, found := cc.items[key]
	cc.mutex.RUnlock()

	if !found {
		return "", false
	}

	if item.Expiration > 0 {
		if time.Now().UnixNano() > item.Expiration {
			return "", false
		}
	}

	return item.Value, true
}

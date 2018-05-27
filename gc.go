package gomem

import "time"

func (cc *CacheContainer) startGC() {
	for {
		<-time.After(cc.gcInterval)

		if cc.items == nil {
			return
		}

		if len(cc.items) > 0 {
			cc.mutex.RLock()
			for key, item := range cc.items {
				if time.Now().UnixNano() > item.Expiration && item.Expiration > 0 {
					delete(cc.items, key)
				}
			}
			cc.mutex.RUnlock()
		}
	}
}

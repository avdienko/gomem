package gomem

func (cc *CacheContainer) Delete(key string) {
	cc.mutex.Lock()
	delete(cc.items, key)
	cc.mutex.Unlock()
}

package gomem

func (cc *CacheContainer) Exist(key string) bool {
	if _, found := cc.items[key]; !found {
		return false
	}

	return true
}

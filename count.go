package gomem

func (cc *CacheContainer) Count() int {
	return len(cc.items)
}

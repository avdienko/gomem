package gomem

import (
	"errors"
)

func (cc *CacheContainer) Rename(key, newName string) error {
	cc.mutex.Lock()

	if _, found := cc.items[key]; !found {
		cc.mutex.Unlock()
		return errors.New("Key not found")
	}

	cc.items[newName] = cc.items[key]
	delete(cc.items, key)

	cc.mutex.Unlock()

	return nil
}

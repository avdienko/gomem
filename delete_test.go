package gomem

import (
	"testing"
	"time"
)

const (
	deleteKey1   = "delete_key1"
	deleteValue1 = "value1"
)

func TestDelete(t *testing.T) {
	cache := New(5*time.Second, 5*time.Second)

	cache.Save(deleteKey1, deleteValue1, 1*time.Minute)

	_, found := cache.Get(deleteKey1)
	if !found {
		t.Error("Error: key not found")
	}

	cache.Delete(deleteKey1)

	_, found = cache.Get(deleteKey1)
	if found {
		t.Error("Error: key found")
	}
}

package gomem

import (
	"testing"
	"time"
)

const (
	existKey1      = "exist_key1"
	existValueKey1 = "exist_value1"
)

func TestExist(t *testing.T) {
	cache := New(5*time.Second, 5*time.Second)

	found := cache.Exist(existKey1)
	if found {
		t.Error("Error: item found")
	}

	cache.Save(existKey1, existValueKey1, 1*time.Minute)

	found = cache.Exist(existKey1)
	if !found {
		t.Error("Error: item not found")
	}
}

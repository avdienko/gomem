package gomem

import (
	"testing"
	"time"
)

const (
	getKey1   = "get_key1"
	getValue1 = "value1"
)

func TestGet(t *testing.T) {
	cache := New(5*time.Second, 5*time.Second)

	_, found := cache.Get(getKey1)
	if found {
		t.Error("Error: key found")
	}

	cache.Save(getKey1, getValue1, 1*time.Minute)

	value, found := cache.Get(getKey1)
	if !found {
		t.Error("Error: key not found")
	}

	if value != getValue1 {
		t.Error("Error: value not valid")
	}
}

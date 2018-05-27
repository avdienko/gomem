package gomem

import (
	"testing"
	"time"
)

const (
	gcKey1   = "gc_key1"
	gcValue1 = "gc_value1"
)

func TestStartGc(t *testing.T) {
	cache := New(5*time.Second, 1*time.Second)

	cache.Save(gcKey1, gcValue1, 1*time.Millisecond)

	time.Sleep(2 * time.Second)

	found := cache.Exist(gcKey1)
	if found {
		t.Error("Error: item found")
	}
}

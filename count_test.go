package gomem

import (
	"testing"
	"time"
)

const (
	countKey1   = "count_key1"
	countValue1 = "value1"
	countKey2   = "count_key2"
	countValue2 = "value2"
)

func TestCount(t *testing.T) {
	cache := New(5*time.Second, 5*time.Second)

	if cache.Count() != 0 {
		t.Error("Error: the number of elements is not 0")
	}

	cache.Save(countKey1, countValue1, 1*time.Minute)

	if cache.Count() != 1 {
		t.Error("Error: the number of elements is not 1")
	}

	cache.Save(countKey2, countValue2, 1*time.Minute)
	if cache.Count() != 2 {
		t.Error("Error: the number of elements is not 2")
	}
}

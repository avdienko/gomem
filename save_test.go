package gomem

import (
	"testing"
	"time"
)

const (
	saveKey1   = "save_key1"
	saveValue1 = "value1"
)

func TestSave(t *testing.T) {
	cache := New(5*time.Second, 5*time.Second)

	_, found := cache.Get(saveKey1)
	if found {
		t.Error("Error: key found")
	}

	cache.Save(saveKey1, saveValue1, 1*time.Minute)

	value, found := cache.Get(saveKey1)
	if !found {
		t.Error("Error: key not found")
	}

	if value != saveValue1 {
		t.Error("Error: value not valid")
	}
}

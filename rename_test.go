package gomem

import (
	"testing"
	"time"
)

const (
	renameKey1   = "rename_key1"
	renameValue1 = "value1"
	renameKey2   = "rename_key2"
	renameValue2 = "value2"
)

func TestRename(t *testing.T) {
	cache := New(5*time.Second, 5*time.Second)

	cache.Save(renameKey1, renameValue1, 1*time.Minute)

	err := cache.Rename(renameKey1, renameKey2)
	if err != nil {
		t.Error(err.Error())
	}

	_, found := cache.Get(renameKey1)
	if found {
		t.Error("Error: key found")
	}

	value, found := cache.Get(renameKey2)
	if !found {
		t.Error("Error: key found")
	}

	if value != renameValue1 {
		t.Error("Error: value not valid")
	}
}

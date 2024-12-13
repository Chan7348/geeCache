package geecache

import (
	"reflect"
	"testing"
)

func TestGet(t *testing.T) {
	var function IGetter = Getter(func(key string) ([]byte, error) {
		return []byte(key), nil
	})

	expect := []byte("key")

	if v, _ := function.Get("key"); !reflect.DeepEqual(v, expect) {
		t.Errorf("Callback failed")
	}
}

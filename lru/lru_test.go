package lru

import (
	"reflect"
	"testing"
)

type String string

func (d String) Len() int {
	return len(d)
}

func TestGet(t *testing.T) {
	lru := New(int64(0), nil)
	lru.Add("key1", String("1234"))
	if v, ok := lru.Get("key1"); !ok || string(v.(String)) != "1234" {
		t.Fatalf("cache hit key1=1234 failed")
	}
	if _, ok := lru.Get("key2"); ok {
		t.Fatalf("cache miss key2 failed")
	}
}

func TestRemoveOldest(t *testing.T) {
	k1, k2, k3 := "key1", "key2", "key3"
	v1, v2, v3 := "value1", "value2", "value3"

	cap := len(k1 + k2 + v1 + v2)
	lru := New(int64(cap), nil)
	lru.Add(k1, String(v1))
	lru.Add(k2, String(v2))
	lru.Add(k3, String(v3))

	if _, ok := lru.Get("key1"); ok || lru.Len() != 2 {
		t.Fatalf("Remove oldest key1 failed")
	}
}

func TestOnEvicted(t *testing.T) {

	evictedKeysSlice := make([]string, 0, 2)
	onEvicted := func(key string, value Value) {
		evictedKeysSlice = append(evictedKeysSlice, key)
	}
	lru := New(int64(10), onEvicted)  // 创造一个最大10字节的lru
	lru.Add("key1", String("123456")) // 向lru增加一个新的kv
	lru.Add("k2", String("k2"))
	lru.Add("k3", String("k3"))
	lru.Add("k4", String("k4"))
	lru.Add("k5", String("k5"))

	expectedEvictedKeys := []string{"key1", "k2", "k3"}

	if !reflect.DeepEqual(expectedEvictedKeys, evictedKeysSlice) {
		t.Fatalf("Call onEvcited failed, expect keys equals to %s", expectedEvictedKeys)
	}
}

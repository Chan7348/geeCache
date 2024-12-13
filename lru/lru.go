package lru

import (
	"container/list"
)

type Cache struct {
	maxBytes   int64                         // 允许使用的最大内存
	nbytes     int64                         // 已使用的内存
	linkedList *list.List                    // 链表
	keyToNode  map[string]*list.Element      // key => node 的 mapping
	OnEvicted  func(key string, value Value) // 当某条记录被移除时触发的callback
}

// 链表节点的数据类型
type entry struct {
	key   string
	value Value
}

type Value interface {
	Len() int // 值所占内存大小
}

func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:   maxBytes,
		linkedList: list.New(),
		keyToNode:  make(map[string]*list.Element),
		OnEvicted:  onEvicted,
	}
}

func (cache *Cache) Get(key string) (value Value, ok bool) {
	if element, ok := cache.keyToNode[key]; ok {
		cache.linkedList.MoveToFront(element)
		kv := element.Value.(*entry)
		return kv.value, true
	}
	return
}

func (cache *Cache) RemoveOldest() {
	element := cache.linkedList.Back()
	if element != nil {
		cache.linkedList.Remove(element)
		kv := element.Value.(*entry)
		delete(cache.keyToNode, kv.key)
		cache.nbytes -= int64(len(kv.key)) + int64(kv.value.Len())
		if cache.OnEvicted != nil {
			cache.OnEvicted(kv.key, kv.value)
		}
	}
}

func (cache *Cache) Add(key string, value Value) {
	if element, ok := cache.keyToNode[key]; ok {
		cache.linkedList.MoveToFront(element)
		kv := element.Value.(*entry)
		cache.nbytes += int64(value.Len()) - int64(kv.value.Len()) // 更新总大小
		kv.value = value
	} else {
		element := cache.linkedList.PushFront(&entry{key, value})
		cache.keyToNode[key] = element
		cache.nbytes += int64(len(key)) + int64(value.Len())
	}

	for cache.maxBytes != 0 && cache.maxBytes < cache.nbytes {
		cache.RemoveOldest()
	}
}

func (cache *Cache) Len() int {
	return cache.linkedList.Len()
}

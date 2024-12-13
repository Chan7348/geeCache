package geecache

// 并发控制
import (
	"sync"

	"github.com/Chan7348/geecache/lru"
)

type cache struct {
	mutex      sync.Mutex
	lru        *lru.Cache
	cacheBytes int64
}

func (cache *cache) add(key string, value ByteView) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	if cache.lru == nil {
		cache.lru = lru.New(cache.cacheBytes, nil)
	}

	cache.lru.Add(key, value)
}

func (cache *cache) get(key string) (value ByteView, ok bool) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	if cache.lru == nil {
		return
	}

	rawValue, ok := cache.lru.Get(key)
	return rawValue.(ByteView), ok
}

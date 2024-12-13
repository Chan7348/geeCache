package geecache

import "sync"

// 与外界交互，控制缓存存储和获取的主流程

type IGetter interface {
	Get(key string) ([]byte, error)
}

type Getter func(key string) ([]byte, error)

// type Group struct {
// 	name      string
// 	getter    IGetter
// 	mainCache cache
// }

func (function Getter) Get(key string) ([]byte, error) {
	return function(key)
}

var (
	mutex sync.RWMutex
)

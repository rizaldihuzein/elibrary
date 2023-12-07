package memcache

import (
	"errors"
	"sync"
)

var (
	ErrNil = errors.New("nil data")
)

type (
	CacheType map[string]interface{}

	MemoryCache struct {
		Persistent CacheType
		mutex      *sync.RWMutex
	}
)

func New() *MemoryCache {
	return &MemoryCache{
		Persistent: make(CacheType),
		mutex:      &sync.RWMutex{},
	}
}

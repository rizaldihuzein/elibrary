package memcache

import (
	"context"
	"errors"
)

func (m *MemoryCache) Upsert(ctx context.Context, key string, data interface{}) (err error) {
	select {
	case <-ctx.Done():
		return errors.New("context done by cancel or timeout")
	default:
	}

	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.Persistent[key] = data

	return
}

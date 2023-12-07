package memcache

import (
	"context"
	"errors"
)

func (m *MemoryCache) Evict(ctx context.Context, key string) (err error) {
	select {
	case <-ctx.Done():
		return errors.New("context done by cancel or timeout")
	default:
	}

	m.mutex.Lock()
	defer m.mutex.Unlock()

	delete(m.Persistent, key)

	return
}

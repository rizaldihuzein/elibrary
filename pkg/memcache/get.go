package memcache

import (
	"context"
	"errors"
)

func (m *MemoryCache) Get(ctx context.Context, key string) (data interface{}, err error) {
	select {
	case <-ctx.Done():
		return nil, errors.New("context done by cancel or timeout")
	default:
	}

	m.mutex.RLock()
	defer m.mutex.RUnlock()

	data, ok := m.Persistent[key]
	if !ok || data == nil {
		return nil, ErrNil
	}

	return
}

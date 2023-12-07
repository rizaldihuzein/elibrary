package memorycache

import (
	"context"
	"cosmart-library/domain"
	"cosmart-library/pkg/logger"
	"cosmart-library/pkg/memcache"
	"time"
)

func (m *memorycache) UpsertBookOrder(ctx context.Context, pickup domain.PickupOrder) (err error) {
	err = m.persistentDriver.Upsert(ctx, pickup.ID, pickup)
	if err != nil {
		logger.Error("[UpsertBookOrder][persistentDriver] error calling upsert", err.Error())
	}
	return
}

func (m *memorycache) GetBookOrder(ctx context.Context, id string) (pickup domain.PickupOrder, err error) {
	pickupInf, err := m.persistentDriver.Get(ctx, id)
	if err == memcache.ErrNil {
		return pickup, nil
	}
	if err != nil {
		logger.Error("[GetBookOrder][persistentDriver] error calling get", err.Error())
		return
	}

	pickup, _ = pickupInf.(domain.PickupOrder)
	return
}

func (m *memorycache) StoreIdempotency(ctx context.Context, key string, ttl time.Time) (err error) {
	err = m.persistentDriver.Upsert(ctx, "id_"+key, ttl)
	if err != nil {
		logger.Error("[StoreIdempotency][persistentDriver] error calling upsert", err.Error())
	}
	return
}

func (m *memorycache) GetIdempotency(ctx context.Context, key string) (ttl time.Time, err error) {
	ttlInf, err := m.persistentDriver.Get(ctx, "id_"+key)
	if err == memcache.ErrNil {
		return ttl, nil
	}
	if err != nil {
		logger.Error("[GetIdempotency][persistentDriver] error calling get", err.Error())
		return
	}

	ttl, _ = ttlInf.(time.Time)
	return
}

func (m *memorycache) EvictIdempotency(ctx context.Context, key string) (ttl time.Time, err error) {
	err = m.persistentDriver.Evict(ctx, "id_"+key)
	if err != nil {
		logger.Error("[EvictIdempotency][persistentDriver] error calling evict", err.Error())
	}
	return
}

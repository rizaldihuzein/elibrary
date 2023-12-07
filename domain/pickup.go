package domain

import (
	"context"
	"time"
)

//go:generate mockgen -source=pickup.go -destination=../mocks/pickup.go -package=mocks
type (
	PickupOrder struct {
		ID         string
		Books      []Book
		PickupDate time.Time
		ReturnDate time.Time
	}

	PickupRepositoryInterface interface {
		UpsertBookOrder(ctx context.Context, pickup PickupOrder) error
		GetBookOrder(ctx context.Context, key string) (pickup PickupOrder, err error)
		StoreIdempotency(ctx context.Context, key string, ttl time.Time) (err error)
		GetIdempotency(ctx context.Context, key string) (ttl time.Time, err error)
		EvictIdempotency(ctx context.Context, key string) (ttl time.Time, err error)
	}

	PickupUsecaseInterface interface {
		MakeBookOrder(ctx context.Context, pickup PickupOrder) (string, error)
		GetBookOrder(ctx context.Context, key string) (pickup PickupOrder, err error)
	}
)

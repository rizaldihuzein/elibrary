package memorycache

import (
	"context"
	"cosmart-library/domain"
)

//go:generate mockgen -source=init.go -destination=../../../mocks/pickup/repository/repository.go -package=mocks_repository
type (
	DriverInterface interface {
		Upsert(ctx context.Context, key string, data interface{}) (err error)
		Get(ctx context.Context, key string) (data interface{}, err error)
		Evict(ctx context.Context, key string) (err error)
	}

	memorycache struct {
		persistentDriver DriverInterface
	}
)

func New(driver DriverInterface) domain.PickupRepositoryInterface {
	return &memorycache{
		persistentDriver: driver,
	}
}

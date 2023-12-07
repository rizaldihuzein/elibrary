package usecase

import (
	"context"
	"cosmart-library/domain"
	"cosmart-library/pick-up/types"
	"cosmart-library/pkg/logger"
	"time"

	"github.com/google/uuid"
)

func (u *usecase) MakeBookOrder(ctx context.Context, pickup domain.PickupOrder) (id string, err error) {
	ctxID := ctx.Value(types.RequestIDKey)
	idempotencyKey, _ := ctxID.(string)
	if idempotencyKey != "" {
		ttl, err := u.repository.GetIdempotency(ctx, idempotencyKey)
		if err != nil {
			logger.Error("[MakeBookOrder] failed to get idempotency", err.Error())
			return "", err
		}
		if !ttl.IsZero() && !ttl.Before(time.Now()) {
			logger.Warn("[MakeBookOrder] request has been done")
			id = pickup.ID
			if id == "" {
				id = "idempotent request"
			}
			return id, nil
		}
		defer u.repository.StoreIdempotency(ctx, idempotencyKey, time.Now().Add(1*time.Hour))
	} else {
		logger.Warn("[MakeBookOrder] no request id, ignoring idempotency")
	}

	if pickup.ID == "" {
		pickup.ID = uuid.New().String()
	}

	if (pickup.ReturnDate.Sub(pickup.PickupDate)) > 7*24*time.Hour || pickup.ReturnDate.IsZero() {
		pickup.ReturnDate = pickup.PickupDate.Add(7 * 24 * time.Hour)
	}

	err = u.repository.UpsertBookOrder(ctx, pickup)
	if err != nil {
		logger.Error("[MakeBookOrder][UpsertBookOrder] error calling repository", err.Error())
		return
	}

	return pickup.ID, err
}

func (u *usecase) GetBookOrder(ctx context.Context, id string) (order domain.PickupOrder, err error) {
	order, err = u.repository.GetBookOrder(ctx, id)
	if err != nil {
		logger.Error("[GetBookOrder][GetBookOrder] error calling repository", err.Error())
	}

	return
}

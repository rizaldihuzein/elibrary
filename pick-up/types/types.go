package types

import (
	"cosmart-library/domain"
)

type CtxKey string

const (
	TimeFormat          = "2006-01-02 15:04:05"
	RequestIDKey CtxKey = "Request-ID"
)

type (
	OrderRequest struct {
		ID            string        `json:"id"`
		Books         []domain.Book `json:"books,omitempty"`
		PickupDateStr string        `json:"pickup_date,omitempty"`
		ReturnDateStr string        `json:"return_date,omitempty"`
	}

	OrderResponse struct {
		ID string `json:"id"`
	}
)

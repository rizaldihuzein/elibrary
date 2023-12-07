package http

import (
	"context"
	"cosmart-library/domain"
	"cosmart-library/pick-up/types"
	"cosmart-library/pkg/logger"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/microcosm-cc/bluemonday"
)

// ArticleHandler  represent the httphandler for article
type pickupHandler struct {
	pickupUsecase domain.PickupUsecaseInterface
	sanitizer     *bluemonday.Policy
}

func InitHTTPHandler(mux *http.ServeMux, uc domain.PickupUsecaseInterface) {
	handler := &pickupHandler{
		pickupUsecase: uc,
		sanitizer:     bluemonday.UGCPolicy(),
	}

	mux.Handle("/book/order", domain.POST(handler.CreateBookOrderHandler))
	mux.Handle("/order", domain.GET(handler.GetOrderHandler))
}

func (p *pickupHandler) GetOrderHandler(w http.ResponseWriter, r *http.Request) (resp domain.GeneralResponse, err error) {
	id := strings.ToLower(p.sanitizer.Sanitize(r.URL.Query().Get("id")))

	order, err := p.pickupUsecase.GetBookOrder(r.Context(), id)
	if err != nil {
		logger.Error("[GetOrderHandler][GetBookOrder] error calling usecase", err.Error())
		return
	}

	data := types.OrderRequest{
		ID:    order.ID,
		Books: order.Books,
	}

	if !order.PickupDate.IsZero() {
		data.PickupDateStr = order.PickupDate.Format(types.TimeFormat)
	}
	if !order.ReturnDate.IsZero() {
		data.ReturnDateStr = order.ReturnDate.Format(types.TimeFormat)
	}

	resp.Data = data

	return
}

func (p *pickupHandler) CreateBookOrderHandler(w http.ResponseWriter, r *http.Request) (resp domain.GeneralResponse, err error) {
	if !domain.ValidateJSONContentType(r) {
		resp.Code = http.StatusBadRequest
		resp.Data = "Unknown request"
		return
	}

	requestID := strings.TrimSpace(r.Header.Get("Request-ID"))
	if requestID == "" {
		logger.Warn("[CreateBookOrder] no request ID provided, idempotency will be omitted")
		requestID = uuid.New().String()
	}

	request := types.OrderRequest{}
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&request)
	if err != nil {
		logger.Error("[CreateBookOrder] error decoding json", err.Error())
		resp.Code = http.StatusBadRequest
		resp.Message = "Invalid request"
		resp.DeveloperMsg = append(resp.DeveloperMsg, err.Error())
		resp.Data = "error"

		return resp, nil
	}

	order := domain.PickupOrder{
		Books: request.Books,
	}

	request.PickupDateStr = strings.TrimSpace(request.PickupDateStr)
	request.ReturnDateStr = strings.TrimSpace(request.ReturnDateStr)
	if request.PickupDateStr == "" {
		resp.Code = http.StatusBadRequest
		resp.Data = "Invalid request"
		resp.Message = "Pickup time cannot be empty"
		return
	}

	if request.PickupDateStr != "" {
		order.PickupDate, err = time.Parse(types.TimeFormat, request.PickupDateStr)
		if err != nil {
			logger.Error("[CreateBookOrder] error parsing pickup time", err.Error())
			resp.Code = http.StatusBadRequest
			resp.Data = "Invalid request"
			resp.DeveloperMsg = append(resp.DeveloperMsg, err.Error())
			resp.Message = "Invalid Pickup time format"

			return resp, nil
		}
	}

	if request.ReturnDateStr != "" {
		order.ReturnDate, err = time.Parse(types.TimeFormat, request.ReturnDateStr)
		if err != nil {
			logger.Error("[CreateBookOrder] error parsing return time", err.Error())
			resp.Code = http.StatusBadRequest
			resp.Data = "Invalid request"
			resp.DeveloperMsg = append(resp.DeveloperMsg, err.Error())
			resp.Message = "Invalid Return time format"

			return resp, nil
		}
	}

	if order.ReturnDate.Before(order.PickupDate) && !order.ReturnDate.IsZero() {
		resp.Code = http.StatusBadRequest
		resp.Data = "Invalid request"
		resp.Message = "Return time cannot be less than Pickup time"

		return
	}

	ctx := context.WithValue(r.Context(), types.RequestIDKey, requestID)
	id, err := p.pickupUsecase.MakeBookOrder(ctx, order)
	if err != nil {
		logger.Error("[CreateBookOrder][MakeBookOrder] error calling usecase", err.Error())
		return
	}

	resp.Data = types.OrderResponse{
		ID: id,
	}

	return
}

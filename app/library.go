package main

import (
	libraryhttpDelivery "cosmart-library/library/delivery/http"
	libraryRepository "cosmart-library/library/repository"
	libraryUsecase "cosmart-library/library/usecase"
	pickuphttpDelivery "cosmart-library/pick-up/delivery/http"
	pickupMemcache "cosmart-library/pick-up/repository/memory-cache"
	pickupUsecase "cosmart-library/pick-up/usecase"
	"cosmart-library/pkg/logger"
	"cosmart-library/pkg/memcache"
	"net/http"
)

func main() {
	libraryRepo := libraryRepository.New()
	libraryUC := libraryUsecase.New(libraryRepo)

	memcache := memcache.New()
	orderRepo := pickupMemcache.New(memcache)
	orderUsecase := pickupUsecase.New(orderRepo)

	mux := http.NewServeMux()
	libraryhttpDelivery.InitHTTPHandler(mux, libraryUC)
	pickuphttpDelivery.InitHTTPHandler(mux, orderUsecase)

	port := ":8080"
	server := &http.Server{
		Addr:    port,
		Handler: mux,
	}
	logger.Info("Starting server on " + port)
	server.ListenAndServe()
}

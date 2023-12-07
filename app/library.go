package main

import (
	httpDelivery "cosmart-library/library/delivery/http"
	"cosmart-library/library/repository"
	"cosmart-library/library/usecase"
	"cosmart-library/pkg/logger"
	"net/http"
)

func main() {
	libraryRepo := repository.New()
	libraryUsecase := usecase.New(libraryRepo)
	mux := http.NewServeMux()
	httpDelivery.InitHTTPHandler(mux, libraryUsecase)

	port := ":8080"
	server := &http.Server{
		Addr:    port,
		Handler: mux,
	}
	logger.Info("Starting server on " + port)
	server.ListenAndServe()
}

package repository

import (
	"context"
	"cosmart-library/domain"
	restapi "cosmart-library/library/repository/rest-api"
)

//go:generate mockgen -source=repository.go -destination=../../mocks/library/repository/repository.go -package=mocks_repository
type (
	APIInterface interface {
		FetchBookListBySubject(ctx context.Context, subject string) (books []domain.Book, err error)
	}

	repository struct {
		api APIInterface
	}
)

func New() domain.LibraryRepositoryInterface {
	api := restapi.New()
	return &repository{
		api: api,
	}
}

func (r *repository) FetchRawBooksBySubject(ctx context.Context, subject string) (books []domain.Book, err error) {
	return r.api.FetchBookListBySubject(ctx, subject)
}

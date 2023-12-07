package domain

import "context"

//go:generate mockgen -source=library.go -destination=../mocks/library.go -package=mocks
type (
	Book struct {
		ID            string   `json:"id"`
		Subject       string   `json:"subject"`
		Title         string   `json:"title"`
		Author        []string `json:"authors"`
		EditionNumber string   `json:"edition"`
	}

	GetBookListRequest struct {
		Subject string
	}

	GetBookListResponse struct {
		Books []Book `json:"books"`
		Page  int    `json:"page"`
	}

	LibraryRepositoryInterface interface {
		FetchRawBooksBySubject(ctx context.Context, subject string) (books []Book, err error)
	}

	LibraryUsecaseInterface interface {
		GetBookList(ctx context.Context, request GetBookListRequest) (response GetBookListResponse, err error)
	}
)

package usecase

import "cosmart-library/domain"

type usecase struct {
	repository domain.LibraryRepositoryInterface
}

func New(repository domain.LibraryRepositoryInterface) domain.LibraryUsecaseInterface {
	return &usecase{
		repository: repository,
	}
}

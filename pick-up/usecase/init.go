package usecase

import "cosmart-library/domain"

type usecase struct {
	repository domain.PickupRepositoryInterface
}

func New(repository domain.PickupRepositoryInterface) domain.PickupUsecaseInterface {
	return &usecase{
		repository: repository,
	}
}

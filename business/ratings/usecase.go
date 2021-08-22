package ratings

import (
	"context"
	"time"
)

type RatingUseCase struct {
	RatingRepository Repository
	ContextTimeout time.Duration
}

func NewRatingUsecase(rt Repository, timeout time.Duration) UseCase {
	return &RatingUseCase{
		RatingRepository: rt,
		ContextTimeout: timeout,
	}
}

func (ru *RatingUseCase) Store(ctx context.Context, catDomain *Domain) error {
	ctx, cancel := context.WithTimeout(ctx, ru.ContextTimeout)
	defer cancel()

	err := ru.RatingRepository.Store(ctx, catDomain)
	if err != nil {
		return err
	}
	return nil
}

func (ru *RatingUseCase) GetById(ctx context.Context, id int) ([]Domain, error) {
	result, err := ru.RatingRepository.GetById(ctx, id)
	if err != nil {
		return []Domain{}, err
	}
	return result, err
}

func (ru *RatingUseCase) Update(ctx context.Context, RatingDomain *Domain, id string) error {
	ctx, cancel := context.WithTimeout(ctx, ru.ContextTimeout)
	defer cancel()

	err := ru.RatingRepository.Update(ctx, RatingDomain, id)
	if err != nil {
		return err
	}
	return nil
}

func (ru *RatingUseCase) Delete(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, ru.ContextTimeout)
	defer cancel()

	err := ru.RatingRepository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

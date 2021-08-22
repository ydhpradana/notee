package categories

import (
	"context"
	"time"
)

type CategoryUseCase struct {
	CategoryRepository     Repository
	ContextTimeout time.Duration
}

func NewCategoryUsecase(cr Repository,  timeout time.Duration) UseCase {
	return &CategoryUseCase{
		CategoryRepository:     cr,
		ContextTimeout: timeout,
	}
}

// func (cu *CategoryUseCase) GetByActive(ctx context.Context, active bool) ([]Domain, error) {
// 	activeCat := "false"
// 	if active {
// 		activeCat = "true"
// 	}
// 	result, err := cu.CategoryRepository.Find(ctx, activeCat)
// 	if err != nil {
// 		return []Domain{}, err
// 	}

// 	return result, nil
// }

func (cu *CategoryUseCase) Store(ctx context.Context, catDomain *Domain) error {
	ctx, cancel := context.WithTimeout(ctx, cu.ContextTimeout)
	defer cancel()

	err := cu.CategoryRepository.Store(ctx, catDomain)
	if err != nil {
		return err
	}
	return nil
}

func (cu *CategoryUseCase) Update(ctx context.Context, catDomain *Domain, id string) error {
	ctx, cancel := context.WithTimeout(ctx, cu.ContextTimeout)
	defer cancel()

	err := cu.CategoryRepository.Update(ctx, catDomain, id)
	if err != nil {
		return err
	}
	return nil
}

func (cu *CategoryUseCase) Delete(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, cu.ContextTimeout)
	defer cancel()

	err := cu.CategoryRepository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (cu *CategoryUseCase) GetById(ctx context.Context, id string) (Domain, error) {
	result, err := cu.CategoryRepository.FindById(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	return result, err
}

func (cu *CategoryUseCase) GetAll(ctx context.Context) ([]Domain, error) {
	result, err := cu.CategoryRepository.FindAll(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return result, err
}

func (cu *CategoryUseCase) GetByName(ctx context.Context, search string) ([]Domain, error) {
	result, err := cu.CategoryRepository.Find(ctx, search)
	if err != nil {
		return []Domain{}, err
	}
	return result, err
}

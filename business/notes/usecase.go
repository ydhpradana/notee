package notes

import (
	"context"
	"time"
)

type NoteUseCase struct {
	NoteRepository Repository
	ContextTimeout     time.Duration
}

func NewNoteUsecase(nt Repository, timeout time.Duration) UseCase {
	return &NoteUseCase{
		NoteRepository: nt,
		ContextTimeout:     timeout,
	}
}

func (nu *NoteUseCase) Store(ctx context.Context, catDomain *Domain) error {
	ctx, cancel := context.WithTimeout(ctx, nu.ContextTimeout)
	defer cancel()

	err := nu.NoteRepository.Store(ctx, catDomain)
	if err != nil {
		return err
	}
	return nil
}

func (nu *NoteUseCase) GetById(ctx context.Context, id string) (Domain, error) {
	result, err := nu.NoteRepository.FindById(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	return result, err
}

func (nu *NoteUseCase) GetAll(ctx context.Context) ([]Domain, error) {
	result, err := nu.NoteRepository.FindAll(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return result, err
}

func (nu *NoteUseCase) GetByName(ctx context.Context, search string) ([]Domain, error) {
	result, err := nu.NoteRepository.Find(ctx, search)
	if err != nil {
		return []Domain{}, err
	}
	return result, err
}

func (nu *NoteUseCase) GetByCatId(ctx context.Context, id string) ([]Domain, error) {
	result, err := nu.NoteRepository.FindByCatId(ctx, id)
	if err != nil {
		return []Domain{}, err
	}
	return result, err
}

func (nu *NoteUseCase) GetByUserId(ctx context.Context, id string) ([]Domain, error) {
	result, err := nu.NoteRepository.FindByUserId(ctx, id)
	if err != nil {
		return []Domain{}, err
	}
	return result, err
}

func (nu *NoteUseCase) GetByIsFree(ctx context.Context, free string) ([]Domain, error) {
	result, err := nu.NoteRepository.FindByIsFree(ctx, free)
	if err != nil {
		return []Domain{}, err
	}
	return result, err
}

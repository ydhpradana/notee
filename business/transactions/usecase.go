package transactions

import (
	"context"
	"time"
)

type TransactionUseCase struct {
	TransactionRepository Repository
	ContextTimeout time.Duration
}

func NewTransactionUsecase(tr Repository, timeout time.Duration) UseCase {
	return &TransactionUseCase{
		TransactionRepository: tr,
		ContextTimeout: timeout,
	}
}

func (tu *TransactionUseCase) Store(ctx context.Context, transactionDomain *Domain) error {
	ctx, cancel := context.WithTimeout(ctx, tu.ContextTimeout)
	defer cancel()

	err := tu.TransactionRepository.Store(ctx, transactionDomain)
	if err != nil {
		return err
	}
	return nil
}

func (tu *TransactionUseCase) GetTrx(ctx context.Context, UserID int) ([]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, tu.ContextTimeout)
	defer cancel()

	resp, err := tu.TransactionRepository.GetTrx(ctx, UserID)
	if err != nil {
		return []Domain{}, err
	}
	return resp, nil
}
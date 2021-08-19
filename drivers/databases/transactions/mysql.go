package transactions

import (
	"context"
	"notee/business/transactions"

	"gorm.io/gorm"
)

type MysqlTransactionRepository struct {
	Conn *gorm.DB
}

func NewMysqlTransactionRepository(conn *gorm.DB) transactions.Repository {
	return &MysqlTransactionRepository{
		Conn: conn,
	}
}

func (repository *MysqlTransactionRepository) Store(ctx context.Context, transactionDomain *transactions.Domain) error {
	rec := fromDomain(*transactionDomain)

	result := repository.Conn.Create(rec)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository *MysqlTransactionRepository) GetTrx(ctx context.Context, userID int) ([]transactions.Domain, error) {
	rec := []Transaction{}

	err := repository.Conn.Preload("Users").Preload("Notes").Preload("Notes.Categories").Preload("Notes.Users").Where("transactions.user_id = ?", userID).Find(&rec).Error
	//err := repository.Conn.Preload(clause.Associations).Where("transactions.user_id = ?", userID).Find(&rec).Error
	if err != nil {
		return []transactions.Domain{}, err
	}

	transactionDomain := []transactions.Domain{}
	for _, value := range rec {
		transactionDomain = append(transactionDomain, value.toDomain())
	}

	return transactionDomain, nil
}

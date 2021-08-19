package transactions

import (
	"notee/business/transactions"
	"notee/drivers/databases/notes"
	"notee/drivers/databases/users"
	"time"
)

type Transaction struct {
	Id         uint                	`gorm:"primarykey" json:"id"`
	CreatedAt  time.Time           	`json:"createdAt"`
	UpdatedAt  time.Time           	`json:"updatedAt"`
	DeletedAt  time.Time           	`gorm:"index" json:"deletedAt"`
	Code      	string              `json:"code"`
	TotalPrice  int         		`json:"total_price"`
	NoteId 		int                 `json:"category_id"`
	Notes 		notes.Note 			`gorm:"foreignKey:NoteId;references:Id"`
	UserId     int                 	`json:"user_id"`
	Users      users.User          	`gorm:"foreignKey:UserId;references:ID"`
}

func (rec *Transaction) toDomain() transactions.Domain {
	return transactions.Domain{
		Id:         rec.Id,
		CreatedAt:  rec.CreatedAt,
		UpdatedAt:  rec.UpdatedAt,
		DeletedAt:  rec.DeletedAt,
		Code:      rec.Code,
		NoteId: rec.NoteId,
		Note:   rec.Notes,
		UserId:     rec.UserId,
		User:       rec.Users,
		TotalPrice:       rec.TotalPrice,
	}
}

func fromDomain(transaction transactions.Domain) *Transaction {
	return &Transaction{
		Id:         transaction.Id,
		CreatedAt:  transaction.CreatedAt,
		UpdatedAt:  transaction.UpdatedAt,
		DeletedAt:  transaction.DeletedAt,
		Code:      transaction.Code,
		NoteId: transaction.NoteId,
		UserId:     transaction.UserId,
		TotalPrice:     transaction.TotalPrice,
	}
}

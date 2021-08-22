package response

import (
	"notee/business/transactions"
	"time"
)

type Transaction struct {
	Id           uint      `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at"`
	Code       	 string    `json:"code"`
	TotalPrice        int       `json:"total_price"`
	UserId       int       `json:"user_id"`
	UserName     string    `json:"user_name"`
	NoteId       int       `json:"note_id"`
	NoteTitle     string    `json:"note_title"`
	// Rating float64 `json:"rating"`
}

func FromDomain(domain transactions.Domain) Transaction {
	return Transaction{
		Id:           domain.Id,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
		DeletedAt:    domain.DeletedAt,
		Code:         domain.Code,
		TotalPrice:        domain.TotalPrice,
		UserId:       domain.UserId,
		UserName:     domain.User.Name,
		NoteId: 	  domain.NoteId,
		NoteTitle: 	  domain.Note.Title,
		// Rating:  domain.Rating,
		//Token:     domain.Token,
	}
}

package response

import (
	"notee/business/notes"
	"time"
)

type Note struct {
	Id        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	Title      string    `json:"title"`
	Note     string    `json:"note"`
	CategoryId int `json:"category_id"`
	UserId int `json:"user_id"`
	CategoryName string `json:"category_name"`
	UserName string `json:"user_name"`
	// Rating float64 `json:"rating"`
}

func FromDomain(domain notes.Domain) Note {
	return Note{
		Id:        domain.Id,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		Title:      domain.Title,
		Note:     domain.Note,
		CategoryId:     domain.CategoryId,
		UserId:     domain.UserId,
		CategoryName:  domain.Category.Name,
		UserName:  domain.User.Name,
		// Rating:  domain.Rating,
		//Token:     domain.Token,
	}
}

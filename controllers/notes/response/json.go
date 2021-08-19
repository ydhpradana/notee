package response

import (
	"notee/business/notes"
	"notee/drivers/databases/categories"
	"notee/drivers/databases/users"
	"time"
)

type Note struct {
	Id        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
	Title      string    `json:"title"`
	Note     string    `json:"note"`
	CategoryId int `json:"category_id"`
	UserId int `json:"user_id"`
	User users.User `json:"user"`
	Category categories.Category `json:"category"`
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
		User:  domain.User,
		Category:  domain.Category,
		//Token:     domain.Token,
	}
}

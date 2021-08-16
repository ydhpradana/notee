package notes

import (
	"notee/business/notes"
	"notee/drivers/databases/categories"
	"notee/drivers/databases/users"
	"time"
)

type Note struct {
	Id        	uint      `gorm:"primarykey" json:"id"`
	CreatedAt 	time.Time `json:"createdAt"`
	UpdatedAt 	time.Time `json:"updatedAt"`
	DeletedAt 	time.Time `gorm:"index" json:"deletedAt"`
	Title      	string    `json:"title"`
	Note      	string    `json:"note"`
	CategoryId  int		  `json:"category_id"`
	Categories  categories.Category `gorm:"foreignKey:CategoryId;references:Id"`
	UserId		int		  `json:"user_id"`
	Users      users.User `gorm:"foreignKey:UserId;references:Id"`
	IsFree      bool      `json:"is_free"`
}

func (rec *Note) toDomain() notes.Domain {
	return notes.Domain{
		Id:        	rec.Id,
		CreatedAt: 	rec.CreatedAt,
		UpdatedAt: 	rec.UpdatedAt,
		DeletedAt: 	rec.DeletedAt,
		Title:      rec.Title,
		Note:      	rec.Note,
		CategoryId: rec.CategoryId,
		UserId: 	rec.UserId,
		IsFree: 	rec.IsFree,
	}
}

func fromDomain(note notes.Domain) *Note {
	return &Note{
		Id:        		note.Id,
		CreatedAt: 		note.CreatedAt,
		UpdatedAt: 		note.UpdatedAt,
		DeletedAt: 		note.DeletedAt,
		Title:      	note.Title,
		Note:      		note.Note,
		CategoryId:     note.CategoryId,
		UserId:      	note.UserId,
		IsFree:      	note.IsFree,
	}
}

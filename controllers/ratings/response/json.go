package response

import (
	"notee/business/ratings"
	"time"
)

type Rating struct {
	Id         uint                `json:"id"`
	CreatedAt  time.Time           `json:"created_at"`
	UpdatedAt  time.Time           `json:"updated_at"`
	DeletedAt  time.Time           `json:"deleted_at"`
	Rating      int                `json:"rating"`
	Review       string            `json:"review"`
	UserId     int                 `json:"user_id"`
	NoteId     int                 `json:"note_id"`
	UserName       string          `json:"user_name"`
}

func FromDomain(domain ratings.Domain) Rating {
	return Rating{
		Id:         domain.Id,
		CreatedAt:  domain.CreatedAt,
		UpdatedAt:  domain.UpdatedAt,
		DeletedAt:  domain.DeletedAt,
		Rating:      domain.Rating,
		Review:       domain.Review,
		UserId:     domain.UserId,
		NoteId:     domain.NoteId,
		UserName:       domain.User.Name,
		//Token:     domain.Token,
	}
}

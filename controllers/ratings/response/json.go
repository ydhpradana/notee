package response

import (
	"notee/business/ratings"
	"notee/drivers/databases/users"
	"time"
)

type Rating struct {
	Id         uint                `json:"id"`
	CreatedAt  time.Time           `json:"createdAt"`
	UpdatedAt  time.Time           `json:"updatedAt"`
	DeletedAt  time.Time           `json:"deletedAt"`
	Rating      int                `json:"rating"`
	Review       string            `json:"review"`
	UserId     int                 `json:"user_id"`
	NoteId     int                 `json:"note_id"`
	User       users.User          `json:"user"`
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
		User:       domain.User,
		//Token:     domain.Token,
	}
}

package ratings

import (
	"notee/business/ratings"
	"notee/drivers/databases/users"
	"time"
)

type Rating struct {
	Id         uint                `gorm:"primarykey" json:"id"`
	CreatedAt  time.Time           `json:"createdAt"`
	UpdatedAt  time.Time           `json:"updatedAt"`
	DeletedAt  time.Time           `gorm:"index" json:"deletedAt"`
	Rating      int              	`json:"rating"`
	Review     string              `json:"review"`
	NoteId 		int                 `json:"note_id"`
	UserId     int                 	`json:"user_id"`
	Users      users.User          	`gorm:"foreignKey:UserId;references:ID"`
}

func (rec *Rating) toDomain() ratings.Domain {
	return ratings.Domain{
		Id:         rec.Id,
		CreatedAt:  rec.CreatedAt,
		UpdatedAt:  rec.UpdatedAt,
		DeletedAt:  rec.DeletedAt,
		Rating:      rec.Rating,
		Review:       rec.Review,
		NoteId: rec.NoteId,
		UserId:     rec.UserId,
		User:       rec.Users,
	}
}

func fromDomain(rating ratings.Domain) *Rating {
	return &Rating{
		Id:         rating.Id,
		CreatedAt:  rating.CreatedAt,
		UpdatedAt:  rating.UpdatedAt,
		DeletedAt:  rating.DeletedAt,
		Rating:      rating.Rating,
		Review:       rating.Review,
		UserId:     rating.UserId,
		NoteId:     rating.NoteId,
	}
}

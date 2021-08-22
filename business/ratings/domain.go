package ratings

import (
	"context"
	"notee/drivers/databases/users"
	"time"
)

type Domain struct {
	Id         uint
	Rating     int
	NoteId     int
	UserId     int
	User       users.User
	Review     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
}

type UseCase interface {
	Store(ctx context.Context, noteDomain *Domain) error
	GetById(ctx context.Context, id int) ([]Domain, error)
	Update(ctx context.Context, ratingDomain *Domain, id string) error
	Delete(ctx context.Context, id string) error
	// GetByActive(ctx context.Context, active bool) ([]Domain, error)
}

type Repository interface {
	Store(ctx context.Context, noteDomain *Domain) error
	GetById(ctx context.Context, id int) ([]Domain, error)
	Update(ctx context.Context, ratingDomain *Domain, id string) error
	Delete(ctx context.Context, id string) error
}

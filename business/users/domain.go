package users

import (
	"context"
	"time"
)

type Domain struct {
	Id        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Name      string
	Email     string
	Password  string
	Token     string
}

type UseCase interface {
	CreateToken(ctx context.Context, username, password string) (string, error)
	Login(ctx context.Context, email, password string) (Domain, error)
	GetById(c context.Context, id string) (Domain, error)
	Store(ctx context.Context, data *Domain) error
}

type Repository interface {
	GetByEmail(ctx context.Context, email string) (Domain, error)
	Login(ctx context.Context, email, password string) (Domain, error)
	Store(ctx context.Context, data *Domain) error
	GetById(ctx context.Context, id string) (Domain, error)
}

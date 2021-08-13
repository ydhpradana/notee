package categories

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
}

type UseCase interface{
	Fetch(ctx context.Context, page, perpage string) ([]Domain, error)
	GetByID(ctx context.Context, catID int) (Domain, error)
	Store(ctx context.Context, catDomain *Domain) error
}

type Repository interface{
	Fetch(ctx context.Context, email, password string) ([]Domain, error)
	GetByID(ctx context.Context, catID int) (Domain, error)
	Store(ctx context.Context, catDomain *Domain) error
}

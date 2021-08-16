package categories

import (
	"context"
	"time"
)

type Domain struct {
	Id        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type UseCase interface{
	GetById(ctx context.Context, catID string) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	GetByName(ctx context.Context, search string) ([]Domain, error)
	Store(ctx context.Context, catDomain *Domain) error
	Update(ctx context.Context, catDomain *Domain, id string) error
	Delete(ctx context.Context, id string) error
	// GetByActive(ctx context.Context, active bool) ([]Domain, error)
}

type Repository interface{
	Find(ctx context.Context, search string) ([]Domain, error)
	FindAll(ctx context.Context) ([]Domain, error)
	FindById(ctx context.Context, catID string) (Domain, error)
	Store(ctx context.Context, catDomain *Domain) error
	Update(ctx context.Context, catDomain *Domain, id string) error
	Delete(ctx context.Context, id string) error
}

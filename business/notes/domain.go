package notes

import (
	"context"
	"time"
)

type Domain struct {
	Id        		uint
	Title      		string
	Note      		string
	CategoryId      int
	CategoryName    string
	UserId      	int
	UserName    	string
	IsFree    		bool
	CreatedAt 		time.Time
	UpdatedAt 		time.Time
	DeletedAt 		time.Time
}

type UseCase interface {
	GetByName(ctx context.Context, search string) ([]Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	GetById(ctx context.Context, noteID string) (Domain, error)
	GetByCatId(ctx context.Context, catID string) ([]Domain, error)
	GetByUserId(ctx context.Context, userID string) ([]Domain, error)
	GetByIsFree(ctx context.Context, isFree string) ([]Domain, error)
	Store(ctx context.Context, catDomain *Domain) error
	// GetByActive(ctx context.Context, active bool) ([]Domain, error)
}

type Repository interface {
	Find(ctx context.Context, search string) ([]Domain, error)
	FindAll(ctx context.Context) ([]Domain, error)
	FindById(ctx context.Context, noteID string) (Domain, error)
	FindByCatId(ctx context.Context, catID string) ([]Domain, error)
	FindByUserId(ctx context.Context, userID string) ([]Domain, error)
	FindByIsFree(ctx context.Context, isFree string) ([]Domain, error)
	Store(ctx context.Context, catDomain *Domain) error
}

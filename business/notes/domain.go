package notes

import (
	"context"
	"notee/drivers/databases/categories"
	"notee/drivers/databases/users"
	"time"
)

type Domain struct {
	Id        		uint
	Title      		string
	Note      		string
	CategoryName    string
	UserName      	string
	CategoryId      int
	Category    	categories.Category
	UserId      	int
	User    		users.User
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
	Store(ctx context.Context, noteDomain *Domain) error
	Update(ctx context.Context, noteDomain *Domain, id string) error
	Delete(ctx context.Context, id string) error
	GetNote(ctx context.Context, id int) ([]Domain, error)
	// GetByActive(ctx context.Context, active bool) ([]Domain, error)
}

type Repository interface {
	Find(ctx context.Context, search string) ([]Domain, error)
	FindAll(ctx context.Context) ([]Domain, error)
	FindById(ctx context.Context, noteID string) (Domain, error)
	FindByCatId(ctx context.Context, catID string) ([]Domain, error)
	FindByUserId(ctx context.Context, userID string) ([]Domain, error)
	FindByIsFree(ctx context.Context, isFree string) ([]Domain, error)
	Store(ctx context.Context, noteDomain *Domain) error
	Update(ctx context.Context, noteDomain *Domain, id string) error
	Delete(ctx context.Context, id string) error
	GetNote(ctx context.Context, id int) ([]Domain, error)
}

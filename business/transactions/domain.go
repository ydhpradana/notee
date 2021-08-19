package transactions

import (
	"context"
	"notee/drivers/databases/notes"
	"notee/drivers/databases/users"
	"time"
)

type Domain struct {
	Id         uint
	Code      string
	NoteId int
	Note   notes.Note
	UserId     int
	User       users.User
	TotalPrice     int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
}

type UseCase interface {
	// GetByName(ctx context.Context, search string) ([]Domain, error)
	GetTrx(ctx context.Context, id int) ([]Domain, error)
	// GetById(ctx context.Context, noteID string) (Domain, error)
	// GetByCatId(ctx context.Context, catID string) ([]Domain, error)
	// GetByUserId(ctx context.Context, userID string) ([]Domain, error)
	// GetByIsFree(ctx context.Context, isFree string) ([]Domain, error)
	Store(ctx context.Context, transactionDomain *Domain) error
	// Update(ctx context.Context, noteDomain *Domain, id string) error
	// Delete(ctx context.Context, id string) error
	// GetByActive(ctx context.Context, active bool) ([]Domain, error)
}

type Repository interface {
	// Find(ctx context.Context, search string) ([]Domain, error)
	GetTrx(ctx context.Context, id int) ([]Domain, error)
	// FindById(ctx context.Context, noteID string) (Domain, error)
	// FindByCatId(ctx context.Context, catID string) ([]Domain, error)
	// FindByUserId(ctx context.Context, userID string) ([]Domain, error)
	// FindByIsFree(ctx context.Context, isFree string) ([]Domain, error)
	Store(ctx context.Context, transactionDomain *Domain) error
	// Update(ctx context.Context, noteDomain *Domain, id string) error
	// Delete(ctx context.Context, id string) error
}

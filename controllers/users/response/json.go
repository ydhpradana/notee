package response

import (
	"notee/business/users"
	"time"
)

type User struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	//Token     string    `json:"token"`
}

func FromDomain(domain users.Domain) User {
	return User{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		Name:      domain.Name,
		Email:     domain.Email,
		//Token:     domain.Token,
	}
}

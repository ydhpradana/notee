package request

import "notee/business/users"

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Users struct {
	Name     string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}


func (req *LoginUser) ToDomain() *users.Domain {
	return &users.Domain{
		Email:    req.Email,
		Password: req.Password,
	}
}

func (req *Users) ToDomain() *users.Domain {
	return &users.Domain{
		Name:     	req.Name,
		Email: 		req.Email,
		Password: 	req.Password,
	}
}

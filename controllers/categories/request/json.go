package request

import "notee/business/categories"

type Categories struct {
	Name      string `json:"name"`
}

func (req *Categories) ToDomain() *categories.Domain {
	return &categories.Domain{
		Name:      req.Name,
	}
}

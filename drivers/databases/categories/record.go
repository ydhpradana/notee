package categories

import (
	"notee/business/categories"
	"time"
)

type Category struct {
	Id        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `gorm:"index" json:"deletedAt"`
	Name      string    `json:"name"`
	
}

func (rec *Category) toDomain() categories.Domain {
	return categories.Domain{
		Id:        rec.Id,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
		DeletedAt: rec.DeletedAt,
		Name:      rec.Name,
		
	}
}

func fromDomain(category categories.Domain) *Category {
	return &Category{
		Id:        category.Id,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
		DeletedAt: category.DeletedAt,
		Name:      category.Name,
	}
}

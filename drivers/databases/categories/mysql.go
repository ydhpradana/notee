package categories

import (
	"context"
	"notee/business/categories"
	"strconv"

	"gorm.io/gorm"
)

type MysqlCategoryRepository struct {
	Conn *gorm.DB
}

func NewMysqlCategoryRepository(conn *gorm.DB) categories.Repository {
	return &MysqlCategoryRepository{
		Conn: conn,
	}
}


func (repository *MysqlCategoryRepository) FindById(ctx context.Context, id string) (categories.Domain, error) {
	rec := Category{}
	id_param, _ := strconv.Atoi(id)
	err := repository.Conn.Where("id = ?", id_param).First(&rec).Error
	if err != nil {
		return categories.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (repository *MysqlCategoryRepository) Find(ctx context.Context, search string) ([]categories.Domain, error) {
	rec := []Category{}
	err := repository.Conn.Where("name LIKE ?", "%"+search+"%").Find(&rec).Error
	if err != nil {
		return []categories.Domain{}, err
	}
	categoryDomain := []categories.Domain{}
	for _, value := range rec {
		categoryDomain = append(categoryDomain, value.toDomain())
	}
	return categoryDomain, nil
}

func (repository *MysqlCategoryRepository) FindAll(ctx context.Context) ([]categories.Domain, error) {
	rec := []Category{}
	err := repository.Conn.Find(&rec).Error
	if err != nil {
		return []categories.Domain{}, err
	}
	categoryDomain := []categories.Domain{}
	for _, value := range rec {
		categoryDomain = append(categoryDomain, value.toDomain())
	}
	return categoryDomain, nil
}

func (repository *MysqlCategoryRepository) Store(ctx context.Context, categoryDomain *categories.Domain) error {
	rec := fromDomain(*categoryDomain)

	result := repository.Conn.Create(rec)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository *MysqlCategoryRepository) Update(ctx context.Context, categoryDomain *categories.Domain, id string) error {
	rec := fromDomain(*categoryDomain)
	id_param, _ := strconv.Atoi(id)
	result := repository.Conn.Where("id = ?", id_param).Updates(rec)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository *MysqlCategoryRepository) Delete(ctx context.Context, id string) error {
	rec := Category{}
	id_param, _ := strconv.Atoi(id)
	result := repository.Conn.Where("id = ?", id_param).Delete(rec)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

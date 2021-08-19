package ratings

import (
	"context"
	"notee/business/ratings"
	"strconv"

	"gorm.io/gorm"
)

type MysqlRatingRepository struct {
	Conn *gorm.DB
}

func NewMysqlRatingRepository(conn *gorm.DB) ratings.Repository {
	return &MysqlRatingRepository{
		Conn: conn,
	}
}


func (repository *MysqlRatingRepository) GetById(ctx context.Context, id int) ([]ratings.Domain, error) {
	rec := []Rating{}
	err := repository.Conn.Preload("Users").Where("ratings.note_id = ?", id).Find(&rec).Error
	if err != nil {
		return []ratings.Domain{}, err
	}

	ratingDomain := []ratings.Domain{}
	for _, value := range rec {
		ratingDomain = append(ratingDomain, value.toDomain())
	}

	return ratingDomain, nil
}

func (repository *MysqlRatingRepository) Store(ctx context.Context, ratingDomain *ratings.Domain) error {
	rec := fromDomain(*ratingDomain)

	result := repository.Conn.Create(rec)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository *MysqlRatingRepository) Update(ctx context.Context, ratingDomain *ratings.Domain, id string) error {
	rec := fromDomain(*ratingDomain)
	id_param, _ := strconv.Atoi(id)
	result := repository.Conn.Where("id = ?", id_param).Updates(rec)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository *MysqlRatingRepository) Delete(ctx context.Context, id string) error {
	rec := Rating{}
	id_param, _ := strconv.Atoi(id)
	result := repository.Conn.Where("id = ?", id_param).Delete(rec)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

package notes

import (
	"context"
	"notee/business/notes"
	"strconv"

	"gorm.io/gorm"
)

type MysqlNoteRepository struct {
	Conn *gorm.DB
}

func NewMysqlNotRepository(conn *gorm.DB) notes.Repository {
	return &MysqlNoteRepository{
		Conn: conn,
	}
}

func (repository *MysqlNoteRepository) FindById(ctx context.Context, id string) (notes.Domain, error) {
	rec := Note{}
	id_param, _ := strconv.Atoi(id)
	err := repository.Conn.Where("id = ?", id_param).First(&rec).Error
	if err != nil {
		return notes.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (repository *MysqlNoteRepository) Find(ctx context.Context, search string) ([]notes.Domain, error) {
	rec := []Note{}
	err := repository.Conn.Where("title LIKE ?", "%"+search+"%").Or("note LIKE ?", "%"+search+"%").Find(&rec).Error
	if err != nil {
		return []notes.Domain{}, err
	}
	noteDomain := []notes.Domain{}
	for _, value := range rec {
		noteDomain = append(noteDomain, value.toDomain())
	}
	return noteDomain, nil
}

func (repository *MysqlNoteRepository) FindAll(ctx context.Context) ([]notes.Domain, error) {
	rec := []Note{}
	err := repository.Conn.Find(&rec).Error
	if err != nil {
		return []notes.Domain{}, err
	}
	noteDomain := []notes.Domain{}
	for _, value := range rec {
		noteDomain = append(noteDomain, value.toDomain())
	}
	return noteDomain, nil
}

func (repository *MysqlNoteRepository) Store(ctx context.Context, noteDomain *notes.Domain) error {
	rec := fromDomain(*noteDomain)

	result := repository.Conn.Create(rec)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

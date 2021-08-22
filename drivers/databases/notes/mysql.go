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

func NewMysqlNoteRepository(conn *gorm.DB) notes.Repository {
	return &MysqlNoteRepository{
		Conn: conn,
	}
}

func (repository *MysqlNoteRepository) FindById(ctx context.Context, id string) (notes.Domain, error) {
	rec := Note{}
	id_param, _ := strconv.Atoi(id)
	err := repository.Conn.Preload("Categories").Preload("Users").Where("id = ?", id_param).First(&rec).Error
	if err != nil {
		return notes.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (repository *MysqlNoteRepository) Find(ctx context.Context, search string) ([]notes.Domain, error) {
	rec := []Note{}
	err := repository.Conn.Preload("Categories").Preload("Users").Where("title LIKE ?", "%"+search+"%").Or("note LIKE ?", "%"+search+"%").Find(&rec).Error
	if err != nil {
		return []notes.Domain{}, err
	}
	noteDomain := []notes.Domain{}
	for _, value := range rec {
		noteDomain = append(noteDomain, value.toDomain())
	}
	return noteDomain, nil
}

func (repository *MysqlNoteRepository) FindByCatId(ctx context.Context, catId string) ([]notes.Domain, error) {
	rec := []Note{}
	id_cat, _ := strconv.Atoi(catId)
	err := repository.Conn.Preload("Categories").Preload("Users").Where("category_id = ?", id_cat).Find(&rec).Error
	if err != nil {
		return []notes.Domain{}, err
	}
	noteDomain := []notes.Domain{}
	for _, value := range rec {
		noteDomain = append(noteDomain, value.toDomain())
	}
	return noteDomain, nil
}

func (repository *MysqlNoteRepository) FindByUserId(ctx context.Context, userId string) ([]notes.Domain, error) {
	rec := []Note{}
	id_user, _ := strconv.Atoi(userId)
	err := repository.Conn.Preload("Categories").Preload("Users").Where("user_id = ?", id_user).Find(&rec).Error
	if err != nil {
		return []notes.Domain{}, err
	}
	noteDomain := []notes.Domain{}
	for _, value := range rec {
		noteDomain = append(noteDomain, value.toDomain())
	}
	return noteDomain, nil
}

func (repository *MysqlNoteRepository) FindByIsFree(ctx context.Context, free string) ([]notes.Domain, error) {
	rec := []Note{}
	is_free, _ := strconv.Atoi(free)
	err := repository.Conn.Preload("Categories").Preload("Users").Where("is_free = ?", is_free).Find(&rec).Error
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
	// query := `SELECT notes.*, users.*, categories.* , AVG(ratings.rating) "rating" FROM notes 
	// 			join ratings on ratings.note_id =notes.id
	// 			join users on users.id =notes.user_id
	// 			join categories on categories.id =notes.category_id`
	err := repository.Conn.Preload("Categories").Preload("Users").Find(&rec).Error
	//err := repository.Conn.Raw(query).Find(&rec).Error
	//fmt.Printf("%+v", rec)
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

func (repository *MysqlNoteRepository) Update(ctx context.Context, noteDomain *notes.Domain, id string) error {
	rec := fromDomain(*noteDomain)
	id_param, _ := strconv.Atoi(id)
	result := repository.Conn.Where("id = ?", id_param).Updates(rec)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository *MysqlNoteRepository) Delete(ctx context.Context, id string) error {
	rec := Note{}
	id_param, _ := strconv.Atoi(id)
	result := repository.Conn.Where("id = ?", id_param).Delete(rec)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository *MysqlNoteRepository) GetNote(ctx context.Context, userID int) ([]notes.Domain, error) {
	rec := []Note{}

	err := repository.Conn.Preload("Categories").Preload("Users").Where("notes.user_id = ?", userID).Find(&rec).Error
	// query := `SELECT notes.*, users.*, categories.* , AVG(ratings.rating) "rating" FROM notes 
	// 			join ratings on ratings.note_id =notes.id
	// 			join users on users.id =notes.user_id
	// 			join categories on categories.id =notes.category_id`
	// err := repository.Conn.Raw(query).Where("notes.user_id = ?", userID).Find(&rec).Error
	//err := repository.Conn.Preload(clause.Associations).Where("transactions.user_id = ?", userID).Find(&rec).Error
	if err != nil {
		return []notes.Domain{}, err
	}

	noteDomain := []notes.Domain{}
	for _, value := range rec {
		noteDomain = append(noteDomain, value.toDomain())
	}

	return noteDomain, nil
}
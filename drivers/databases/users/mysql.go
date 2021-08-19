package users

import (
	"context"
	"notee/business"
	"notee/business/users"
	"notee/helpers/hash"

	"gorm.io/gorm"
)

type MysqlUserRepository struct {
	Conn *gorm.DB
}

func NewMysqlUserRepository(conn *gorm.DB) users.Repository {
	return &MysqlUserRepository{
		Conn: conn,
	}
}

func (repository *MysqlUserRepository) Login(ctx context.Context, email, password string) (users.Domain , error) {
	var userLogin = User{}
	result := repository.Conn.Where("email = ? ", email).First(&userLogin)
	if result.Error != nil {
		return users.Domain{}, result.Error
	}
	match := hash.ValidateHash(password, userLogin.Password)
	if !match {
		return users.Domain{}, business.ErrInternalServer
	}

	return userLogin.toDomain(), nil
}

func (repository *MysqlUserRepository) GetByEmail(ctx context.Context, email string) (users.Domain, error) {
	rec := User{}
	err := repository.Conn.Where("email = ?", email).First(&rec).Error
	if err != nil {
		return users.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (repository *MysqlUserRepository) GetById(ctx context.Context, id int) (users.Domain, error) {
	rec := User{}
	err := repository.Conn.Where("id = ?", id ).First(&rec).Error
	if err != nil {
		return users.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (repository *MysqlUserRepository) Store(ctx context.Context, userDomain *users.Domain) error {
	rec := fromDomain(*userDomain)

	result := repository.Conn.Create(rec)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

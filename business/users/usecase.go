package users

import (
	"context"
	"notee/app/middleware"
	"notee/business"
	"notee/helpers/hash"
	"strings"
	"time"
)

type UserUseCase struct {
	Repository Repository
	ContextTimeout time.Duration 
	jwtAuth * middleware.ConfigJWT

}


func NewUserUsecase(repository Repository, jwtauth *middleware.ConfigJWT, timeout time.Duration) UseCase {
	return &UserUseCase{
		Repository : repository,
		jwtAuth:        jwtauth,
		ContextTimeout : timeout,
	}
}

func (usecase *UserUseCase) Login(ctx context.Context, email, password string) (Domain, error){
	
	result, err := usecase.Repository.Login(ctx, email, password)
	if err != nil {
		return Domain{}, err
	}
	//hash_password, err := hash.Hash(password)
	return result, nil
}

func (uc *UserUseCase) Store(ctx context.Context, userDomain *Domain) error {
	ctx, cancel := context.WithTimeout(ctx, uc.ContextTimeout)
	defer cancel()
	
	existedUser, err := uc.Repository.GetByEmail(ctx, userDomain.Email)
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return err
		}
	}
	if existedUser != (Domain{}) {
		return business.ErrDuplicateData
	}

	userDomain.Password, err = hash.Hash(userDomain.Password)
	if err != nil {
		return business.ErrInternalServer
	}
	err = uc.Repository.Store(ctx, userDomain)
	if err != nil {
		return err
	}

	return nil
}

func (uc *UserUseCase) CreateToken(ctx context.Context, email, password string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.ContextTimeout)
	defer cancel()

	if strings.TrimSpace(email) == "" && strings.TrimSpace(password) == "" {
		return "", business.ErrUsernamePasswordNotFound
	}

	userDomain, err := uc.Repository.GetByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	if !hash.ValidateHash(password, userDomain.Password) {
		return "", business.ErrInternalServer
	}

	token := uc.jwtAuth.GenerateToken(int(userDomain.ID))
	return token, nil
}

func (usecase *UserUseCase) GetById(ctx context.Context, id int) (Domain, error) {
	result, err := usecase.Repository.GetById(ctx, id)
		if err != nil {
		return Domain{}, err
	}
	return result, nil
}


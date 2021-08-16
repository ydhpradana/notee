package users

import (
	"context"
	"notee/app/middleware"
	"notee/business"
	"notee/helpers/hash"
	"notee/helpers/middlewares"
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
	match := hash.ValidateHash(password, result.Password)
	if !match {
		return Domain{}, business.ErrInternalServer
	}
	result.Token, _ = middlewares.GenerateTokenJWT(int32(result.Id))
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

func (uc *UserUseCase) CreateToken(ctx context.Context, username, password string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.ContextTimeout)
	defer cancel()

	if strings.TrimSpace(username) == "" && strings.TrimSpace(password) == "" {
		return "", business.ErrUsernamePasswordNotFound
	}

	userDomain, err := uc.Repository.GetByEmail(ctx, username)
	if err != nil {
		return "", err
	}

	if !hash.ValidateHash(password, userDomain.Password) {
		return "", business.ErrInternalServer
	}

	token := uc.jwtAuth.GenerateToken(int(userDomain.Id))
	return token, nil
}

func (usecase *UserUseCase) GetById(ctx context.Context, id string) (Domain, error) {
	result, err := usecase.Repository.GetById(ctx, id)
		if err != nil {
		return Domain{}, err
	}

	
	return result, nil
}
package users_test

import (
	"context"
	"errors"
	"notee/app/middleware"
	business "notee/business"
	"notee/business/users"
	userMock "notee/business/users/mocks"
	"notee/helpers/hash"

	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	userRepository userMock.Repository
	userUsecase    users.UseCase
	jwtAuth        *middleware.ConfigJWT
)

func setup() {
	jwtAuth = &middleware.ConfigJWT{SecretJWT: "secret", ExpiresDuration: 2}
	userUsecase = users.NewUserUsecase(&userRepository, jwtAuth, 2)

}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())

}

func TestGetById(t *testing.T) {
	t.Run("test case 1, valid test ", func(t *testing.T) {
		domain := users.Domain{
			ID:       1,
			Name:     "Yudha Pradana",
			Email:    "yudhapradana@gmail.com",
		}
		userRepository.On("GetById", mock.Anything, mock.AnythingOfType("int")).Return(domain, nil).Once()

		result, err := userUsecase.GetById(context.Background(), 1)

		assert.Equal(t, err, nil)
		assert.Equal(t, result, domain)
	})

	// t.Run("test case 2, invalid id", func(t *testing.T) {
	// 	userRepository.On("GetById", mock.Anything, mock.AnythingOfType("int")).Return(users.Domain{}, business.ErrNotFound).Once()
	// 	result, err := userUsecase.GetById(context.Background(), -1)

	// 	assert.Equal(t, result, users.Domain{})
	// 	assert.Equal(t, err, business.ErrNotFound)
	// })

}

func TestRegister(t *testing.T) {
	t.Run("test case 1, valid test ", func(t *testing.T) {
		domain := users.Domain{
			ID:       1,
			Name:     "Yudha Pradana",
			Email:    "yudhapradana@gmail.com",
		}
		userRepository.On("GetByEmail", mock.Anything, mock.AnythingOfType("string")).Return(users.Domain{}, nil).Once()
		userRepository.On("Store", mock.Anything, mock.AnythingOfType("*users.Domain")).Return(nil).Once()

		err := userUsecase.Store(context.Background(), &domain)

		assert.NoError(t, err)
	})

	t.Run("test case 2, duplicate data", func(t *testing.T) {
		domain := users.Domain{
			ID:       1,
			Name:     "Yudha Pradana",
			Email:    "yudhapradana@gmail.com",
		}
		errRepository := errors.New("duplicate data")
		userRepository.On("GetByEmail", mock.Anything, mock.AnythingOfType("string")).Return(domain, errRepository).Once()

		err := userUsecase.Store(context.Background(), &domain)

		assert.Equal(t, err, errRepository)
	})

	t.Run("test case 3, duplicate data", func(t *testing.T) {
		domain := users.Domain{
			ID:       1,
			Name:     "Yudha Pradana",
			Email:    "yudhapradana@gmail.com",
		}
		userRepository.On("GetByEmail", mock.Anything, mock.AnythingOfType("string")).Return(domain, business.ErrDuplicateData).Once()

		err := userUsecase.Store(context.Background(), &domain)

		assert.Equal(t, err, business.ErrDuplicateData)
	})

	t.Run("test case 4, duplicate data", func(t *testing.T) {

		userRepository.On("GetByEmail", mock.Anything, mock.AnythingOfType("string")).Return(users.Domain{}, business.ErrInternalServer).Once()

		err := userUsecase.Store(context.Background(), &users.Domain{})

		assert.Equal(t, err, business.ErrInternalServer)
	})

	t.Run("test case 5, register failed", func(t *testing.T) {
		domain := users.Domain{
			ID:       1,
			Name:     "Yudha Pradana",
			Email:    "yudhapradana@gmail.com",
		}
		errRepository := errors.New("register failed")
		userRepository.On("GetByEmail", mock.Anything, mock.AnythingOfType("string")).Return(users.Domain{}, nil).Once()
		userRepository.On("Store", mock.Anything, mock.AnythingOfType("*users.Domain")).Return(errRepository).Once()

		err := userUsecase.Store(context.Background(), &domain)

		assert.Equal(t, err, errRepository)
	})

}

func TestLogin(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		pass, _ := hash.Hash("123456")
		usersDomain := users.Domain{
			ID:       1,
			Name:     "Yudha Pradana",
			Email:    "yudhapradana@gmail.com",
			Password: pass,
		}

		userRepository.On("GetByEmail", mock.Anything, mock.AnythingOfType("string")).Return(usersDomain, nil).Once()

		_, err := userUsecase.CreateToken(context.Background(), "yudha", "123456")
		assert.Nil(t, err)
	})
	t.Run("test case 2, password error", func(t *testing.T) {
		pass, _ := hash.Hash("123456")
		usersDomain := users.Domain{
			ID:       1,
			Name:     "Yudha Pradana",
			Email:    "yudhapradana@gmail.com",
			Password: pass,
		}

		userRepository.On("GetByEmail", mock.Anything, mock.AnythingOfType("string")).Return(usersDomain, nil).Once()

		_, err := userUsecase.CreateToken(context.Background(), "yudha", "1234567")
		assert.Equal(t, err, business.ErrInternalServer)

	})

	t.Run("test case 3, error record", func(t *testing.T) {

		errRepository := errors.New("error record")
		userRepository.On("GetByEmail", mock.Anything, mock.AnythingOfType("string")).Return(users.Domain{}, errRepository).Once()

		result, err := userUsecase.CreateToken(context.Background(), "yudh", "123456")

		assert.Equal(t, err, errRepository)
		assert.Equal(t, result, result)
	})

}

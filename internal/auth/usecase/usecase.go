package usecase

import (
	"context"
	"errors"

	"github.com/BIC-Final-Project/backend/internal/auth/entity"
	"github.com/BIC-Final-Project/backend/internal/auth/http/middlewares"
	"github.com/BIC-Final-Project/backend/internal/auth/repository"
)

type AuthUsecase interface {
	CreateUser(ctx context.Context, user entity.User) (*entity.User, error)
	Login(ctx context.Context, user entity.User) (*entity.UserResponse, *entity.Token, error)
	FindUserByEmail(ctx context.Context, email string) (*entity.User, error)
}

type authUsecase struct {
	repo repository.AuthRepository
}

// CreateUser implements AuthUsecase.
func (a *authUsecase) CreateUser(ctx context.Context, user entity.User) (*entity.User, error) {
	checkUsername, _ := a.repo.FindUserByEmail(ctx, user.Email)
	if checkUsername != nil {
		return nil, errors.New("username already exists")
	}

	hashedPass, errHashed := middlewares.HashPassword(user.Password)
	if errHashed != nil {
		return nil, errHashed
	}

	user.Password = hashedPass

	data, err := a.repo.SaveUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// FindUserByUsername implements AuthUsecase.
func (a *authUsecase) FindUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	userData, err := a.repo.FindUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return userData, nil
}

// Login implements AuthUsecase.
func (a *authUsecase) Login(ctx context.Context, user entity.User) (*entity.UserResponse, *entity.Token, error) {
	userData, err := a.repo.FindUserByEmail(ctx, user.Email)
	if err != nil {
		return nil, nil, err
	}

	errBcrypt := middlewares.CheckPassword(user.Password, userData.Password)
	if errBcrypt != nil {
		return nil, nil, errors.New("invalid password")
	}

	Token, err := middlewares.SignJWT(*userData, 3)
	if err != nil {
		return nil, nil, err
	}

	dataToken := &entity.Token{
		Token: Token,
	}

	userRes := &entity.UserResponse{
		ID:          userData.ID,
		NamaLengkap: userData.NamaLengkap,
		NoHandphone: userData.NoHandphone,
		Email:       userData.Email,
		Role:        userData.Role,
	}

	return userRes, dataToken, nil
}

func NewAuthUsecase(repo repository.AuthRepository) AuthUsecase {
	return &authUsecase{repo}
}

//go:build wireinject
// +build wireinject

package di

import (
	"github.com/BIC-Final-Project/backend/internal/auth/http/handler"
	"github.com/BIC-Final-Project/backend/internal/auth/repository"
	"github.com/BIC-Final-Project/backend/internal/auth/usecase"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitAdmin(db *mongo.Database) *handler.AuthHandler {
	wire.Build(
		repository.NewAuthRepository,
		usecase.NewAuthUsecase,
		handler.NewAuthHandler,
	)

	return &handler.AuthHandler{}
}

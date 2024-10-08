// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/BIC-Final-Project/backend/internal/auth/http/handler"
	"github.com/BIC-Final-Project/backend/internal/auth/repository"
	"github.com/BIC-Final-Project/backend/internal/auth/usecase"
	"go.mongodb.org/mongo-driver/mongo"
)

// Injectors from wire.go:

func InitAdmin(db *mongo.Database) *handler.AuthHandler {
	authRepository := repository.NewAuthRepository(db)
	authUsecase := usecase.NewAuthUsecase(authRepository)
	authHandler := handler.NewAuthHandler(authUsecase)
	return authHandler
}

//go:build wireinject
// +build wireinject

package di

import (
	"github.com/BIC-Final-Project/backend/configs/env"
	"github.com/BIC-Final-Project/backend/internal/operational/http/handler"
	"github.com/BIC-Final-Project/backend/internal/operational/repository"
	"github.com/BIC-Final-Project/backend/internal/operational/usecase"
	"github.com/BIC-Final-Project/backend/internal/storage"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitFasilitas(db *mongo.Database, env env.EnvVars) *handler.FasilitasHandler {
	wire.Build(
		repository.NewFasilitasRepository,
		usecase.NewFasilitasUsecase,
		handler.NewFasilitasHandler,
		storage.NewS3Service,
	)

	return &handler.FasilitasHandler{}
}

func InitMembershipType(db *mongo.Database) *handler.MembershipTypeHandler {
	wire.Build(
		repository.NewMembershipTypeRepository,
		usecase.NewMembershipTypeUsecase,
		handler.NewMembershipTypeHandler,
	)

	return &handler.MembershipTypeHandler{}
}

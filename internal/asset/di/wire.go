//go:build wireinject
// +build wireinject

package di

import (
	"github.com/BIC-Final-Project/backend/configs/env"
	"github.com/BIC-Final-Project/backend/internal/asset/http/handler"
	"github.com/BIC-Final-Project/backend/internal/asset/repository"
	"github.com/BIC-Final-Project/backend/internal/asset/usecase"
	"github.com/BIC-Final-Project/backend/internal/storage"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitAset(db *mongo.Database, env env.EnvVars) *handler.AsetHandler {
	wire.Build(
		repository.NewAsetRepository,
		usecase.NewAsetUsecase,
		handler.NewAsetHandler,
		storage.NewS3Service,
	)

	return &handler.AsetHandler{}
}

func InitVendor(db *mongo.Database, env env.EnvVars) *handler.VendorHandler {
	wire.Build(
		repository.NewVendorRepository,
		usecase.NewVendorUsecase,
		handler.NewVendorHandler,
	)

	return &handler.VendorHandler{}
}

func InitPerencanaan(db *mongo.Database, env env.EnvVars) *handler.PerencanaanHandler {
	wire.Build(
		repository.NewPerencanaanRepository,
		usecase.NewPerencanaanUsecase,
		handler.NewPerencanaanHandler,
	)

	return &handler.PerencanaanHandler{}
}

func InitPemeliharaan(db *mongo.Database, env env.EnvVars) *handler.PemeliharaanHandler {
	wire.Build(
		repository.NewPemeliharaanRepository,
		usecase.NewPemeliharaanUsecase,
		handler.NewPemeliharaanHandler,
	)

	return &handler.PemeliharaanHandler{}
}

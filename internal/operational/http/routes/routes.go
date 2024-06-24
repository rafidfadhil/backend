package routes

import (
	"github.com/BIC-Final-Project/backend/configs/env"
	"github.com/BIC-Final-Project/backend/internal/operational/di"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupOperationalRoutes(app *fiber.App, db *mongo.Database, env env.EnvVars) {
	operational := app.Group("api/v1/operational")

	// FASILITAS
	fasilitasHandler := di.InitFasilitas(db, env)
	fasilitas := operational.Group("/fasilitas")
	fasilitas.Post("/", fasilitasHandler.CreateFasilitas)
	fasilitas.Get("/", fasilitasHandler.GetAllFasilitas)
	fasilitas.Get("/name", fasilitasHandler.GetAllFasilitasName)
	fasilitas.Get("/:id", fasilitasHandler.GetFasilitas)
	fasilitas.Put("/:id", fasilitasHandler.UpdateFasilitas)
	fasilitas.Delete("/:id", fasilitasHandler.DeleteFasilitas)

	// MEMBERSHIP TYPE
	membershipTypeHandler := di.InitMembershipType(db)
	membershipType := operational.Group("/membership-type")
	membershipType.Post("/", membershipTypeHandler.CreateMembershipType)
	membershipType.Get("/", membershipTypeHandler.GetAllMembershipType)
	membershipType.Get("/:id", membershipTypeHandler.GetMembershipType)
	membershipType.Put("/:id", membershipTypeHandler.UpdateMembershipType)
	membershipType.Delete("/:id", membershipTypeHandler.DeleteMembershipType)
}

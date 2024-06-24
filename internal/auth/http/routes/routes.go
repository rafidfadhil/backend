package routes

import (
	"github.com/BIC-Final-Project/backend/internal/auth/di"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupAuthRoutes(app *fiber.App, db *mongo.Database) {
	auth := di.InitAdmin(db)

	a := app.Group("api/v1/auth")

	a.Post("/register", auth.Register)
	a.Post("/login", auth.Login)
}

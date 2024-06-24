package main

import (
	"log"

	"github.com/BIC-Final-Project/backend/configs/env"
	assetRoute "github.com/BIC-Final-Project/backend/internal/asset/http/routes"
	authRoute "github.com/BIC-Final-Project/backend/internal/auth/http/routes"
	"github.com/BIC-Final-Project/backend/internal/mongodb"
	opRoute "github.com/BIC-Final-Project/backend/internal/operational/http/routes"
	"github.com/BIC-Final-Project/backend/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	env, err := env.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	db, err := mongodb.ConnectDB(env.MONGODBURI, env.MONGODBNAME)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	defer func(db *mongo.Database) {
		err := mongodb.CloseMongoDB(db)
		if err != nil {
			log.Fatalf("Error closing MongoDB: %s", err)
		}
	}(db)

	app := setupApp()

	authRoute.SetupAuthRoutes(app, db)
	opRoute.SetupOperationalRoutes(app, db, env)
	assetRoute.SetupAssetRoutes(app, db, env)

	err = app.Listen(":" + env.PORT)
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
	}

	log.Printf("Server is running on port %s", env.PORT)
}

func setupApp() *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: utils.ErrorHandler,
	})

	app.Use(cors.New())
	app.Use(logger.New())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("OK!")
	})

	return app
}

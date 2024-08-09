package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	userRoutes "github.com/mauFade/infinity/internal/handlers/user"
)

func SetupRoutes(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins: os.Getenv("APP_URL"),
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	v1 := app.Group("/v1")

	users := v1.Group("/users")

	userRoutes.UseUserRoutes(users)
}

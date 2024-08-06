package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mauFade/infinity/internal/handlers"
)

func SetupRoutes(app *fiber.App) {
	v1 := app.Group("/v1")

	users := v1.Group("/users")

	users.Post("/", handlers.CreateUserHandler)
}

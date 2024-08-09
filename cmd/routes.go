package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	userHandlers "github.com/mauFade/infinity/internal/handlers/user"
	"github.com/mauFade/infinity/internal/middleware"
)

func SetupRoutes(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins: os.Getenv("APP_URL"),
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	v1 := app.Group("/v1")

	{
		users := v1.Group("/users")

		users.Post("/", userHandlers.CreateUserHandler)

		users.Post("/auth", userHandlers.AuthenticateHandler)

		users.Use(middleware.EnsureAuthenticated())

		users.Get("/", userHandlers.GetUserProfileHandler)

		users.Post("/address", userHandlers.CreateAddressHandler)
	}

}

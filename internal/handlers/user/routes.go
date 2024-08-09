package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mauFade/infinity/internal/middleware"
)

func UseUserRoutes(router fiber.Router) {
	router.Post("/", CreateUserHandler)

	router.Post("/auth", AuthenticateHandler)

	router.Use(middleware.EnsureAuthenticated())

	router.Get("/", GetUserProfileHandler)

	router.Post("/address", CreateAddressHandler)
}

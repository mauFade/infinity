package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mauFade/infinity/internal/config"
	"github.com/mauFade/infinity/internal/repositories"
	"github.com/mauFade/infinity/internal/usecase/user"
)

func GetUserProfileHandler(c *fiber.Ctx) error {
	id, err := config.GetIdToken(c.Get("Authorization"))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"status":  fiber.StatusInternalServerError,
		})
	}

	uc := user.NewGetUserProfileUseCase(
		repositories.NewUserRepository(config.Database.DataBase),
		repositories.NewAddressRepository(config.Database.DataBase),
	)

	data, err := uc.Execute(&user.GetUserProfileInput{
		UserId: id,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"status":  fiber.StatusInternalServerError,
		})
	}

	return c.Status(fiber.StatusOK).JSON(data)
}

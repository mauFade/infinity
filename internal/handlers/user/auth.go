package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mauFade/infinity/internal/config"
	"github.com/mauFade/infinity/internal/repositories"
	"github.com/mauFade/infinity/internal/usecase/user"
)

type authUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func AuthenticateHandler(c *fiber.Ctx) error {
	payload := authUserRequest{}

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"status":  fiber.StatusInternalServerError,
		})
	}

	uc := user.NewAuthUserUseCase(repositories.NewUserRepository(config.Database.DataBase))

	data, err := uc.Execute(&user.AuthInput{
		Email:    payload.Email,
		Password: payload.Password,
	})

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"status":  fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(data)
}

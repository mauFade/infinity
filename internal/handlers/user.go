package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mauFade/infinity/internal/config"
	"github.com/mauFade/infinity/internal/repositories"
	"github.com/mauFade/infinity/internal/usecase/user"
)

func CreateUserHandler(c *fiber.Ctx) error {
	payload := struct {
		Name       string `json:"name"`
		Email      string `json:"email"`
		Phone      string `json:"phone"`
		Password   string `json:"password"`
		Profession string `json:"profession"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	uc := user.NewCreateUserUseCase(repositories.NewUserRepository(config.Database.DataBase))

	uc.Execute(&user.CreateUserInput{
		Name:       payload.Name,
		Email:      payload.Email,
		Phone:      payload.Phone,
		Password:   payload.Password,
		Profession: payload.Profession,
	})

	return c.JSON(payload)
}

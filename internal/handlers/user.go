package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mauFade/infinity/internal/config"
	"github.com/mauFade/infinity/internal/repositories"
	"github.com/mauFade/infinity/internal/usecase/user"
)

type Request struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Password   string `json:"password"`
	Profession string `json:"profession"`
}

func CreateUserHandler(c *fiber.Ctx) error {
	payload := Request{}

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	uc := user.NewCreateUserUseCase(repositories.NewUserRepository(config.Database.DataBase))

	output, err := uc.Execute(&user.CreateUserInput{
		Name:       payload.Name,
		Email:      payload.Email,
		Phone:      payload.Phone,
		Password:   payload.Password,
		Profession: payload.Profession,
	})

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"status":  fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(output)
}

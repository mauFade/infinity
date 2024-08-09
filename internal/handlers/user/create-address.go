package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mauFade/infinity/internal/config"
	"github.com/mauFade/infinity/internal/repositories"
	"github.com/mauFade/infinity/internal/usecase/user"
)

type createAddressInput struct {
	Country       string `json:"country"`
	ZipCode       string `json:"zip_code"`
	Street        string `json:"street"`
	Number        string `json:"number"`
	Neighbourhood string `json:"neighbourhood"`
	City          string `json:"city"`
	Estate        string `json:"estate"`
}

func CreateAddressHandler(c *fiber.Ctx) error {
	id, err := config.GetIdToken(c.Get("Authorization"))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"status":  fiber.StatusInternalServerError,
		})
	}

	payload := createAddressInput{}

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"status":  fiber.StatusInternalServerError,
		})
	}

	uc := user.NewCreateAddressUseCase(
		repositories.NewAddressRepository(config.Database.DataBase),
		repositories.NewUserRepository(config.Database.DataBase),
	)

	data, err := uc.Execute(
		&user.CreateAddressInput{
			UserID:        id,
			Country:       payload.Country,
			ZipCode:       payload.ZipCode,
			Street:        payload.Street,
			Number:        payload.Number,
			Neighbourhood: payload.Neighbourhood,
			City:          payload.City,
			Estate:        payload.Estate,
		},
	)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"status":  fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(data)
}

package user

import (
	"github.com/google/uuid"
	"github.com/mauFade/infinity/internal/repositories"
)

type CreateAddressUseCase struct {
	addrRepository *repositories.AddressRepository
	userRepository *repositories.UserRepository
}

type CreateAddressInput struct {
	UserID        string
	Country       string
	ZipCode       string
	Street        string
	Number        string
	Neighbourhood string
	City          string
	Estate        string
}

type CreateAddressOutput struct {
	ID            uuid.UUID `json:"id"`
	UserID        uuid.UUID `json:"user_id"`
	Country       string    `json:"country"`
	ZipCode       string    `json:"zip_code"`
	Street        string    `json:"street"`
	Number        string    `json:"number"`
	Neighbourhood string    `json:"neighbourhood"`
	City          string    `json:"city"`
	Estate        string    `json:"estate"`
}

func NewCreateAddressUseCase(r *repositories.AddressRepository, u *repositories.UserRepository) *CreateAddressUseCase {
	return &CreateAddressUseCase{
		addrRepository: r,
		userRepository: u,
	}
}

func (u *CreateAddressUseCase) Execute() (*CreateAddressOutput, error) {

	return &CreateAddressOutput{}, nil
}

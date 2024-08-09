package user

import (
	"errors"

	"github.com/google/uuid"
	"github.com/mauFade/infinity/internal/models"
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

func (u *CreateAddressUseCase) Execute(data *CreateAddressInput) (*CreateAddressOutput, error) {
	user := u.userRepository.FindByID(data.UserID)

	if user == nil {
		return nil, errors.New("user not found with this id")
	}

	addr := u.addrRepository.FindByUserID(data.UserID)

	if addr != nil {
		return nil, errors.New("user already has an address")
	}

	addr = models.NewAddress(
		uuid.New(),
		user.ID,
		data.Country,
		data.ZipCode,
		data.Street,
		data.Number,
		data.Neighbourhood,
		data.City,
		data.Estate,
	)

	u.addrRepository.Create(*addr)

	return &CreateAddressOutput{
		ID:            addr.ID,
		UserID:        addr.UserID,
		Country:       addr.Country,
		ZipCode:       addr.ZipCode,
		Street:        addr.Street,
		Number:        addr.Number,
		Neighbourhood: addr.Neighbourhood,
		City:          addr.City,
		Estate:        addr.Estate,
	}, nil
}

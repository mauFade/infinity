package user

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/mauFade/infinity/internal/repositories"
)

type GetUserProfileUseCase struct {
	userRepository *repositories.UserRepository
	addrRepository *repositories.AddressRepository
}

func NewGetUserProfileUseCase(r *repositories.UserRepository, a *repositories.AddressRepository) *GetUserProfileUseCase {
	return &GetUserProfileUseCase{
		userRepository: r,
		addrRepository: a,
	}
}

type GetUserProfileInput struct {
	UserId string
}

type AddressOutput struct {
	ID            uuid.UUID `json:"id"`
	Country       string    `json:"country"`
	ZipCode       string    `json:"zip_code"`
	Street        string    `json:"street"`
	Number        string    `json:"number"`
	Neighbourhood string    `json:"neighbourhood"`
	City          string    `json:"city"`
	Estate        string    `json:"estate"`
}

type GetUserProfileOutput struct {
	ID         uuid.UUID      `json:"id"`
	Name       string         `json:"name"`
	Email      string         `json:"email"`
	Phone      string         `json:"phone"`
	Agency     string         `json:"agency"`
	Bank       string         `json:"bank"`
	Serial     string         `json:"serial"`
	Profession string         `json:"profession"`
	UpdatedAt  time.Time      `json:"updated_at"`
	CreatedAt  time.Time      `json:"created_at"`
	Address    *AddressOutput `json:"address"`
}

func (u *GetUserProfileUseCase) Execute(data *GetUserProfileInput) (*GetUserProfileOutput, error) {
	user := u.userRepository.FindByID(data.UserId)

	if user == nil {
		return nil, errors.New("user not found with this id")
	}

	userAddr := u.addrRepository.FindByUserID(user.ID.String())

	var outputAddr *AddressOutput

	if userAddr != nil {
		outputAddr = &AddressOutput{
			ID:            userAddr.ID,
			Country:       userAddr.Country,
			ZipCode:       userAddr.ZipCode,
			Street:        userAddr.Street,
			Number:        userAddr.Number,
			Neighbourhood: userAddr.Neighbourhood,
			City:          userAddr.City,
			Estate:        userAddr.Estate,
		}
	}

	return &GetUserProfileOutput{
		ID:         user.ID,
		Name:       user.Name,
		Email:      user.Email,
		Phone:      user.Phone,
		Agency:     user.Agency,
		Bank:       user.Bank,
		Serial:     user.Serial,
		Profession: user.Profession,
		UpdatedAt:  user.UpdatedAt,
		CreatedAt:  user.CreatedAt,
		Address:    outputAddr,
	}, nil
}

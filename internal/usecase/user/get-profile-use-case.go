package user

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/mauFade/infinity/internal/repositories"
)

type GetUserProfileUseCase struct {
	userRepository *repositories.UserRepository
}

func NewGetUserProfileUseCase(r *repositories.UserRepository) *GetUserProfileUseCase {
	return &GetUserProfileUseCase{
		userRepository: r,
	}
}

type GetUserProfileInput struct {
	UserId string
}

type GetUserProfileOutput struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
	Agency     string    `json:"agency"`
	Bank       string    `json:"bank"`
	Serial     string    `json:"serial"`
	Profession string    `json:"profession"`
	UpdatedAt  time.Time `json:"updated_at"`
	CreatedAt  time.Time `json:"created_at"`
}

func (u *GetUserProfileUseCase) Execute(data *GetUserProfileInput) (*GetUserProfileOutput, error) {
	user := u.userRepository.FindByID(data.UserId)

	if user == nil {
		return nil, errors.New("user not found with this id")
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
	}, nil
}

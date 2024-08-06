package user

import (
	"fmt"

	"github.com/mauFade/infinity/internal/repositories"
)

type CreateUserUseCase struct {
	userRepository *repositories.UserRepository
}

type CreateUserInput struct {
	Name       string
	Email      string
	Phone      string
	Password   string
	Profession string
}

func NewCreateUserUseCase(r *repositories.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{
		userRepository: r,
	}
}

func (u *CreateUserUseCase) Execute(data *CreateUserInput) {
	u.validateEmail(data.Email)
}

func (u *CreateUserUseCase) validateEmail(email string) {
	userExist := u.userRepository.FindByEmail(email)

	fmt.Println(userExist)

	if userExist != nil {
		fmt.Println("EXISTE CARAI")
	} else {
		fmt.Println("NÃO EXISTE CARAI")
	}
}

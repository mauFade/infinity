package user

import (
	"errors"
	"fmt"
	"math/rand"
	"net/mail"
	"regexp"
	"time"

	"github.com/google/uuid"
	"github.com/mauFade/infinity/internal/models"
	"github.com/mauFade/infinity/internal/repositories"
	"golang.org/x/crypto/bcrypt"
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

type CreateUserOutput struct {
	ID         uuid.UUID  `json:"id"`
	Name       string     `json:"name"`
	Email      string     `json:"email"`
	Phone      string     `json:"phone"`
	Agency     string     `json:"agency"`
	Bank       string     `json:"bank"`
	Serial     string     `json:"serial"`
	Profession string     `json:"profession"`
	Deleted    bool       `json:"is_deleted"`
	DeletedAt  *time.Time `json:"deleted_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	CreatedAt  time.Time  `json:"created_at"`
}

func NewCreateUserUseCase(r *repositories.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{
		userRepository: r,
	}
}

func (u *CreateUserUseCase) Execute(data *CreateUserInput) (*CreateUserOutput, error) {
	err := u.validateEmail(data.Email)

	if err != nil {
		return nil, err
	}

	err = u.validatePhone(data.Phone)

	if err != nil {
		return nil, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password), 8)

	if err != nil {
		return nil, err
	}

	user := models.NewUser(
		uuid.New(),
		data.Name,
		data.Email,
		fmt.Sprintf("55%s", data.Phone),
		string(hash),
		"0001",
		"001",
		fmt.Sprintf("%s-%v", u.randStringBytes(), rand.Intn(8)+1),
		data.Profession,
		false,
		nil,
		time.Now(),
		time.Now(),
	)

	err = u.userRepository.Create(user)

	if err != nil {
		return nil, err
	}

	return &CreateUserOutput{
		ID:         user.ID,
		Name:       user.Name,
		Email:      user.Email,
		Phone:      user.Phone,
		Agency:     user.Agency,
		Bank:       user.Bank,
		Serial:     user.Serial,
		Profession: user.Profession,
		Deleted:    user.Deleted,
		DeletedAt:  user.DeletedAt,
		UpdatedAt:  user.UpdatedAt,
		CreatedAt:  user.CreatedAt,
	}, nil
}

func (u *CreateUserUseCase) randStringBytes() string {
	numbers := "1234567890"

	b := make([]byte, 6)

	for i := range b {
		b[i] = numbers[rand.Intn(len(numbers))]
	}

	return string(b)
}

func (u *CreateUserUseCase) validatePhone(phone string) error {
	re := regexp.MustCompile(`^\d{11}$`)

	if !re.MatchString(phone) {
		return errors.New("invalid phone number")
	}

	userExist := u.userRepository.FindByPhone(fmt.Sprintf("55%s", phone))

	if userExist != nil && !userExist.Deleted {
		return errors.New("this cellphone is already in use")
	}

	return nil
}

func (u *CreateUserUseCase) validateEmail(email string) error {
	_, err := mail.ParseAddress(email)

	if err != nil {
		return errors.New("invalid email")
	}

	userExist := u.userRepository.FindByEmail(email)

	if userExist != nil && !userExist.Deleted {
		return errors.New("this email is already in use")
	}

	return nil
}

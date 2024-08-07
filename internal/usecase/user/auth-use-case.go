package user

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mauFade/infinity/internal/repositories"
)

type AuthUserUseCase struct {
	userRepository *repositories.UserRepository
}

type generateTokenDTO struct {
	ID    string
	Name  string
	Email string
}

func NewAuthUserUseCase(r *repositories.UserRepository) *AuthUserUseCase {
	return &AuthUserUseCase{
		userRepository: r,
	}
}

type AuthInput struct {
	Email    string
	Password string
}

type authOutput struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}

func (u *AuthUserUseCase) Execute(data *AuthInput) (*authOutput, error) {
	user := u.userRepository.FindByEmail(data.Email)

	if user == nil {
		return nil, errors.New("invalid email")
	}

	err := user.ComparePasswords(data.Password)

	if err != nil {
		return nil, err
	}

	token, err := u.generateToken(&generateTokenDTO{
		ID:    user.ID.String(),
		Name:  user.Name,
		Email: user.Email,
	})

	if err != nil {
		return nil, err
	}

	return &authOutput{
		ID:    user.ID.String(),
		Token: token,
	}, nil
}

func (u *AuthUserUseCase) generateToken(data *generateTokenDTO) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["userID"] = data.ID
	claims["userEmail"] = data.Email
	claims["userName"] = data.Name
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString([]byte("JWT_SECRET"))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

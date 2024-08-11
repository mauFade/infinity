package mocks

import (
	"github.com/mauFade/infinity/internal/models"
	"github.com/stretchr/testify/mock"
)

type UserRepository struct {
	mock.Mock
}

func (m *UserRepository) FindByEmail(email string) *models.User {
	args := m.Called(email)
	if args.Get(0) != nil {
		return args.Get(0).(*models.User)
	}
	return nil
}

func (m *UserRepository) FindByPhone(phone string) *models.User {
	args := m.Called(phone)
	if args.Get(0) != nil {
		return args.Get(0).(*models.User)
	}
	return nil
}

func (m *UserRepository) Create(user *models.User) error {
	return m.Called(user).Error(0)
}

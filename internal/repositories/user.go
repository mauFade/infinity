package repositories

import (
	"github.com/mauFade/infinity/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(d *gorm.DB) *UserRepository {
	return &UserRepository{
		db: d,
	}
}

func (r *UserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) Update(user *models.User) {
	r.db.Save(&user)
}

func (r *UserRepository) FindByEmail(email string) *models.User {
	var user *models.User

	result := r.db.Where("email = ?", email).First(&user)

	if result.RowsAffected == 0 {
		return nil
	}

	return user
}

func (r *UserRepository) FindByPhone(phone string) *models.User {
	var user *models.User

	result := r.db.Where("phone = ?", phone).First(&user)

	if result.RowsAffected == 0 {
		return nil
	}

	return user
}

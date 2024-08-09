package repositories

import (
	"github.com/mauFade/infinity/internal/models"
	"gorm.io/gorm"
)

type AddressRepository struct {
	db *gorm.DB
}

func NewAddressRepository(d *gorm.DB) *AddressRepository {
	return &AddressRepository{
		db: d,
	}
}

func (r *AddressRepository) Create(data models.Address) error {
	return r.db.Create(data).Error
}

func (r *AddressRepository) FindByUserID(userId string) *models.Address {
	var addr *models.Address

	result := r.db.Where("user_id = ?", userId).First(&addr)

	if result.RowsAffected == 0 {
		return nil
	}

	return addr
}

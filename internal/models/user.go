package models

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Address struct {
	ID            uuid.UUID `gorm:"type:uuid"`
	UserID        uuid.UUID `gorm:"type:uuid"`
	Country       string    `gorm:"type:varchar"`
	ZipCode       string    `gorm:"type:varchar"`
	Street        string    `gorm:"type:varchar"`
	Number        string    `gorm:"type:varchar"`
	Neighbourhood string    `gorm:"type:varchar"`
	City          string    `gorm:"type:varchar"`
	Estate        string    `gorm:"type:varchar"`
}

type User struct {
	ID         uuid.UUID  `gorm:"type:uuid"`
	Name       string     `gorm:"type:varchar"`
	Email      string     `gorm:"type:varchar"`
	Phone      string     `gorm:"type:varchar"`
	Password   string     `gorm:"type:varchar"`
	Agency     string     `gorm:"type:varchar"`
	Bank       string     `gorm:"type:varchar"`
	Serial     string     `gorm:"type:varchar"`
	Profession string     `gorm:"type:varchar"`
	Address    Address    `gorm:"foreignKey:UserID"`
	Deleted    bool       `gorm:"type:bool"`
	DeletedAt  *time.Time `gorm:"type:timestamp"`
	UpdatedAt  time.Time  `gorm:"type:timestamp"`
	CreatedAt  time.Time  `gorm:"type:timestamp"`
}

func NewUser(
	id uuid.UUID,
	name,
	email,
	phone,
	password,
	agency,
	bank,
	serial,
	profession string,
	deleted bool,
	deletedAt *time.Time,
	updatedAt,
	createdAt time.Time,
) *User {
	return &User{
		ID:         id,
		Name:       name,
		Email:      email,
		Phone:      phone,
		Password:   password,
		Agency:     agency,
		Bank:       bank,
		Serial:     serial,
		Profession: profession,
		Deleted:    deleted,
		DeletedAt:  deletedAt,
		UpdatedAt:  updatedAt,
		CreatedAt:  createdAt,
	}
}

func NewAddress(
	id,
	userID uuid.UUID,
	country,
	zipCode,
	street,
	number,
	neighbourhood,
	city,
	estate string,
) *Address {
	return &Address{
		ID:            id,
		UserID:        userID,
		Country:       country,
		ZipCode:       zipCode,
		Street:        street,
		Number:        number,
		Neighbourhood: neighbourhood,
		City:          city,
		Estate:        estate,
	}
}

func (u *User) ComparePasswords(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return err
	}

	return nil
}

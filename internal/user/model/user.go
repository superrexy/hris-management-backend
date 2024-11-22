package model

import (
	"hris-management/utils/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	model.BaseModel
	Name     string  `json:"name" validate:"required" gorm:"not null"`
	Email    string  `json:"email" validate:"required,email" gorm:"unique;not null"`
	Role     string  `json:"role" validate:"required,oneof=admin employee manager" gorm:"not null"`
	Password string  `json:"-" validate:"required" gorm:"not null"`
	Avatar   *string `json:"avatar" validate:"omitempty,url" gorm:"default:null"`
	Position string  `json:"position" validate:"required" gorm:"not null"`
	IsActive bool    `json:"is_active" validate:"omitempty" gorm:"default:true"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	// hash the password
	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashed)

	tx.Statement.SetColumn("password", u.Password)

	return
}

func (u *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if u.Password == "" {
		return
	}

	// hash the password
	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashed)

	tx.Statement.SetColumn("password", u.Password)

	return
}

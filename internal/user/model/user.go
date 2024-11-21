package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Role     string  `json:"role" gorm:"type:enum('admin', 'user', 'manager')"`
	Password string  `json:"password"`
	Avatar   *string `json:"avatar"`
	IsActive bool    `json:"is_active" gorm:"default:true"`
}

func (u *User) TableName() string {
	return "users"
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

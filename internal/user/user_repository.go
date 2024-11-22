package user

import (
	"hris-management/config"
	"hris-management/internal/user/model"
)

type UserRepository interface {
	GetAll() ([]model.User, error)
	GetByAttribute(attribute string, value interface{}) (model.User, error)
	Create(user *model.User) (model.User, error)
	Update(user *model.User) (model.User, error)
	Delete(user *model.User) error
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) GetAll() ([]model.User, error) {
	var users []model.User
	result := config.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (r *userRepository) GetByAttribute(attribute string, value interface{}) (model.User, error) {
	var user model.User
	result := config.DB.Where(attribute+" = ?", value).First(&user)
	if result.Error != nil {
		return model.User{}, result.Error
	}
	return user, nil
}

func (r *userRepository) Create(user *model.User) (model.User, error) {
	var data model.User
	result := config.DB.Create(user)
	if result.Error != nil {
		return model.User{}, result.Error
	}

	config.DB.Where("id = ?", user.ID).First(&data)

	return data, nil
}

func (r *userRepository) Update(user *model.User) (model.User, error) {
	result := config.DB.Save(user)
	if result.Error != nil {
		return model.User{}, result.Error
	}
	return *user, nil
}

func (r *userRepository) Delete(user *model.User) error {
	result := config.DB.Delete(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

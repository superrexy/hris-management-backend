package user

import (
	"hris-management/internal/user/model"
)

type UserService interface {
	GetAllUsers() ([]model.User, error)
	GetUserByEmail(email string) (model.User, error)
	GetUserByID(id uint) (model.User, error)
	CreateUser(user *model.User) (model.User, error)
	UpdateUser(user *model.User) (model.User, error)
	DeleteUser(user *model.User) error
}

type userService struct {
	userRepository UserRepository
}

// CreateUser implements UserService.
func (u *userService) CreateUser(user *model.User) (model.User, error) {
	return u.userRepository.Create(user)
}

// DeleteUser implements UserService.
func (u *userService) DeleteUser(user *model.User) error {
	return u.userRepository.Delete(user)
}

// GetAllUsers implements UserService.
func (u *userService) GetAllUsers() ([]model.User, error) {
	return u.userRepository.GetAll()
}

// GetUserByEmail implements UserService.
func (u *userService) GetUserByEmail(email string) (model.User, error) {
	return u.userRepository.GetByAttribute("email", email)
}

// GetUserByID implements UserService.
func (u *userService) GetUserByID(id uint) (model.User, error) {
	return u.userRepository.GetByAttribute("id", id)
}

// UpdateUser implements UserService.
func (u *userService) UpdateUser(user *model.User) (model.User, error) {
	return u.userRepository.Update(user)
}

func NewUserService(repository UserRepository) UserService {
	return &userService{
		userRepository: repository,
	}
}

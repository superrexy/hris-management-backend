package user

import (
	"hris-management/internal/user/dto"
	"hris-management/internal/user/model"
	"hris-management/utils/exception"

	"github.com/gofiber/fiber/v2"
)

type UserService interface {
	GetAllUsers() ([]model.User, error)
	GetUserByEmail(email string) (model.User, error)
	GetUserByID(id uint) (model.User, error)
	CreateUser(payload dto.StoreUserPayload) (model.User, error)
	UpdateUser(id uint, payload dto.UpdateUserPayload) (model.User, error)
	DeleteUser(user model.User) error
}

type userService struct {
	userRepository UserRepository
}

// CreateUser implements UserService.
func (u *userService) CreateUser(payload dto.StoreUserPayload) (model.User, error) {
	if emailExists, _ := u.userRepository.GetByAttribute("email", payload.Email); emailExists.ID != 0 {
		return model.User{}, exception.NewServiceError(fiber.StatusConflict, "Email already exists", nil)
	}

	return u.userRepository.Create(&model.User{
		Name:     payload.Name,
		Email:    payload.Email,
		Role:     payload.Role,
		Password: payload.Password,
		Avatar:   payload.Avatar,
		IsActive: payload.IsActive,
	})
}

// DeleteUser implements UserService.
func (u *userService) DeleteUser(user model.User) error {
	return u.userRepository.Delete(&user)
}

// GetAllUsers implements UserService.
func (u *userService) GetAllUsers() ([]model.User, error) {
	return u.userRepository.GetAll()
}

// GetUserByEmail implements UserService.
func (u *userService) GetUserByEmail(email string) (model.User, error) {
	user, _ := u.userRepository.GetByAttribute("email", email)

	if user.ID == 0 {
		return model.User{}, exception.NewServiceError(404, "User not found", nil)
	}

	return user, nil
}

// GetUserByID implements UserService.
func (u *userService) GetUserByID(id uint) (model.User, error) {
	user, _ := u.userRepository.GetByAttribute("id", id)

	if user.ID == 0 {
		return model.User{}, exception.NewServiceError(404, "User not found", nil)
	}

	return user, nil
}

// UpdateUser implements UserService.
func (u *userService) UpdateUser(id uint, payload dto.UpdateUserPayload) (model.User, error) {
	user, err := u.userRepository.GetByAttribute("id", id)
	if err != nil {
		return model.User{}, exception.NewServiceError(404, "User not found", nil)
	}

	if user.Email != payload.Email {
		if emailExists, _ := u.userRepository.GetByAttribute("email", payload.Email); emailExists.ID != 0 {
			return model.User{}, exception.NewServiceError(fiber.StatusConflict, "Email already exists", nil)
		}
	}

	user.Name = payload.Name
	user.Email = payload.Email
	user.Role = payload.Role
	user.Avatar = payload.Avatar

	if payload.IsActive != nil {
		user.IsActive = *payload.IsActive
	}

	return u.userRepository.Update(&user)
}

func NewUserService(repository UserRepository) UserService {
	return &userService{
		userRepository: repository,
	}
}

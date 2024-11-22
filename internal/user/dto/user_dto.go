package dto

type StoreUserPayload struct {
	Name     string  `json:"name" validate:"required"`
	Email    string  `json:"email" validate:"required,email"`
	Role     string  `json:"role" validate:"required,oneof=admin employee manager"`
	Password string  `json:"password" validate:"required"`
	Avatar   *string `json:"avatar" validate:"omitempty,url"`
	IsActive bool    `json:"is_active" validate:"omitempty"`
}

type UpdateUserPayload struct {
	Name  string `json:"name" validate:"omitempty"`
	Email string `json:"email" validate:"omitempty,email"`
	Role  string `json:"role" validate:"omitempty,oneof=admin employee manager"`
	// Password is omitted because it should not be updated
	Avatar   *string `json:"avatar" validate:"omitempty,url"`
	IsActive *bool   `json:"is_active" validate:"omitempty"`
}

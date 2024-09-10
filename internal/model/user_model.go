package model

type UserResponse struct {
	Name         string `json:"name,omitempty"`
	Email        string `json:"email,omitempty"`
	Phone        string `json:"phone,omitempty"`
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	CreatedAt    int64  `json:"created_at,omitempty"`
	UpdatedAt    int64  `json:"updated_at,omitempty"`
}

type RegisterUserRequest struct {
	Name     string `json:"name" validate:"required,max=100"`
	Email    string `json:"email" validate:"required,max=100"`
	Phone    string `json:"phone" validate:"required,max=20"`
	Password string `json:"password" validate:"required,max=100"`
}

type LoginUserRequest struct {
	Email    string `json:"email" validate:"required,max=100"`
	Password string `json:"password" validate:"required,max=100"`
}

type LogoutUserRequest struct {
	Email string `json:"email" validate:"required,max=100"`
}

type VerifyUserRequest struct {
	Token string `json:"token" validate:"required,max=255"`
}

type UpdateUserRequest struct {
	Email    string `json:"-" validate:"required,max=100"`
	Password string `json:"password" validate:"max=100"`
	Phone    string `json:"phone" validate:"max=20"`
	Name     string `json:"name" validate:"max=100"`
}
type GetUserRequest struct {
	Email string `json:"email" validate:"required,max=100"`
}

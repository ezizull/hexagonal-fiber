// Package auth contains the auth controller
package auth

// LoginRequest is a struct that contains the login request information
type LoginRequest struct {
	Email    string `json:"email" example:"user@mail.com" gorm:"unique" validate:"required,email"`
	Password string `json:"password" example:"Password123" validate:"required,min=8,regexp=^(?=.*[a-z])(?=.*[A-Z])(?=.*\\d)(?=.*[@$!%*?&])[A-Za-z\\d@$!%*?&]+$"`
}

// AccessTokenRequest is a struct that contains the login request information
type AccessTokenRequest struct {
	RefreshToken string `json:"refreshToken" example:"badbunybabybebe" validate:"required"`
}

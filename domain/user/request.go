package user

// NewUser is a struct that contains the request body for the new user
type NewUser struct {
	UserName string `json:"username" example:"someUser" gorm:"unique" validate:"required"`
	Email    string `json:"email" example:"user@mail.com" gorm:"unique" validate:"required,email"`
	Password string `json:"password" example:"Pass@Word123" validate:"required,password"`
	Age      int    `json:"age" example:"1" validate:"required"`
	RoleID   string `json:"role_id" gorm:"index" validate:"required"`
}

// UpdateUser is a struct that contains the request body for the update user
type UpdateUser struct {
	UserName *string `json:"username,omitempty" example:"someUser" gorm:"unique" validate:"-"`
	Email    *string `json:"email,omitempty" example:"mail@mail.com" gorm:"unique" validate:"-"`
	Password *string `json:"password,omitempty" example:"Pass@Word123" validate:"-"`
	Age      *int    `json:"age,omitempty" example:"1" validate:"-"`
	RoleID   *string `json:"role_id,omitempty" gorm:"index" validate:"-"`
}

// LoginRequest is a struct that contains the request body for the login user
type LoginRequest struct {
	Email    string `json:"email" example:"user@mail.com" gorm:"unique" validate:"required,email"`
	Password string `json:"password" example:"Pass@Word123" validate:"required,password"`
}

// AccessTokenRequest is a struct that contains the login request information
type AccessTokenRequest struct {
	RefreshToken string `json:"refreshToken" example:"badbunybabybebe" validate:"required"`
}

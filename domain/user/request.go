package user

// NewUser is a struct that contains the request body for the new user
type NewUser struct {
	UserName string `json:"username" example:"someUser" gorm:"unique" binding:"required"`
	Email    string `json:"email" example:"mail@mail.com" gorm:"unique" binding:"required"`
	Password string `json:"password" example:"Password123" binding:"required"`
	Age      int    `json:"age" example:"1" binding:"required"`
	RoleID   string `json:"role_id" gorm:"index" binding:"required"`
}

// UpdateUser is a struct that contains the request body for the update user
type UpdateUser struct {
	UserName *string `json:"username,omitempty" example:"someUser" gorm:"unique" binding:"-"`
	Email    *string `json:"email,omitempty" example:"mail@mail.com" gorm:"unique" binding:"-"`
	Password *string `json:"password,omitempty" example:"Password123" binding:"-"`
	Age      *int    `json:"age,omitempty" example:"1" binding:"-"`
	RoleID   *string `json:"role_id,omitempty" gorm:"index" binding:"-"`
}

// LoginUser is a struct that contains the request body for the login user
type LoginUser struct {
	Email    string
	Password string
}

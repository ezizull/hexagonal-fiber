package user

import (
	"time"
)

// ResponseUser is a struct that contains the response body for the user
type ResponseUser struct {
	ID        string     `json:"id" example:"cef47ee2-7211-452a-a087-79ce4b8ec3a3"`
	UserName  string     `json:"user" example:"BossonH"`
	Email     string     `json:"email" example:"user@mail.com" gorm:"unique" validate:"required,email"`
	Age       int        `json:"age" example:"1" validate:"required"`
	CreatedAt time.Time  `json:"createdAt,omitempty" example:"2021-02-24 20:19:39" gorm:"autoCreateTime:mili"`
	UpdatedAt time.Time  `json:"updatedAt,omitempty" example:"2021-02-24 20:19:39" gorm:"autoUpdateTime:mili"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" example:"null"`
}

// ResponseUser is a struct that contains the response body for the user
type ResponseUserRole struct {
	ID       string `json:"id" example:"cef47ee2-7211-452a-a087-79ce4b8ec3a3"`
	UserName string `json:"user" example:"BossonH"`
	Email    string `json:"email" example:"user@mail.com" gorm:"unique" validate:"required,email"`
	Role     Role
}

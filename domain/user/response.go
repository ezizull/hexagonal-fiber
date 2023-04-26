package user

import (
	"time"
)

// ResponseUser is a struct that contains the response body for the user
type ResponseUser struct {
	ID        int        `json:"id" example:"1099"`
	UserName  string     `json:"user" example:"BossonH"`
	Email     string     `json:"email" example:"some@mail.com"`
	Age       int        `json:"age" example:"1" binding:"required"`
	CreatedAt time.Time  `json:"createdAt,omitempty" example:"2021-02-24 20:19:39" gorm:"autoCreateTime:mili"`
	UpdatedAt time.Time  `json:"updatedAt,omitempty" example:"2021-02-24 20:19:39" gorm:"autoUpdateTime:mili"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" example:"null"`
}

// ResponseUser is a struct that contains the response body for the user
type ResponseUserRole struct {
	ID       int    `json:"id" example:"1099"`
	UserName string `json:"user" example:"BossonH"`
	Email    string `json:"email" example:"some@mail.com"`
	Role     Role
}

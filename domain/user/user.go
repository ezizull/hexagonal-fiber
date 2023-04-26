// Package user contains the business logic for the user entity
package user

import (
	"time"
)

// User is a struct that contains the user information
type User struct {
	ID           int        `json:"id" example:"1099" gorm:"primaryKey"`
	UserName     string     `json:"userName" example:"UserName" gorm:"column:user_name;uniqueIndex"`
	Email        string     `json:"email" example:"user@mail.com" gorm:"unique" validate:"required,email"`
	HashPassword string     `json:"hash_password" example:"has@Password1"`
	Age          int        `json:"age" example:"1" validate:"required"`
	RoleID       string     `json:"role_id" gorm:"index"`
	CreatedAt    time.Time  `json:"created_at,omitempty" example:"2021-02-24 20:19:39" gorm:"autoCreateTime:mili"`
	UpdatedAt    time.Time  `json:"updated_at,omitempty" example:"2021-02-24 20:19:39" gorm:"autoUpdateTime:mili"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty" example:"null"`
}

// TableName overrides the table name used by User to `users`
func (*User) TableName() string {
	return "users"
}

// PaginationUser is a struct that contains the pagination result for user
type PaginationUser struct {
	Data       []User
	Total      int
	Limit      int
	Current    int
	NextCursor uint
	PrevCursor uint
	NumPages   int
}

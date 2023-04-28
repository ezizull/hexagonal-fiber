// Package user contains the business logic for the user entity
package user

import (
	"time"

	"github.com/google/uuid"
)

// User is a struct that contains the user information
type User struct {
	ID           uuid.UUID  `json:"id" example:"cef47ee2-7211-452a-a087-79ce4b8ec3a3" gorm:"gorm:"type:uuid;default:uuid_generate_v4()"`
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

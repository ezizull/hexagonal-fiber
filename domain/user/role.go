package user

import (
	"time"
)

// Role is a struct that contains the role information
type Role struct {
	ID        int        `json:"id" gorm:"primaryKey"`
	Name      string     `json:"name" gorm:"unique"`
	CreatedAt time.Time  `json:"created_at,omitempty" example:"2021-02-24 20:19:39" gorm:"autoCreateTime:mili"`
	UpdatedAt time.Time  `json:"updated_at,omitempty" example:"2021-02-24 20:19:39" gorm:"autoUpdateTime:mili"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" example:"null"`
}

// UserRole is a struct that contains role of user
type UserRole struct {
	User
	Role Role `gorm:"foreignKey:ID;references:RoleID"`
}

// TableName overrides the table name used by User to `users`
func (*UserRole) TableName() string {
	return "users"
}

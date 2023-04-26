package user

import (
	"time"

	_ "gorm.io/gorm"
)

// Auth contains the data of the authentication
type Auth struct {
	AccessToken               string
	RefreshToken              string
	ExpirationAccessDateTime  time.Time
	ExpirationRefreshDateTime time.Time
}

// DataUserAuthenticated is a struct that contains the data for the authenticated user
type DataUserAuthenticated struct {
	ID       int    `json:"id" example:"123"`
	UserName string `json:"userName" example:"UserName" gorm:"unique"`
	Email    string `json:"email" example:"some@mail.com" gorm:"unique"`
	RoleID   string `json:"role_id" example:"admin"`
	Role     Role
}

// DataSecurityAuthenticated is a struct that contains the security data for the authenticated user
type DataSecurityAuthenticated struct {
	JWTAccessToken            string    `json:"jwtAccessToken" example:"SomeAccessToken"`
	JWTRefreshToken           string    `json:"jwtRefreshToken" example:"SomeRefreshToken"`
	ExpirationAccessDateTime  time.Time `json:"expirationAccessDateTime" example:"2023-02-02T21:03:53.196419-06:00"`
	ExpirationRefreshDateTime time.Time `json:"expirationRefreshDateTime" example:"2023-02-03T06:53:53.196419-06:00"`
}

// SecurityAuthenticatedUser is a struct that contains the data for the authenticated user
type SecurityAuthenticatedUser struct {
	Data     DataUserAuthenticated     `json:"data"`
	Security DataSecurityAuthenticated `json:"security"`
}

package user

import "golang.org/x/crypto/bcrypt"

func (n *NewUser) ToDomainMapper() *User {
	return &User{
		UserName: n.UserName,
		Email:    n.Email,
		Age:      n.Age,
		RoleID:   n.RoleID,
	}
}

func (n UpdateUser) ToDomainMapper() User {
	updateDomain := User{}

	if n.UserName != nil {
		updateDomain.UserName = *n.UserName
	}

	if n.Password != nil {
		hash, _ := bcrypt.GenerateFromPassword([]byte(*n.Password), bcrypt.DefaultCost)
		updateDomain.HashPassword = string(hash)
	}

	if n.Email != nil {
		updateDomain.Email = *n.Email
	}

	if n.Age != nil {
		updateDomain.Age = *n.Age
	}

	return updateDomain
}

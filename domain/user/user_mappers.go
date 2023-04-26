// Package user contains the business logic for the user entity
package user

// ArrayToDomainMapper function to convert list user domain to list user repo
func ArrayToDomainMapper(users *[]User) *[]User {
	usersDomain := make([]User, len(*users))
	for i, user := range *users {
		usersDomain[i] = user
	}

	return &usersDomain
}

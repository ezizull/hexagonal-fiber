package user

func (userRole *UserRole) UserToResponseMapper() (createUserRoleResponse *ResponseUserRole) {
	return &ResponseUserRole{
		ID:       userRole.ID,
		UserName: userRole.UserName,
		Email:    userRole.Email,
		Role:     userRole.Role,
	}
}

func (user *User) DomainToResponseMapper() (createUserResponse *ResponseUser) {
	return &ResponseUser{
		ID:        user.ID,
		UserName:  user.UserName,
		Email:     user.Email,
		Age:       user.Age,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

}

func ArrayDomainToResponseMapper(users *[]User) *[]ResponseUser {
	usersResponse := make([]ResponseUser, len(*users))
	for i, user := range *users {
		usersResponse[i] = *user.DomainToResponseMapper()
	}
	return &usersResponse
}

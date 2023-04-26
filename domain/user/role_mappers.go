package user

// ToRoleDomainMapper function to convert role of user role repo to role domain
func (userRole *UserRole) ToRoleDomainMapper() *Role {
	return &Role{
		ID:        userRole.ID,
		Name:      userRole.Role.Name,
		CreatedAt: userRole.CreatedAt,
		UpdatedAt: userRole.UpdatedAt,
	}
}

// ToDomainMapper function to convert userRole repo to userRole domain
func (userRole *UserRole) ToDomainMapper() *User {
	return &User{
		ID:           userRole.ID,
		UserName:     userRole.UserName,
		Email:        userRole.Email,
		HashPassword: userRole.HashPassword,
		RoleID:       userRole.RoleID,
		CreatedAt:    userRole.CreatedAt,
		UpdatedAt:    userRole.UpdatedAt,
	}
}

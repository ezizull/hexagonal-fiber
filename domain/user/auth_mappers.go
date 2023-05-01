package user

func SecAuthUserMapper(user *User, authInfo *Auth) *SecurityAuthenticatedUser {
	return &SecurityAuthenticatedUser{
		Data: DataUserAuthenticated{
			UserName: user.UserName,
			Email:    user.Email,
			ID:       user.ID.String(),
			RoleID:   user.RoleID,
		},
		Security: DataSecurityAuthenticated{
			JWTAccessToken:            authInfo.AccessToken,
			JWTRefreshToken:           authInfo.RefreshToken,
			ExpirationAccessDateTime:  authInfo.ExpirationAccessDateTime,
			ExpirationRefreshDateTime: authInfo.ExpirationRefreshDateTime,
		},
	}

}

func SecAuthUserRoleMapper(userRole *UserRole, authInfo *Auth) *SecurityAuthenticatedUser {
	return &SecurityAuthenticatedUser{
		Data: DataUserAuthenticated{
			ID:       userRole.ID.String(),
			UserName: userRole.UserName,
			Email:    userRole.Email,
			RoleID:   userRole.RoleID,
			Role:     userRole.Role,
		},
		Security: DataSecurityAuthenticated{
			JWTAccessToken:            authInfo.AccessToken,
			JWTRefreshToken:           authInfo.RefreshToken,
			ExpirationAccessDateTime:  authInfo.ExpirationAccessDateTime,
			ExpirationRefreshDateTime: authInfo.ExpirationRefreshDateTime,
		},
	}

}

func (secureAuth *SecurityAuthenticatedUser) ToUserRoleResponse() *ResponseUserRole {
	return &ResponseUserRole{
		ID:       secureAuth.Data.ID,
		UserName: secureAuth.Data.UserName,
		Email:    secureAuth.Data.Email,
		Role:     secureAuth.Data.Role,
	}
}

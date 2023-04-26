package user

func SecAuthUserMapper(user *User, authInfo *Auth) *SecurityAuthenticatedUser {
	return &SecurityAuthenticatedUser{
		Data: DataUserAuthenticated{
			UserName: user.UserName,
			Email:    user.Email,
			ID:       user.ID,
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
			ID:       userRole.ID,
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

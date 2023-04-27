// Package auth provides the use case for authentication
package auth

import (
	"time"

	"hexagonal-fiber/application/security/jwt"
	userDomain "hexagonal-fiber/domain/user"

	userRepository "hexagonal-fiber/infrastructure/repository/postgres/user"

	"github.com/gofiber/fiber/v2"
)

// Service is a struct that contains the repository implementation for auth use case
type Service struct {
	UserRepository userRepository.Repository
}

// LoginJWT implements the login with jwt methode use case
func (s *Service) LoginJWT(user userDomain.LoginRequest) (*userDomain.SecurityAuthenticatedUser, error) {
	userMap := map[string]interface{}{"email": user.Email}
	userRole, err := s.UserRepository.GetWithRoleByMap(userMap)

	if err != nil || userRole.ID == 0 {
		err = fiber.NewError(fiber.StatusUnauthorized, "email or password does not match")
		return nil, err
	}

	isAuthenticated := CheckPasswordHash(user.Password, userRole.HashPassword)
	if !isAuthenticated {
		err = fiber.NewError(fiber.StatusUnauthorized, "email or password does not match")
		return nil, err
	}

	accessTokenClaims, err := jwt.GenerateJWTToken(userRole.ID, "access", userRole.Role.Name)
	if err != nil {
		return nil, err
	}
	refreshTokenClaims, err := jwt.GenerateJWTToken(userRole.ID, "refresh", userRole.Role.Name)
	if err != nil {
		return nil, err
	}

	return userDomain.SecAuthUserRoleMapper(userRole, &userDomain.Auth{
		AccessToken:               accessTokenClaims.Token,
		RefreshToken:              refreshTokenClaims.Token,
		ExpirationAccessDateTime:  accessTokenClaims.ExpirationTime,
		ExpirationRefreshDateTime: refreshTokenClaims.ExpirationTime,
	}), err
}

// AccessTokenByRefreshToken implements the Access Token By Refresh Token use case
func (s *Service) AccessTokenByRefreshToken(refreshToken string) (*userDomain.SecurityAuthenticatedUser, error) {
	claimsMap, err := jwt.GetClaimsAndVerifyToken(refreshToken, "refresh")
	if err != nil {
		return nil, err
	}

	userMap := map[string]interface{}{"id": claimsMap["user_id"]}
	userRole, err := s.UserRepository.GetWithRoleByMap(userMap)
	if err != nil || userRole.ID == 0 {
		err = fiber.NewError(fiber.StatusNotFound, "user not found")
		return nil, err
	}

	accessTokenClaims, err := jwt.GenerateJWTToken(userRole.ID, "access", userRole.Role.Name)
	if err != nil {
		return nil, err
	}

	var expTime = int64(claimsMap["exp"].(float64))

	return userDomain.SecAuthUserRoleMapper(userRole, &userDomain.Auth{
		AccessToken:               accessTokenClaims.Token,
		ExpirationAccessDateTime:  accessTokenClaims.ExpirationTime,
		RefreshToken:              refreshToken,
		ExpirationRefreshDateTime: time.Unix(expTime, 0),
	}), nil
}

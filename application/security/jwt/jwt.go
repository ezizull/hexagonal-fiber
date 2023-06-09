// Package jwt implements the JWT authentication
package jwt

import (
	"fmt"
	secureDomain "hexagonal-fiber/domain/security"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// GenerateJWTToken generates a JWT token (refresh or access)
func GenerateJWTToken(userID string, tokenType string, roleName string) (appToken *secureDomain.AppToken, err error) {
	tokenTimeUnix, err := getTimeExpire(tokenType)
	if err != nil {
		return
	}

	nowTime := time.Now()
	expirationTokenTime := nowTime.Add(tokenTimeUnix)

	tokenClaims := &secureDomain.Claims{
		UserID: userID,
		Type:   tokenType,
		Role:   roleName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTokenTime.Unix(),
			IssuedAt:  nowTime.UTC().Unix(),
		},
	}
	tokenWithClaims := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), tokenClaims)

	// Sign and get the complete encoded token as a string using the secret
	tokenStr, err := tokenWithClaims.SignedString(secureDomain.PrivateKey)
	if err != nil {
		return
	}

	appToken = &secureDomain.AppToken{
		Token:          tokenStr,
		TokenType:      tokenType,
		ExpirationTime: expirationTokenTime,
	}

	return
}

// GetClaimsAndVerifyToken verifies the token and returns the claims
func GetClaimsAndVerifyToken(tokenString string, tokenType string) (claims jwt.MapClaims, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			message := fmt.Sprintf("unexpected signing method: %v", token.Header["alg"])
			return nil, fiber.NewError(fiber.StatusUnauthorized, message)
		}

		return secureDomain.PublicKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["type"] != tokenType {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "invalid token type")
		}

		var timeExpire = claims["exp"].(float64)
		if time.Now().Unix() > int64(timeExpire) {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "token expired")
		}

		return claims, nil
	}

	return nil, err
}

// ReGenerateCustomJWT regenerate jwt with custom modified data
func ReGenerateCustomJWT(tokenString string, oldClaims *secureDomain.Claims) (newClaims *secureDomain.Claims, err error) {

	return
}

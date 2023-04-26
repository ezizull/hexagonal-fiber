// Package middlewares contains the middlewares for the restapi api
package middlewares

import (
	"net/http"

	"hacktiv/final-project/utils/lists"

	secureDomain "hacktiv/final-project/domain/security"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// AuthJWTMiddleware is a function that validates the jwt token
func AuthJWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token not provided"})
			c.Abort()
			return
		}

		claims := &secureDomain.Claims{}
		_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return secureDomain.PublicKey, nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("Authorized", *claims)
		c.Next()
	}
}

// AuthRoleMiddleware is a function that validates the role of user
func AuthRoleMiddleware(validRoles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get your object from the context
		authData := c.MustGet("Authorized").(secureDomain.Claims)

		if authData.Role == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized"})
			c.Abort()
			return
		}

		if !lists.Contains(validRoles, authData.Role) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized for this path"})
			c.Abort()
			return
		}

		c.Next()
	}
}

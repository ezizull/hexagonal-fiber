// Package middlewares contains the middlewares for the restapi api
package middlewares

import (
	"strconv"

	"hexagonal-fiber/utils/lists"

	secureDomain "hexagonal-fiber/domain/security"

	authConst "hexagonal-fiber/utils/constant/auth"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
)

func AuthJWTMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		tokenString := ctx.Get("Authorization")
		if tokenString == "" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token not provided"})
		}

		claims := &secureDomain.Claims{}
		_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return secureDomain.PublicKey, nil
		})

		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}

		ctx.Set(authConst.AuthRole, claims.Role)
		ctx.Set(authConst.AuthUserID, strconv.Itoa(claims.UserID))

		return ctx.Next()
	}
}

// AuthRoleMiddleware is a function that validates the role of user
func AuthRoleMiddleware(validRoles []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get your object from the context
		authData := ctx.MustGet("Authorized").(secureDomain.Claims)

		if authData.Role == "" {
			ctx.JSON(fiber.StatusUnauthorized, gin.H{"error": "You are not authorized"})
			ctx.Abort()
			return
		}

		if !lists.Contains(validRoles, authData.Role) {
			ctx.JSON(fiber.StatusUnauthorized, gin.H{"error": "You are not authorized for this path"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

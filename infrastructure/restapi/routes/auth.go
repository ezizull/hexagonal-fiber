// Package routes contains all routes of the application
package routes

import (
	authController "hacktiv/final-project/infrastructure/restapi/controllers/auth"

	"github.com/gofiber/fiber/v2"
)

// AuthRoutes is a function that contains all routes of the auth
func AuthRoutes(router fiber.Router, controller *authController.Controller) {

	routerAuth := router.Group("/auth")
	{
		routerAuth.Post("/login", controller.Login)
		routerAuth.Post("/access-token", controller.GetAccessTokenByRefreshToken)
	}

}

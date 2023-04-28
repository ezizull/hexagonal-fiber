package routes

import (
	userController "hexagonal-fiber/infrastructure/restapi/controllers/user"
	"hexagonal-fiber/infrastructure/restapi/middlewares"

	"github.com/gofiber/fiber/v2"
)

// UserRoutes is a function that contains all routes of the user
func UserRoutes(router fiber.Router, controller *userController.Controller) {
	routerAuth := router.Group("/user")

	// authentication
	routerAuth.Use(middlewares.AuthJWTMiddleware())
	{
		routerAuth.Get("/:id", controller.GetUsersByID)
		routerAuth.Put("/:id", controller.UpdateUser)
		routerAuth.Delete("/:id", controller.DeleteUser)
	}

	// authorization
	routerAuth.Use(middlewares.AuthRoleMiddleware([]string{"admin"}))
	{
		routerAuth.Get("", controller.GetAllUsers)
	}
}

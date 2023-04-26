package routes

import (
	userController "hacktiv/final-project/infrastructure/restapi/controllers/user"
	"hacktiv/final-project/infrastructure/restapi/middlewares"

	"github.com/gofiber/fiber/v2"
)

// UserRoutes is a function that contains all routes of the user
func UserRoutes(router fiber.Router, controller *userController.Controller) {
	routerAuth := router.Group("/user")

	// public
	{
		routerAuth.Post("", controller.NewUser)
	}

	// authentication
	routerAuth.Use(middlewares.AuthJWTMiddleware())
	{
		routerAuth.Get("/:id", controller.GetUsersByID)
		routerAuth.Put("/:id", controller.UpdateUser)
		routerAuth.Delete("/:id", controller.DeleteUser)
	}

	// admin role
	routerAuth.Use(middlewares.AuthRoleMiddleware([]string{"admin"}))
	{
		routerAuth.Get("", controller.GetAllUsers)
	}
}

package routes

import (
	sosmedController "hacktiv/final-project/infrastructure/restapi/controllers/sosmed"
	"hacktiv/final-project/infrastructure/restapi/middlewares"

	"github.com/gofiber/fiber/v2"
)

// SocialMediaRoutes is a function that contains all routes of the sosmed
func SocialMediaRoutes(router fiber.Router, controller *sosmedController.Controller) {
	routerSocialMedia := router.Group("/social-media")

	// authentication
	routerSocialMedia.Use(middlewares.AuthJWTMiddleware())
	{
		routerSocialMedia.Get("", controller.GetAllSocialMedia)
		routerSocialMedia.Get("/own", controller.GetAllOwnSocialMedia)
		routerSocialMedia.Get("/:id", controller.GetSocialMediaByID)
		routerSocialMedia.Post("", controller.NewSocialMedia)
		routerSocialMedia.Put("/:id", controller.UpdateSocialMedia)
		routerSocialMedia.Delete("/:id", controller.DeleteSocialMedia)
	}
}

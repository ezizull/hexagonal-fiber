package routes

import (
	photoController "hexagonal-fiber/infrastructure/restapi/controllers/photo"
	"hexagonal-fiber/infrastructure/restapi/middlewares"

	"github.com/gofiber/fiber/v2"
)

// PhotoRoutes is a function that contains all routes of the photo
func PhotoRoutes(router fiber.Router, controller *photoController.Controller) {
	routerPhoto := router.Group("/photos")

	// authentication
	routerPhoto.Use(middlewares.AuthJWTMiddleware())
	{
		routerPhoto.Get("", controller.GetAllPhotos)
		routerPhoto.Get("/own", controller.GetAllOwnPhotos)
		routerPhoto.Get("/:id/comments", controller.GetPhotoWithComments)
		routerPhoto.Get("/:id", controller.GetPhotoByID)
		routerPhoto.Post("", controller.NewPhoto)
		routerPhoto.Put("/:id", controller.UpdatePhoto)
		routerPhoto.Delete("/:id", controller.DeletePhoto)
	}
}

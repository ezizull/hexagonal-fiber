// Package routes contains all routes of the application
package routes

import (
	// swaggerFiles for documentation
	_ "hexagonal-fiber/docs"
	"hexagonal-fiber/infrastructure/restapi/adapter"

	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"

	"gorm.io/gorm"
)

// Security is a struct that contains the security of the application
// @SecurityDefinitions.jwt
type Security struct {
	Authorization string `header:"Authorization" json:"Authorization"`
}

func ApplicationRootRouter(router fiber.Router, db *gorm.DB) {
	// Documentation Swagger
	{
		router.Get("/swagger/*any", fiberSwagger.WrapHandler)
		router.Get("/", func(c *fiber.Ctx) error {
			return c.Redirect("/swagger/index.html", fiber.StatusMovedPermanently)
		})
	}
}

func ApplicationV1Router(router fiber.Router, db *gorm.DB) {
	routerV1 := router.Group("/v1")

	{
		// Auth Routes
		AuthRoutes(routerV1, adapter.AuthAdapter(db))

		// User Routes
		UserRoutes(routerV1, adapter.UserAdapter(db))

		// Photo Routes
		PhotoRoutes(routerV1, adapter.PhotoAdapter(db))

		// SocialMedia Routes
		SocialMediaRoutes(routerV1, adapter.SocialMediaAdapter(db))

		// Comment Routes
		CommentRoutes(routerV1, adapter.CommentAdapter(db))

	}
}

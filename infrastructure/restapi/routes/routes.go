// Package routes contains all routes of the application
package routes

import (
	_ "hexagonal-fiber/docs"
	databsDomain "hexagonal-fiber/domain/database"
	"hexagonal-fiber/infrastructure/restapi/adapter"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func ApplicationRootRouter(router fiber.Router, db databsDomain.Database) {
	// Monitoring
	{
		router.Get("/metrics", monitor.New())
	}

	// Documentation Swagger
	{
		router.Get("/swagger/*any", fiberSwagger.WrapHandler)
		router.Get("/", func(c *fiber.Ctx) error {
			return c.Redirect("/swagger/index.html", fiber.StatusMovedPermanently)
		})
	}
}

func ApplicationV1Router(router fiber.Router, db databsDomain.Database) {
	routerV1 := router.Group("/v1")

	{
		// Auth Routes
		AuthRoutes(routerV1, adapter.AuthAdapter(db))

		// CSRF Middleware
		{
			router.Use(csrf.New(csrf.ConfigDefault))
		}

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

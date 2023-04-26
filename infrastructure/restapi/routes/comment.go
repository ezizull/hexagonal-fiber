package routes

import (
	commentController "hexagonal-fiber/infrastructure/restapi/controllers/comment"
	"hexagonal-fiber/infrastructure/restapi/middlewares"

	"github.com/gofiber/fiber/v2"
)

// CommentRoutes is a function that contains all routes of the comment
func CommentRoutes(router fiber.Router, controller *commentController.Controller) {
	routerComment := router.Group("/comments")

	// authentication
	routerComment.Use(middlewares.AuthJWTMiddleware())
	{
		routerComment.Get("", controller.GetAllComments)
		routerComment.Get("/own", controller.GetAllOwnComments)
		routerComment.Get("/:id", controller.GetCommentByID)
		routerComment.Post("", controller.NewComment)
		routerComment.Put("/:id", controller.UpdateComment)
		routerComment.Delete("/:id", controller.DeleteComment)
	}
}

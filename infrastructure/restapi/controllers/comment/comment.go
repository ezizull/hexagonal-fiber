package comment

import (
	useCaseComment "hexagonal-fiber/application/usecases/comment"
	commentDomain "hexagonal-fiber/domain/comment"
	secureDomain "hexagonal-fiber/domain/security"

	authConst "hexagonal-fiber/utils/constant/auth"
	mssgConst "hexagonal-fiber/utils/constant/message"

	"github.com/gofiber/fiber/v2"
)

// Controller is a struct that contains the comment service
type Controller struct {
	CommentService useCaseComment.Service
}

// NewComment godoc
// @Tags comment
// @Summary Create New CommentName
// @Description Create new comment on the system
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param data body NewComment true "body data"
// @Success 200 {object} commentDomain.Comment
// @Failure 400 {object} controllers.MessageResponse
// @Failure 500 {object} controllers.MessageResponse
// @Router /comment [post]
func (c *Controller) NewComment(ctx *fiber.Ctx) (err error) {
	authData := ctx.Locals(authConst.Authorized).(*secureDomain.Claims)

	var request commentDomain.NewComment
	if err := ctx.BodyParser(&request); err != nil {
		appError := fiber.NewError(fiber.StatusBadRequest, mssgConst.StatusBadRequest)
		return ctx.Status(fiber.StatusBadRequest).JSON(appError)
	}

	request.UserID = authData.UserID
	if err := createValidation(request); err != nil {
		appError := fiber.NewError(fiber.StatusBadRequest, mssgConst.ValidationError)
		return ctx.Status(fiber.StatusBadRequest).JSON(appError)
	}

	comment, err := c.CommentService.Create(&request)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
		return
	}

	return ctx.Status(fiber.StatusCreated).JSON(comment)
}

// GetAllComments godoc
// @Tags comment
// @Summary Get all Comments
// @Security ApiKeyAuth
// @Description Get all Comments on the system
// @Success 200 {object} commentDomain.PaginationComment
// @Failure 400 {object} controllers.MessageResponse
// @Failure 500 {object} controllers.MessageResponse
// @Router /comment [get]
func (c *Controller) GetAllComments(ctx *fiber.Ctx) (err error) {
	page := ctx.QueryInt("page", 1)
	limit := ctx.QueryInt("limit", 20)

	var request commentDomain.GetComment
	if err := ctx.BodyParser(&request); err != nil {
		appError := fiber.NewError(fiber.StatusBadRequest, mssgConst.StatusBadRequest)
		return ctx.Status(fiber.StatusBadRequest).JSON(appError)
	}

	comments, err := c.CommentService.GetAll(page, limit)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
		return
	}

	return ctx.Status(fiber.StatusOK).JSON(comments)
}

// GetAllOwnComments godoc
// @Tags comment
// @Summary Get all Comments
// @Security ApiKeyAuth
// @Description Get all Comments on the system
// @Success 200 {object} commentDomain.PaginationComment
// @Failure 400 {object} controllers.MessageResponse
// @Failure 500 {object} controllers.MessageResponse
// @Router /comment [get]
func (c *Controller) GetAllOwnComments(ctx *fiber.Ctx) (err error) {
	authData := ctx.Locals(authConst.Authorized).(*secureDomain.Claims)

	page := ctx.QueryInt("page", 1)
	limit := ctx.QueryInt("limit", 20)

	comments, err := c.CommentService.UserGetAll(authData.UserID, page, limit)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
		return
	}

	return ctx.Status(fiber.StatusOK).JSON(comments)
}

// GetCommentByID godoc
// @Tags comment
// @Summary Get comments by ID
// @Description Get Comments by ID on the system
// @Param comment_id path int true "id of comment"
// @Security ApiKeyAuth
// @Success 200 {object} commentDomain.Comment
// @Failure 400 {object} controllers.MessageResponse
// @Failure 500 {object} controllers.MessageResponse
// @Router /comment/{comment_id} [get]
func (c *Controller) GetCommentByID(ctx *fiber.Ctx) (err error) {
	commentID := ctx.Params("id")
	comment, err := c.CommentService.GetByID(commentID)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
		return
	}

	return ctx.Status(fiber.StatusOK).JSON(comment)
}

// UpdateComment godoc
// @Tags comment
// @Summary Get comments by ID
// @Description Get Comments by ID on the system
// @Param comment_id path int true "id of comment"
// @Security ApiKeyAuth
// @Success 200 {object} commentDomain.Comment
// @Failure 400 {object} controllers.MessageResponse
// @Failure 500 {object} controllers.MessageResponse
// @Router /comment/{comment_id} [get]
func (c *Controller) UpdateComment(ctx *fiber.Ctx) (err error) {
	authData := ctx.Locals(authConst.Authorized).(*secureDomain.Claims)
	commentID := ctx.Params("id")

	var request commentDomain.UpdateComment
	if err := ctx.BodyParser(&request); err != nil {
		appError := fiber.NewError(fiber.StatusBadRequest, mssgConst.StatusBadRequest)
		return ctx.Status(fiber.StatusBadRequest).JSON(appError)
	}

	if err = updateValidation(&request); err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
		return
	}

	var comment *commentDomain.Comment

	if authData.Role == "admin" {
		comment, err = c.CommentService.Update(commentID, request)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
			return
		}
	} else {
		comment, err = c.CommentService.UserUpdate(commentID, authData.UserID, request)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
			return
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(comment)
}

// DeleteComment godoc
// @Tags comment
// @Summary Get comments by ID
// @Description Get Comments by ID on the system
// @Param comment_id path int true "id of comment"
// @Security ApiKeyAuth
// @Success 200 {object} controllers.MessageResponse
// @Failure 400 {object} controllers.MessageResponse
// @Failure 500 {object} controllers.MessageResponse
// @Router /comment/{comment_id} [get]
func (c *Controller) DeleteComment(ctx *fiber.Ctx) (err error) {
	commentID := ctx.Params("id")
	if err = c.CommentService.Delete(commentID); err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
		return
	}

	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "resource deleted successfully"})
}

package comment

import (
	"errors"
	"net/http"
	"strconv"

	useCaseComment "hacktiv/final-project/application/usecases/comment"
	commentDomain "hacktiv/final-project/domain/comment"
	errorDomain "hacktiv/final-project/domain/errors"
	secureDomain "hacktiv/final-project/domain/security"
	"hacktiv/final-project/infrastructure/restapi/controllers"

	"github.com/gin-gonic/gin"
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
func (c *Controller) NewComment(ctx *gin.Context) {
	// Get your object from the context
	authData := ctx.MustGet("Authorized").(secureDomain.Claims)

	var request commentDomain.NewComment
	if err := controllers.BindJSON(ctx, &request); err != nil {
		appError := errorDomain.NewAppError(err, errorDomain.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	request.UserID = authData.UserID
	err := createValidation(request)
	if err != nil {
		appError := errorDomain.NewAppError(err, errorDomain.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	comment, err := c.CommentService.Create(&request)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, comment)
}

// GetAllComments godoc
// @Tags comment
// @Summary Get all Comments
// @Security ApiKeyAuth
// @Description Get all Comments on the system
// @Success 200 {object} commentDomain.PaginationResultComment
// @Failure 400 {object} controllers.MessageResponse
// @Failure 500 {object} controllers.MessageResponse
// @Router /comment [get]
func (c *Controller) GetAllComments(ctx *gin.Context) {
	pageStr := ctx.DefaultQuery("page", "1")
	limitStr := ctx.DefaultQuery("limit", "20")

	page, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		appError := errorDomain.NewAppError(errors.New("param page is necessary to be an integer"), errorDomain.ValidationError)
		_ = ctx.Error(appError)
		return
	}
	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		appError := errorDomain.NewAppError(errors.New("param limit is necessary to be an integer"), errorDomain.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	var request commentDomain.GetComment
	if err := controllers.BindJSON(ctx, &request); err != nil {
		appError := errorDomain.NewAppError(err, errorDomain.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	comments, err := c.CommentService.GetAll(page, limit)
	if err != nil {
		appError := errorDomain.NewAppErrorWithType(errorDomain.UnknownError)
		_ = ctx.Error(appError)
		return
	}

	ctx.JSON(http.StatusOK, comments)
}

// GetAllOwnComments godoc
// @Tags comment
// @Summary Get all Comments
// @Security ApiKeyAuth
// @Description Get all Comments on the system
// @Success 200 {object} commentDomain.PaginationResultComment
// @Failure 400 {object} controllers.MessageResponse
// @Failure 500 {object} controllers.MessageResponse
// @Router /comment [get]
func (c *Controller) GetAllOwnComments(ctx *gin.Context) {
	// Get your object from the context
	authData := ctx.MustGet("Authorized").(secureDomain.Claims)

	pageStr := ctx.DefaultQuery("page", "1")
	limitStr := ctx.DefaultQuery("limit", "20")

	page, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		appError := errorDomain.NewAppError(errors.New("param page is necessary to be an integer"), errorDomain.ValidationError)
		_ = ctx.Error(appError)
		return
	}
	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		appError := errorDomain.NewAppError(errors.New("param limit is necessary to be an integer"), errorDomain.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	comments, err := c.CommentService.UserGetAll(authData.UserID, page, limit)
	if err != nil {
		appError := errorDomain.NewAppErrorWithType(errorDomain.UnknownError)
		_ = ctx.Error(appError)
		return
	}

	ctx.JSON(http.StatusOK, comments)
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
func (c *Controller) GetCommentByID(ctx *gin.Context) {
	commentID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		appError := errorDomain.NewAppError(errors.New("comment id is invalid"), errorDomain.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	comment, err := c.CommentService.GetByID(commentID)
	if err != nil {
		appError := errorDomain.NewAppError(err, errorDomain.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	ctx.JSON(http.StatusOK, comment)
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
func (c *Controller) UpdateComment(ctx *gin.Context) {
	// Get your object from the context
	authData := ctx.MustGet("Authorized").(secureDomain.Claims)

	commentID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		appError := errorDomain.NewAppError(errors.New("param id is necessary in the url"), errorDomain.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	var request commentDomain.UpdateComment
	err = controllers.BindJSON(ctx, &request)
	if err != nil {
		appError := errorDomain.NewAppError(err, errorDomain.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	err = updateValidation(&request)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	var comment *commentDomain.Comment

	if authData.Role == "admin" {
		comment, err = c.CommentService.Update(commentID, request)
		if err != nil {
			_ = ctx.Error(err)
			return
		}
	} else {
		comment, err = c.CommentService.UserUpdate(commentID, authData.UserID, request)
		if err != nil {
			_ = ctx.Error(err)
			return
		}
	}

	ctx.JSON(http.StatusOK, comment)
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
func (c *Controller) DeleteComment(ctx *gin.Context) {
	commentID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		appError := errorDomain.NewAppError(errors.New("param id is necessary in the url"), errorDomain.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	err = c.CommentService.Delete(commentID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "resource deleted successfully"})

}

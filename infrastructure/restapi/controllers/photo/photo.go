// Package photo contains the photo controller
package photo

import (
	"errors"
	"net/http"
	"strconv"

	useCasePhoto "hacktiv/final-project/application/usecases/photo"
	errorDomain "hacktiv/final-project/domain/errors"
	photoDomain "hacktiv/final-project/domain/photo"
	secureDomain "hacktiv/final-project/domain/security"
	"hacktiv/final-project/infrastructure/restapi/controllers"

	"github.com/gin-gonic/gin"
)

// Controller is a struct that contains the photo service
type Controller struct {
	PhotoService useCasePhoto.Service
}

// NewPhoto godoc
// @Tags photo
// @Summary Create New PhotoName
// @Description Create new photo on the system
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param data body NewPhoto true "body data"
// @Success 200 {object} photoDomain.Photo
// @Failure 400 {object} controllers.MessageResponse
// @Failure 500 {object} controllers.MessageResponse
// @Router /photo [post]
func (c *Controller) NewPhoto(ctx *gin.Context) {
	// Get your object from the context
	authData := ctx.MustGet("Authorized").(secureDomain.Claims)

	var request photoDomain.NewPhoto
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

	photo, err := c.PhotoService.Create(&request)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, photo)
}

// GetAllPhotos godoc
// @Tags photo
// @Summary Get all Photos
// @Security ApiKeyAuth
// @Description Get all Photos on the system
// @Success 200 {object} photoDomain.PaginationResultPhoto
// @Failure 400 {object} controllers.MessageResponse
// @Failure 500 {object} controllers.MessageResponse
// @Router /photo [get]
func (c *Controller) GetAllPhotos(ctx *gin.Context) {
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

	photos, err := c.PhotoService.GetAll(page, limit)
	if err != nil {
		appError := errorDomain.NewAppErrorWithType(errorDomain.UnknownError)
		_ = ctx.Error(appError)
		return
	}

	ctx.JSON(http.StatusOK, photos)
}

// GetAllOwnPhotos godoc
// @Tags photo
// @Summary Get all Photos
// @Security ApiKeyAuth
// @Description Get all Photos on the system
// @Success 200 {object} photoDomain.PaginationResultPhoto
// @Failure 400 {object} controllers.MessageResponse
// @Failure 500 {object} controllers.MessageResponse
// @Router /photo [get]
func (c *Controller) GetAllOwnPhotos(ctx *gin.Context) {
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

	photos, err := c.PhotoService.UserGetAll(authData.UserID, page, limit)
	if err != nil {
		appError := errorDomain.NewAppErrorWithType(errorDomain.UnknownError)
		_ = ctx.Error(appError)
		return
	}

	ctx.JSON(http.StatusOK, photos)
}

// GetPhotoWithComments godoc
// @Tags photo
// @Summary Get photos by ID
// @Description Get Photos by ID on the system
// @Param photo_id path int true "id of photo"
// @Security ApiKeyAuth
// @Success 200 {object} photoDomain.ResponsePhotoComments
// @Failure 400 {object} controllers.MessageResponse
// @Failure 500 {object} controllers.MessageResponse
// @Router /photo/{photo_id} [get]
func (c *Controller) GetPhotoWithComments(ctx *gin.Context) {
	photoID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		appError := errorDomain.NewAppError(errors.New("photo id is invalid"), errorDomain.ValidationError)
		_ = ctx.Error(appError)
		return
	}

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

	photo, err := c.PhotoService.GetWithComments(photoID, page, limit)
	if err != nil {
		appError := errorDomain.NewAppError(err, errorDomain.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	ctx.JSON(http.StatusOK, photo)
}

// GetPhotoByID godoc
// @Tags photo
// @Summary Get photos by ID
// @Description Get Photos by ID on the system
// @Param photo_id path int true "id of photo"
// @Security ApiKeyAuth
// @Success 200 {object} photoDomain.Photo
// @Failure 400 {object} controllers.MessageResponse
// @Failure 500 {object} controllers.MessageResponse
// @Router /photo/{photo_id} [get]
func (c *Controller) GetPhotoByID(ctx *gin.Context) {
	photoID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		appError := errorDomain.NewAppError(errors.New("photo id is invalid"), errorDomain.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	photo, err := c.PhotoService.GetByID(photoID)
	if err != nil {
		appError := errorDomain.NewAppError(err, errorDomain.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	ctx.JSON(http.StatusOK, photo)
}

// UpdatePhoto godoc
// @Tags photo
// @Summary Get photos by ID
// @Description Get Photos by ID on the system
// @Param photo_id path int true "id of photo"
// @Security ApiKeyAuth
// @Success 200 {object} photoDomain.Photo
// @Failure 400 {object} controllers.MessageResponse
// @Failure 500 {object} controllers.MessageResponse
// @Router /photo/{photo_id} [get]
func (c *Controller) UpdatePhoto(ctx *gin.Context) {
	// Get your object from the context
	authData := ctx.MustGet("Authorized").(secureDomain.Claims)

	photoID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		appError := errorDomain.NewAppError(errors.New("param id is necessary in the url"), errorDomain.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	var request photoDomain.UpdatePhoto
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

	var photo *photoDomain.Photo

	if authData.Role == "admin" {
		photo, err = c.PhotoService.Update(photoID, request)
		if err != nil {
			_ = ctx.Error(err)
			return
		}
	} else {
		photo, err = c.PhotoService.UserUpdate(photoID, authData.UserID, request)
		if err != nil {
			_ = ctx.Error(err)
			return
		}
	}

	ctx.JSON(http.StatusOK, photo)
}

// DeletePhoto godoc
// @Tags photo
// @Summary Get photos by ID
// @Description Get Photos by ID on the system
// @Param photo_id path int true "id of photo"
// @Security ApiKeyAuth
// @Success 200 {object} controllers.MessageResponse
// @Failure 400 {object} controllers.MessageResponse
// @Failure 500 {object} controllers.MessageResponse
// @Router /photo/{photo_id} [get]
func (c *Controller) DeletePhoto(ctx *gin.Context) {
	photoID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		appError := errorDomain.NewAppError(errors.New("param id is necessary in the url"), errorDomain.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	err = c.PhotoService.Delete(photoID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "resource deleted successfully"})

}

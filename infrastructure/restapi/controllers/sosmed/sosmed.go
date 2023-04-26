// Package sosmed contains the sosmed controller
package sosmed

import (
	"errors"
	"net/http"
	"strconv"

	useCaseSosmed "hacktiv/final-project/application/usecases/sosmed"
	errorDomain "hacktiv/final-project/domain/errors"
	secureDomain "hacktiv/final-project/domain/security"
	sosmedDomain "hacktiv/final-project/domain/sosmed"
	"hacktiv/final-project/infrastructure/restapi/controllers"

	"github.com/gin-gonic/gin"
)

// Controller is a struct that contains the sosmed service
type Controller struct {
	SocialMediaService useCaseSosmed.Service
}

// NewSocialMedia godoc
// @Tags sosmed
// @Summary Create New SocialMediaName
// @Description Create new sosmed on the system
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param data body NewSocialMedia true "body data"
// @Success 200 {object} sosmedDomain.SocialMedia
// @Failure 400 {object} controllers.MessageResponse
// @Failure 500 {object} controllers.MessageResponse
// @Router /sosmed [post]
func (c *Controller) NewSocialMedia(ctx *gin.Context) {
	// Get your object from the context
	authData := ctx.MustGet("Authorized").(secureDomain.Claims)

	var request sosmedDomain.NewSocialMedia
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

	sosmed, err := c.SocialMediaService.Create(&request)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, sosmed)
}

// GetAllSocialMedia godoc
// @Tags sosmed
// @Summary Get all SocialMedia
// @Security ApiKeyAuth
// @Description Get all SocialMedia on the system
// @Success 200 {object} sosmedDomain.PaginationResultSocialMedia
// @Failure 400 {object} controllers.MessageResponse
// @Failure 500 {object} controllers.MessageResponse
// @Router /sosmed [get]
func (c *Controller) GetAllSocialMedia(ctx *gin.Context) {
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

	sosmeds, err := c.SocialMediaService.GetAll(page, limit)
	if err != nil {
		appError := errorDomain.NewAppErrorWithType(errorDomain.UnknownError)
		_ = ctx.Error(appError)
		return
	}

	ctx.JSON(http.StatusOK, sosmeds)
}

// GetAllOwnSocialMedia godoc
// @Tags sosmed
// @Summary Get all SocialMedia
// @Security ApiKeyAuth
// @Description Get all SocialMedia on the system
// @Success 200 {object} sosmedDomain.PaginationResultSocialMedia
// @Failure 400 {object} controllers.MessageResponse
// @Failure 500 {object} controllers.MessageResponse
// @Router /sosmed [get]
func (c *Controller) GetAllOwnSocialMedia(ctx *gin.Context) {
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

	sosmeds, err := c.SocialMediaService.UserGetAll(authData.UserID, page, limit)
	if err != nil {
		appError := errorDomain.NewAppErrorWithType(errorDomain.UnknownError)
		_ = ctx.Error(appError)
		return
	}

	ctx.JSON(http.StatusOK, sosmeds)
}

// GetSocialMediaByID godoc
// @Tags sosmed
// @Summary Get sosmeds by ID
// @Description Get SocialMedia by ID on the system
// @Param sosmed_id path int true "id of sosmed"
// @Security ApiKeyAuth
// @Success 200 {object} sosmedDomain.SocialMedia
// @Failure 400 {object} controllers.MessageResponse
// @Failure 500 {object} controllers.MessageResponse
// @Router /sosmed/{sosmed_id} [get]
func (c *Controller) GetSocialMediaByID(ctx *gin.Context) {
	sosmedID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		appError := errorDomain.NewAppError(errors.New("sosmed id is invalid"), errorDomain.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	sosmed, err := c.SocialMediaService.GetByID(sosmedID)
	if err != nil {
		appError := errorDomain.NewAppError(err, errorDomain.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	ctx.JSON(http.StatusOK, sosmed)
}

// UpdateSocialMedia godoc
// @Tags sosmed
// @Summary Get sosmeds by ID
// @Description Get SocialMedia by ID on the system
// @Param sosmed_id path int true "id of sosmed"
// @Security ApiKeyAuth
// @Success 200 {object} sosmedDomain.SocialMedia
// @Failure 400 {object} controllers.MessageResponse
// @Failure 500 {object} controllers.MessageResponse
// @Router /sosmed/{sosmed_id} [get]
func (c *Controller) UpdateSocialMedia(ctx *gin.Context) {
	// Get your object from the context
	authData := ctx.MustGet("Authorized").(secureDomain.Claims)

	sosmedID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		appError := errorDomain.NewAppError(errors.New("param id is necessary in the url"), errorDomain.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	var request sosmedDomain.UpdateSocialMedia
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

	var sosmed *sosmedDomain.SocialMedia

	if authData.Role == "admin" {
		sosmed, err = c.SocialMediaService.Update(sosmedID, request)
		if err != nil {
			_ = ctx.Error(err)
			return
		}
	} else {
		sosmed, err = c.SocialMediaService.UserUpdate(sosmedID, authData.UserID, request)
		if err != nil {
			_ = ctx.Error(err)
			return
		}
	}

	ctx.JSON(http.StatusOK, sosmed)
}

// DeleteSocialMedia godoc
// @Tags sosmed
// @Summary Get sosmeds by ID
// @Description Get SocialMedia by ID on the system
// @Param sosmed_id path int true "id of sosmed"
// @Security ApiKeyAuth
// @Success 200 {object} controllers.MessageResponse
// @Failure 400 {object} controllers.MessageResponse
// @Failure 500 {object} controllers.MessageResponse
// @Router /sosmed/{sosmed_id} [get]
func (c *Controller) DeleteSocialMedia(ctx *gin.Context) {
	sosmedID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		appError := errorDomain.NewAppError(errors.New("param id is necessary in the url"), errorDomain.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	err = c.SocialMediaService.Delete(sosmedID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "resource deleted successfully"})

}

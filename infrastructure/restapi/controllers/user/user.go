// Package user contains the user controller
package user

import (
	"errors"
	"net/http"
	"strconv"

	useCaseUser "hacktiv/final-project/application/usecases/user"
	errorDomain "hacktiv/final-project/domain/errors"
	secureDomain "hacktiv/final-project/domain/security"
	userDomain "hacktiv/final-project/domain/user"
	"hacktiv/final-project/infrastructure/restapi/controllers"

	"github.com/gin-gonic/gin"
)

// Controller is a struct that contains the user service
type Controller struct {
	UserService useCaseUser.Service
}

// NewUser godoc
// @Tags user
// @Summary Create New UserName
// @Description Create new user on the system
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param data body NewUser true "body data"
// @Success 200 {object} ResponseUser
// @Failure 400 {object} controllers.MessageResponse
// @Failure 500 {object} controllers.MessageResponse
// @Router /user [post]
func (c *Controller) NewUser(ctx *gin.Context) {
	var request userDomain.NewUser

	if err := controllers.BindJSON(ctx, &request); err != nil {
		appError := errorDomain.NewAppError(err, errorDomain.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	err := createValidation(request)
	if err != nil {
		appError := errorDomain.NewAppError(err, errorDomain.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	user, err := c.UserService.Create(request)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	userResponse := user.DomainToResponseMapper()
	ctx.JSON(http.StatusOK, userResponse)
}

// GetAllUsers godoc
// @Tags user
// @Summary Get all Users
// @Security ApiKeyAuth
// @Description Get all Users on the system
// @Success 200 {object} []ResponseUser
// @Failure 400 {object} controllers.MessageResponse
// @Failure 500 {object} controllers.MessageResponse
// @Router /user [get]
func (c *Controller) GetAllUsers(ctx *gin.Context) {
	users, err := c.UserService.GetAll()
	if err != nil {
		appError := errorDomain.NewAppErrorWithType(errorDomain.UnknownError)
		_ = ctx.Error(appError)
		return
	}

	ctx.JSON(http.StatusOK, userDomain.ArrayDomainToResponseMapper(users))
}

// GetUsersByID godoc
// @Tags user
// @Summary Get users by ID
// @Description Get Users by ID on the system
// @Param user_id path int true "id of user"
// @Security ApiKeyAuth
// @Success 200 {object} ResponseUser
// @Failure 400 {object} controllers.MessageResponse
// @Failure 500 {object} controllers.MessageResponse
// @Router /user/{user_id} [get]
func (c *Controller) GetUsersByID(ctx *gin.Context) {
	// Get your object from the context
	authData := ctx.MustGet("Authorized").(secureDomain.Claims)

	if authData.Role == "admin" {
		userID, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			appError := errorDomain.NewAppError(errors.New("user id is invalid"), errorDomain.ValidationError)
			_ = ctx.Error(appError)
			return
		}

		userRole, err := c.UserService.GetWithRole(userID)
		if err != nil {
			appError := errorDomain.NewAppError(err, errorDomain.ValidationError)
			_ = ctx.Error(appError)
			return
		}

		ctx.JSON(http.StatusOK, userRole)
		return

	} else {
		userRole, err := c.UserService.GetWithRole(authData.UserID)
		if err != nil {
			appError := errorDomain.NewAppError(err, errorDomain.ValidationError)
			_ = ctx.Error(appError)
			return
		}

		ctx.JSON(http.StatusOK, userRole)
		return

	}

}

// UpdateUser godoc
// @Tags user
// @Summary Get users by ID
// @Description Get Users by ID on the system
// @Param user_id path int true "id of user"
// @Security ApiKeyAuth
// @Success 200 {object} ResponseUser
// @Failure 400 {object} controllers.MessageResponse
// @Failure 500 {object} controllers.MessageResponse
// @Router /user/{user_id} [get]
func (c *Controller) UpdateUser(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		appError := errorDomain.NewAppError(errors.New("param id is necessary in the url"), errorDomain.ValidationError)
		_ = ctx.Error(appError)
		return
	}
	var request userDomain.UpdateUser

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

	user, err := c.UserService.Update(userID, request)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, user.DomainToResponseMapper())
}

// DeleteUser godoc
// @Tags user
// @Summary Get users by ID
// @Description Get Users by ID on the system
// @Param user_id path int true "id of user"
// @Security ApiKeyAuth
// @Success 200 {object} controllers.MessageResponse
// @Failure 400 {object} controllers.MessageResponse
// @Failure 500 {object} controllers.MessageResponse
// @Router /user/{user_id} [get]
func (c *Controller) DeleteUser(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		appError := errorDomain.NewAppError(errors.New("param id is necessary in the url"), errorDomain.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	err = c.UserService.Delete(userID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "resource deleted successfully"})

}

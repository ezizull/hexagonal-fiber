// Package auth contains the auth controller
package auth

import (
	useCaseAuth "hexagonal-fiber/application/usecases/auth"
	userDomain "hexagonal-fiber/domain/user"

	messageUtil "hexagonal-fiber/utils/constant/message"

	"hexagonal-fiber/infrastructure/restapi/controllers"

	"github.com/gofiber/fiber/v2"
)

// Controller is a struct that contains the auth service
type Controller struct {
	AuthService useCaseAuth.Service
}

// Login godoc
// @Tags auth
// @Summary Login UserName
// @Description Auth user by email and password
// @Param data body LoginRequest true "body data"
// @Success 200 {object} userDomain.SecurityAuthenticatedUser
// @Failure 400 {object} controllers.MessageResponse
// @Failure 500 {object} controllers.MessageResponse
// @Router /auth/login [post]
func (c *Controller) Login(ctx *fiber.Ctx) (err error) {
	var request userDomain.LoginRequest

	if err = ctx.BodyParser(&request); err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": messageUtil.ValidationError,
		})
		return
	}

	if err = controllers.Validation(request); err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
		return
	}

	authDataUser, err := c.AuthService.Login(request)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
		return
	}

	return ctx.Status(fiber.StatusOK).JSON(authDataUser)
}

// GetAccessTokenByRefreshToken godoc
// @Tags auth
// @Summary GetAccessTokenByRefreshToken UserName
// @Description Auth user by email and password
// @Param data body AccessTokenRequest true "body data"
// @Success 200 {object} userDomain.SecurityAuthenticatedUser
// @Failure 400 {object} controllers.MessageResponse
// @Failure 500 {object} controllers.MessageResponse
// @Router /auth/access-token [post]
func (c *Controller) GetAccessTokenByRefreshToken(ctx *fiber.Ctx) (err error) {
	var request userDomain.AccessTokenRequest

	if err = ctx.BodyParser(request); err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": messageUtil.ValidationError,
		})
		return
	}

	authDataUser, err := c.AuthService.AccessTokenByRefreshToken(request.RefreshToken)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
		return
	}

	return ctx.Status(fiber.StatusOK).JSON(authDataUser)
}

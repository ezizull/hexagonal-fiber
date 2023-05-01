// Package user contains the user controller
package user

import (
	"encoding/json"
	"fmt"
	useCaseUser "hexagonal-fiber/application/usecases/user"
	userDomain "hexagonal-fiber/domain/user"

	secureDomain "hexagonal-fiber/domain/security"
	redisRepo "hexagonal-fiber/infrastructure/repository/redis"

	authConst "hexagonal-fiber/utils/constant/auth"
	mssgConst "hexagonal-fiber/utils/constant/message"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

// Controller is a struct that contains the user service
type Controller struct {
	InfoRedis   *redisRepo.InfoDatabaseRedis
	UserService useCaseUser.Service
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
func (c *Controller) GetAllUsers(ctx *fiber.Ctx) (err error) {
	users, err := c.UserService.GetAll()
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
		return
	}

	return ctx.Status(fiber.StatusCreated).JSON(userDomain.ArrayDomainToResponseMapper(users))
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
func (c *Controller) GetUsersByID(ctx *fiber.Ctx) (err error) {
	authData := ctx.Locals(authConst.Authorized).(*secureDomain.Claims)

	if authData.Role == "admin" {
		var userRole *userDomain.ResponseUserRole
		userID := ctx.Params("id")

		userRole, err = c.UserService.GetWithRole(userID)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
			return
		}

		return ctx.Status(fiber.StatusOK).JSON(userRole)

	} else {
		var userRole *userDomain.ResponseUserRole

		redisDB := c.InfoRedis.NewRedis(0)
		redisData, redisErr := redisDB.Get(c.InfoRedis.CTX, ctx.IP()).Result()

		if redisErr == redis.Nil {
			userRole, err = c.UserService.GetWithRole(authData.UserID)
			if err != nil {
				ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": redisErr})
				return
			}

		} else {
			authDataUser := userDomain.SecurityAuthenticatedUser{}

			err = json.Unmarshal([]byte(redisData), &authDataUser)
			if err != nil {
				ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
				return
			}

			userRole = authDataUser.ToUserRoleResponse()
			fmt.Println("check using redis", userRole)
		}

		return ctx.Status(fiber.StatusOK).JSON(userRole)

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
func (c *Controller) UpdateUser(ctx *fiber.Ctx) (err error) {
	authData := ctx.Locals(authConst.Authorized).(*secureDomain.Claims)

	var userID string
	if authData.Role == "admin" {
		userID = ctx.Params("id")
	} else {
		userID = authData.UserID
	}

	var request userDomain.UpdateUser
	if err = ctx.BodyParser(&request); err != nil {
		appError := fiber.NewError(fiber.StatusBadRequest, mssgConst.StatusBadRequest)
		return ctx.Status(fiber.StatusBadRequest).JSON(appError)
	}

	if err = updateValidation(&request); err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
		return
	}

	user, err := c.UserService.Update(userID, request)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
		return
	}

	return ctx.Status(fiber.StatusOK).JSON(user.DomainToResponseMapper())
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
func (c *Controller) DeleteUser(ctx *fiber.Ctx) (err error) {
	authData := ctx.Locals(authConst.Authorized).(*secureDomain.Claims)

	var userID string
	if authData.Role == "admin" {
		userID = ctx.Params("id")
	} else {
		userID = authData.UserID
	}

	if err = c.UserService.Delete(userID); err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
		return
	}

	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "resource deleted successfully"})
}

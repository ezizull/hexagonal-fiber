// Package sosmed contains the sosmed controller
package sosmed

import (
	useCaseSosmed "hexagonal-fiber/application/usecases/sosmed"
	sosmedDomain "hexagonal-fiber/domain/sosmed"

	secureDomain "hexagonal-fiber/domain/security"

	authConst "hexagonal-fiber/utils/constant/auth"
	mssgConst "hexagonal-fiber/utils/constant/message"

	"github.com/gofiber/fiber/v2"
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
func (c *Controller) NewSocialMedia(ctx *fiber.Ctx) (err error) {
	authData := ctx.Locals(authConst.Authorized).(*secureDomain.Claims)

	var request sosmedDomain.NewSocialMedia
	if err := ctx.BodyParser(&request); err != nil {
		appError := fiber.NewError(fiber.StatusBadRequest, mssgConst.StatusBadRequest)
		return ctx.Status(fiber.StatusBadRequest).JSON(appError)
	}

	request.UserID = authData.UserID
	if err = createValidation(request); err != nil {
		appError := fiber.NewError(fiber.StatusBadRequest, mssgConst.ValidationError)
		return ctx.Status(fiber.StatusBadRequest).JSON(appError)
	}

	sosmed, err := c.SocialMediaService.Create(&request)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
		return
	}

	return ctx.Status(fiber.StatusCreated).JSON(sosmed)
}

// GetAllSocialMedia godoc
// @Tags sosmed
// @Summary Get all SocialMedia
// @Security ApiKeyAuth
// @Description Get all SocialMedia on the system
// @Success 200 {object} sosmedDomain.PaginationSocialMedia
// @Failure 400 {object} controllers.MessageResponse
// @Failure 500 {object} controllers.MessageResponse
// @Router /sosmed [get]
func (c *Controller) GetAllSocialMedia(ctx *fiber.Ctx) (err error) {
	page := ctx.QueryInt("page", 1)
	limit := ctx.QueryInt("limit", 20)

	sosmeds, err := c.SocialMediaService.GetAll(page, limit)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
		return
	}

	return ctx.Status(fiber.StatusOK).JSON(sosmeds)
}

// GetAllOwnSocialMedia godoc
// @Tags sosmed
// @Summary Get all SocialMedia
// @Security ApiKeyAuth
// @Description Get all SocialMedia on the system
// @Success 200 {object} sosmedDomain.PaginationSocialMedia
// @Failure 400 {object} controllers.MessageResponse
// @Failure 500 {object} controllers.MessageResponse
// @Router /sosmed [get]
func (c *Controller) GetAllOwnSocialMedia(ctx *fiber.Ctx) (err error) {
	authData := ctx.Locals(authConst.Authorized).(*secureDomain.Claims)

	page := ctx.QueryInt("page", 1)
	limit := ctx.QueryInt("limit", 20)

	sosmeds, err := c.SocialMediaService.UserGetAll(authData.UserID, page, limit)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
		return
	}

	return ctx.Status(fiber.StatusOK).JSON(sosmeds)
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
func (c *Controller) GetSocialMediaByID(ctx *fiber.Ctx) (err error) {
	sosmedID := ctx.Params("id")
	sosmed, err := c.SocialMediaService.GetByID(sosmedID)
	if err != nil {
		appError := fiber.NewError(fiber.StatusBadRequest, mssgConst.ValidationError)
		return ctx.Status(fiber.StatusBadRequest).JSON(appError)
	}

	return ctx.Status(fiber.StatusOK).JSON(sosmed)
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
func (c *Controller) UpdateSocialMedia(ctx *fiber.Ctx) (err error) {
	authData := ctx.Locals(authConst.Authorized).(*secureDomain.Claims)
	sosmedID := ctx.Params("id")

	var request sosmedDomain.UpdateSocialMedia
	if err := ctx.BodyParser(&request); err != nil {
		appError := fiber.NewError(fiber.StatusBadRequest, mssgConst.StatusBadRequest)
		return ctx.Status(fiber.StatusBadRequest).JSON(appError)
	}

	if err = updateValidation(&request); err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
		return
	}

	var sosmed *sosmedDomain.SocialMedia

	if authData.Role == "admin" {
		sosmed, err = c.SocialMediaService.Update(sosmedID, request)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
			return
		}
	} else {
		sosmed, err = c.SocialMediaService.UserUpdate(sosmedID, authData.UserID, request)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
			return
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(sosmed)
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
func (c *Controller) DeleteSocialMedia(ctx *fiber.Ctx) (err error) {
	sosmedID := ctx.Params("id")
	if err = c.SocialMediaService.Delete(sosmedID); err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
		return
	}

	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "resource deleted successfully"})

}

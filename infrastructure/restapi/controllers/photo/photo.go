// Package photo contains the photo controller
package photo

import (
	useCasePhoto "hexagonal-fiber/application/usecases/photo"
	photoDomain "hexagonal-fiber/domain/photo"

	secureDomain "hexagonal-fiber/domain/security"

	authConst "hexagonal-fiber/utils/constant/auth"
	mssgConst "hexagonal-fiber/utils/constant/message"

	"github.com/gofiber/fiber/v2"
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
func (c *Controller) NewPhoto(ctx *fiber.Ctx) (err error) {
	authData := ctx.Locals(authConst.Authorized).(*secureDomain.Claims)

	var request photoDomain.NewPhoto
	if err := ctx.BodyParser(&request); err != nil {
		appError := fiber.NewError(fiber.StatusBadRequest, mssgConst.StatusBadRequest)
		return ctx.Status(fiber.StatusBadRequest).JSON(appError)
	}

	request.UserID = authData.UserID
	if err = createValidation(request); err != nil {
		appError := fiber.NewError(fiber.StatusBadRequest, mssgConst.ValidationError)
		return ctx.Status(fiber.StatusBadRequest).JSON(appError)
	}

	photo, err := c.PhotoService.Create(&request)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
		return
	}

	return ctx.Status(fiber.StatusCreated).JSON(photo)
}

// GetAllPhotos godoc
// @Tags photo
// @Summary Get all Photos
// @Security ApiKeyAuth
// @Description Get all Photos on the system
// @Success 200 {object} photoDomain.PaginationPhoto
// @Failure 400 {object} controllers.MessageResponse
// @Failure 500 {object} controllers.MessageResponse
// @Router /photo [get]
func (c *Controller) GetAllPhotos(ctx *fiber.Ctx) (err error) {
	page := ctx.QueryInt("page", 1)
	limit := ctx.QueryInt("limit", 20)

	photos, err := c.PhotoService.GetAll(page, limit)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
		return
	}

	return ctx.Status(fiber.StatusCreated).JSON(photos)
}

// GetAllOwnPhotos godoc
// @Tags photo
// @Summary Get all Photos
// @Security ApiKeyAuth
// @Description Get all Photos on the system
// @Success 200 {object} photoDomain.PaginationPhoto
// @Failure 400 {object} controllers.MessageResponse
// @Failure 500 {object} controllers.MessageResponse
// @Router /photo [get]
func (c *Controller) GetAllOwnPhotos(ctx *fiber.Ctx) (err error) {
	authData := ctx.Locals(authConst.Authorized).(*secureDomain.Claims)

	page := ctx.QueryInt("page", 1)
	limit := ctx.QueryInt("limit", 20)

	photos, err := c.PhotoService.UserGetAll(authData.UserID, page, limit)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
		return
	}

	return ctx.Status(fiber.StatusCreated).JSON(photos)
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
func (c *Controller) GetPhotoWithComments(ctx *fiber.Ctx) (err error) {
	page := ctx.QueryInt("page", 1)
	limit := ctx.QueryInt("limit", 20)

	photoID := ctx.Params("id")
	photoComments, err := c.PhotoService.GetWithComments(photoID, page, limit)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
		return
	}

	return ctx.Status(fiber.StatusCreated).JSON(photoComments)
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
func (c *Controller) GetPhotoByID(ctx *fiber.Ctx) (err error) {
	photoID := ctx.Params("id")
	photo, err := c.PhotoService.GetByID(photoID)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
		return
	}

	return ctx.Status(fiber.StatusCreated).JSON(photo)
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
func (c *Controller) UpdatePhoto(ctx *fiber.Ctx) (err error) {
	authData := ctx.Locals(authConst.Authorized).(*secureDomain.Claims)
	photoID := ctx.Params("id")

	var request photoDomain.UpdatePhoto
	if err := ctx.BodyParser(&request); err != nil {
		appError := fiber.NewError(fiber.StatusBadRequest, mssgConst.StatusBadRequest)
		return ctx.Status(fiber.StatusBadRequest).JSON(appError)
	}

	if err = updateValidation(&request); err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
		return
	}

	var photo *photoDomain.Photo

	if authData.Role == "admin" {
		photo, err = c.PhotoService.Update(photoID, request)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
			return
		}
	} else {
		photo, err = c.PhotoService.UserUpdate(photoID, authData.UserID, request)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
			return
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(photo)
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
func (c *Controller) DeletePhoto(ctx *fiber.Ctx) (err error) {
	photoID := ctx.Params("id")
	if err = c.PhotoService.Delete(photoID); err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
		return
	}

	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "resource deleted successfully"})

}

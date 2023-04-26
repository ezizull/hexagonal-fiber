// Package photo contains the photo controller
package photo

import (
	"strconv"

	useCasePhoto "hexagonal-fiber/application/usecases/photo"
	photoDomain "hexagonal-fiber/domain/photo"

	// secureDomain "hexagonal-fiber/domain/security"

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
	UserID := ctx.Locals(authConst.AuthUserID).(int)

	var request photoDomain.NewPhoto
	if err := ctx.BodyParser(&request); err != nil {
		appError := fiber.NewError(fiber.StatusBadRequest, mssgConst.StatusBadRequest)
		return ctx.Status(fiber.StatusBadRequest).JSON(appError)
	}

	request.UserID = UserID
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
	UserID := ctx.Locals(authConst.AuthUserID).(int)

	page := ctx.QueryInt("page", 1)
	limit := ctx.QueryInt("limit", 20)

	photos, err := c.PhotoService.UserGetAll(UserID, page, limit)
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
	photoID, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		appError := fiber.NewError(fiber.StatusBadRequest, "photo id is invalid")
		return ctx.Status(fiber.StatusBadRequest).JSON(appError)
	}

	page := ctx.QueryInt("page", 1)
	limit := ctx.QueryInt("limit", 20)

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
	photoID, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		appError := fiber.NewError(fiber.StatusBadRequest, "photo id is invalid")
		return ctx.Status(fiber.StatusBadRequest).JSON(appError)
	}

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
	authRole := ctx.Locals(authConst.AuthRole).(string)
	authUserID := ctx.Locals(authConst.AuthUserID).(int)

	photoID, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		appError := fiber.NewError(fiber.StatusBadRequest, "photo id is invalid")
		return ctx.Status(fiber.StatusBadRequest).JSON(appError)
	}

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

	if authRole == "admin" {
		photo, err = c.PhotoService.Update(photoID, request)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
			return
		}
	} else {
		photo, err = c.PhotoService.UserUpdate(photoID, authUserID, request)
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
	photoID, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		appError := fiber.NewError(fiber.StatusBadRequest, "photo id is invalid")
		return ctx.Status(fiber.StatusBadRequest).JSON(appError)
	}

	if err = c.PhotoService.Delete(photoID); err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
		return
	}

	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "resource deleted successfully"})

}

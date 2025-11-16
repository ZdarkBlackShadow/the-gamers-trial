package controller

import (
	"net/url"

	"github.com/ZdarkBlackShadow/the-gamers-trial/config"
	"github.com/ZdarkBlackShadow/the-gamers-trial/service"
	"github.com/gofiber/fiber/v2"
)

type ImageController struct {
	imageService *service.ImageService
}

func InitImageController(is *service.ImageService) *ImageController {
	return &ImageController{
		imageService: is,
	}
}

func (ic *ImageController) GetImage(c *fiber.Ctx) error {
	imagePathSegment := c.Params("path_of_image")
	config.Log.Debug(imagePathSegment)

	dbPath := "/" + imagePathSegment

	decodedPath, err := url.PathUnescape(dbPath)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid path format")
	}

	fullFilePath, err := ic.imageService.GetImagePathByDBPath(decodedPath)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Image not found")
	}

	return c.SendFile(fullFilePath)
}

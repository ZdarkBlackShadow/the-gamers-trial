package routes

import (
	"github.com/ZdarkBlackShadow/the-gamers-trial/controller"
	"github.com/gofiber/fiber/v2"
)

func RegisterImageRoutes(group fiber.Router, ic *controller.ImageController) {
	group.Get("/:path_of_image", ic.GetImage)
}

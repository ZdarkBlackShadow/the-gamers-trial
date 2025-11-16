package routes

import (
	"github.com/ZdarkBlackShadow/the-gamers-trial/controller"
	"github.com/gofiber/fiber/v2"
)

func RegisterHomeRoutes(group fiber.Router, hc *controller.HomeController) {
	group.Get("/home", hc.Home)
	group.Get("/", hc.RedirectionHome)
}

package routes

import (
	"github.com/ZdarkBlackShadow/the-gamers-trial/controller"
	"github.com/gofiber/fiber/v2"
)

func RegisterAuthentificationRoutes(group fiber.Router, ac *controller.AuthentificationController) {
	group.Get("/", ac.Authentification)
	group.Post("/login", ac.Login)
	group.Post("/register", ac.Register)
	group.Post("/logout", ac.Logout)
}

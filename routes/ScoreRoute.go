package routes

import (
	"github.com/ZdarkBlackShadow/the-gamers-trial/controller"
	"github.com/gofiber/fiber/v2"
)

func RegisterScoreRoutes(group fiber.Router, sc *controller.ScoreController) {
	group.Get("/", sc.Score)
}

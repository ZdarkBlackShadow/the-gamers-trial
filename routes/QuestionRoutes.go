package routes

import (
	"github.com/ZdarkBlackShadow/the-gamers-trial/controller"
	"github.com/gofiber/fiber/v2"
)

func RegisterQuestionRoutes(group fiber.Router, qc *controller.QuestionController) {
	group.Get("", qc.Home)
	group.Post("", qc.Response)
}
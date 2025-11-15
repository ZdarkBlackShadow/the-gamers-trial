package routes

import (
	"github.com/ZdarkBlackShadow/the-gamers-trial/controller"
	"github.com/gofiber/fiber/v2"
)

func RegisterQuestionsRoutes(group fiber.Router, qc *controller.QuestionsController) {
	group.Get("/", qc.Questions)
}
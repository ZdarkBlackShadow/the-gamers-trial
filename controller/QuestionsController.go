package controller

import (
	"html/template"
	"log"

	"github.com/ZdarkBlackShadow/the-gamers-trial/service"
	"github.com/gofiber/fiber/v2"
)

type QuestionsController struct {
	template                *template.Template
	authentificationService *service.AuthentificationService
}

func InitQuestionsController(template *template.Template, authService *service.AuthentificationService) *QuestionsController {
	return &QuestionsController{
		template:                template,
		authentificationService: authService,
	}
}

func (qc *QuestionsController) Questions(c *fiber.Ctx) error {
	c.Set("Content-Type", "text/html; charset=utf-8")

	err := qc.template.ExecuteTemplate(c.Response().BodyWriter(), "questions", nil)
	if err != nil {
		log.Println("Error when rendering the template:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]string{"error": err.Error()})
	}
	return nil
}





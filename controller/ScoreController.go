package controller

import (
	"html/template"
	"log"

	"github.com/ZdarkBlackShadow/the-gamers-trial/service"
	"github.com/gofiber/fiber/v2"
)

type ScoreController struct {
	template                *template.Template
	authentificationService *service.AuthentificationService
}

func InitScoreController(template *template.Template, authService *service.AuthentificationService) *ScoreController {
	return &ScoreController{
		template:                template,
		authentificationService: authService,
	}
}

func (sc *ScoreController) Score(c *fiber.Ctx) error {
	c.Set("Content-Type", "text/html; charset=utf-8")

	err := sc.template.ExecuteTemplate(c.Response().BodyWriter(), "score", nil)
	if err != nil {
		log.Println("Error when rendering the template:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]string{"error": err.Error()})
	}
	return nil
}
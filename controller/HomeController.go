package controller

import (
	"html/template"
	"log"

	"github.com/ZdarkBlackShadow/the-gamers-trial/service"
	"github.com/gofiber/fiber/v2"
)

type HomeController struct {
	template                *template.Template
	authentificationService *service.AuthentificationService
}

func InitHomeController(template *template.Template, authService *service.AuthentificationService) *HomeController {
	return &HomeController{
		template:                template,
		authentificationService: authService,
	}
}

func (hc *HomeController) Home(c *fiber.Ctx) error {
	c.Set("Content-Type", "text/html; charset=utf-8")

	err := hc.template.ExecuteTemplate(c.Response().BodyWriter(), "home", nil)
	if err != nil {
		log.Println("Error when rendering the template:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]string{"error": err.Error()})
	}
	return nil
}

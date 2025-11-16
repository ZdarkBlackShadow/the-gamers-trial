package controller

import (
	"html/template"

	"github.com/ZdarkBlackShadow/the-gamers-trial/config"
	"github.com/ZdarkBlackShadow/the-gamers-trial/model/views"
	"github.com/ZdarkBlackShadow/the-gamers-trial/service"
	"github.com/ZdarkBlackShadow/the-gamers-trial/util"
	"github.com/gofiber/fiber/v2"
)

type ScoreController struct {
	template                *template.Template
	authentificationService *service.AuthentificationService
	scoreService            *service.ScoreService
}

func InitScoreController(template *template.Template, authService *service.AuthentificationService, scoreService *service.ScoreService) *ScoreController {
	return &ScoreController{
		template:                template,
		authentificationService: authService,
		scoreService: scoreService,
	}
}

func (sc *ScoreController) Score(c *fiber.Ctx) error {
	var data views.Score
	message, ranking, err := sc.scoreService.GetRanking()
	if err != nil {
		data.HasError = true
		data.ErrorDetail = util.CreateError(message, "error when getting ranking")
	} else {
		data.HasError = false
		data.List = ranking
	}

	c.Set("Content-Type", "text/html; charset=utf-8")

	err = sc.template.ExecuteTemplate(c.Response().BodyWriter(), "score", data)
	if err != nil {
		config.Log.Error("Error when rendering the template:", err)
		return err
	}
	return nil
}

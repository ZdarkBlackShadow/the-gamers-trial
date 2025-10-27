package controller

import (
	"html/template"

	"github.com/ZdarkBlackShadow/the-gamers-trial/config"
	"github.com/ZdarkBlackShadow/the-gamers-trial/model/entity"
	"github.com/ZdarkBlackShadow/the-gamers-trial/model/views"
	"github.com/ZdarkBlackShadow/the-gamers-trial/service"
	"github.com/ZdarkBlackShadow/the-gamers-trial/util"
	"github.com/gofiber/fiber/v2"
)

type HomeController struct {
	template       *template.Template
	sessionService *service.SessionService
}

func InitHomeController(template *template.Template, sessService *service.SessionService) *HomeController {
	return &HomeController{
		template:       template,
		sessionService: sessService,
	}
}

func (hc *HomeController) Home(c *fiber.Ctx) error {
	data := views.Home{}
	detail, err := hc.sessionService.GetInterfaceByKeys(c, []string{"user"})
	if err != nil {
		config.Log.Error("Error when read the session:" + err.Error())
		data.HasError = true
		data.ErrorDetail = util.CreateError(fiber.ErrInternalServerError.Message, "error when loading your data")
	}
	c.Set("Content-Type", "text/html; charset=utf-8")

	if detail["user"] == nil {
		data.IsLogin = false
	} else {
		data.IsLogin = true
		user, ok := detail["user"].(entity.User)
		if !ok {
			config.Log.Error("error with interface user:" + err.Error())
			data.HasError = true
			data.ErrorDetail = util.CreateError(fiber.ErrInternalServerError.Message, "error while recuperate your data")
		} else {
			data.UserInfo.Name = user.Pseudo
			data.UserInfo.Score = user.Score
		}
	}
	err = hc.template.ExecuteTemplate(c.Response().BodyWriter(), "home", data)
	if err != nil {
		config.Log.Error("Error when rendering the template:" + err.Error())
		return fiber.DefaultErrorHandler(c, err)
	}
	return nil
}

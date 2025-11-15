package controller

import (
	"html/template"

	"github.com/ZdarkBlackShadow/the-gamers-trial/config"
	"github.com/ZdarkBlackShadow/the-gamers-trial/model/request"
	"github.com/ZdarkBlackShadow/the-gamers-trial/model/views"
	"github.com/ZdarkBlackShadow/the-gamers-trial/service"
	"github.com/ZdarkBlackShadow/the-gamers-trial/util"
	"github.com/gofiber/fiber/v2"
)

type AuthentificationController struct {
	template                *template.Template
	sessionService          *service.SessionService
	authentificationService *service.AuthentificationService
}

func InitAuthentificationController(
	tpl *template.Template,
	sessService *service.SessionService,
	authService *service.AuthentificationService) *AuthentificationController {
	return &AuthentificationController{
		template:                tpl,
		sessionService:          sessService,
		authentificationService: authService,
	}
}

func (ac *AuthentificationController) Authentification(c *fiber.Ctx) error {
	sess, err := ac.sessionService.GetSession(c)
	if err != nil {
		return fiber.DefaultErrorHandler(c, err)
	}

	var authStruct views.Authentification
	potentialErr := sess.Get("error")

	if potentialErr != nil {
		if errStruct, ok := potentialErr.(views.Error); ok {
			authStruct.ErrorDetail = errStruct
			authStruct.HasError = true
			//err = ac.sessionService.DeleteKey(sess, "error")
		} else {
			config.Log.Error("Error type in session is not views.Authentification")
		}
	} else {
		authStruct.HasError = false
	}
	config.Log.Debug(authStruct)
	c.Set("Content-Type", "text/html; charset=utf-8")

	err = ac.template.ExecuteTemplate(c.Response().BodyWriter(), "authentification", authStruct)
	if err != nil {
		config.Log.Error("Error when rendering the template:", err)
		return fiber.DefaultErrorHandler(c, err)
	}
	return nil

}

func (ac *AuthentificationController) Register(c *fiber.Ctx) error {
	var userRequest request.Register
	err := c.BodyParser(&userRequest)
	if err != nil {
		config.Log.Debug("bad request from " + c.IP() + ":" + err.Error())
		errStruct := util.CreateError(fiber.ErrBadRequest.Message, "the json is incorrect")
		err = ac.sessionService.CreateSession(c, map[string]interface{}{"error": errStruct})
		if err != nil {
			config.Log.Error("error when trying to save the session : " + err.Error())
			return fiber.DefaultErrorHandler(c, err)
		}
		return c.Redirect("/authentification", fiber.StatusSeeOther)
	}

	hasError, errorDetail, newUser := ac.authentificationService.Register(userRequest)
	if hasError {
		err = ac.sessionService.CreateSession(c, map[string]interface{}{"error": errorDetail})
		if err != nil {
			config.Log.Error("error when trying to save the session : " + err.Error())
			return fiber.DefaultErrorHandler(c, err)
		}
		return c.Redirect("/authentification", fiber.StatusSeeOther)
	}
	config.Log.Debug("test")

	err = ac.sessionService.CreateSession(c, map[string]interface{}{"user": newUser})
	if err != nil {
		return fiber.DefaultErrorHandler(c, err)
	}
	return c.Redirect("/home", fiber.StatusSeeOther)
}

func (ac *AuthentificationController) Login(c *fiber.Ctx) error {
	var userRequest request.Login
	err := c.BodyParser(&userRequest)
	if err != nil {
		config.Log.Debug("bad request from " + c.IP() + ":" + err.Error())
		errStruct := util.CreateError(fiber.ErrBadGateway.Message, "the json is incorect")
		err = ac.sessionService.CreateSession(c, map[string]interface{}{"error": errStruct})
		if err != nil {
			return fiber.DefaultErrorHandler(c, err)
		}
		return c.Redirect("/authentification", fiber.StatusSeeOther)
	}

	hasError, errorDetail, user := ac.authentificationService.Login(userRequest)
	if hasError {
		err = ac.sessionService.CreateSession(c, map[string]interface{}{"error": errorDetail})
		if err != nil {
			return fiber.DefaultErrorHandler(c, err)
		}
		return c.Redirect("/authentification", fiber.StatusSeeOther)
	}

	err = ac.sessionService.CreateSession(c, map[string]interface{}{"user": user})
	if err != nil {
		return fiber.DefaultErrorHandler(c, err)
	}
	return c.Redirect("/home", fiber.StatusSeeOther)
}

func (ac *AuthentificationController) Logout(c *fiber.Ctx) error {
    sess, err := ac.sessionService.GetSession(c)
    if err != nil {
        config.Log.Error("Error retrieving session: ", err)
        return fiber.DefaultErrorHandler(c, err)
    }

    err = ac.sessionService.DeleteKey(sess, "user")
    if err != nil {
        config.Log.Error("Error deleting user from session: ", err)
        return fiber.DefaultErrorHandler(c, err)
    }
    return c.Redirect("/home", fiber.StatusSeeOther)
}

package controller

import (
	"html/template"

	"github.com/ZdarkBlackShadow/the-gamers-trial/config"
	"github.com/ZdarkBlackShadow/the-gamers-trial/model/entity"
	"github.com/ZdarkBlackShadow/the-gamers-trial/model/request"
	"github.com/ZdarkBlackShadow/the-gamers-trial/model/session"
	"github.com/ZdarkBlackShadow/the-gamers-trial/model/views"
	"github.com/ZdarkBlackShadow/the-gamers-trial/service"
	"github.com/ZdarkBlackShadow/the-gamers-trial/util"
	"github.com/gofiber/fiber/v2"
)

type QuestionController struct {
	questionService *service.QuestionService
	sessionService  *service.SessionService
	template        *template.Template
}

func InitQuestionController(qs *service.QuestionService, sessService *service.SessionService, template *template.Template) *QuestionController {
	return &QuestionController{
		questionService: qs,
		sessionService:  sessService,
		template:        template,
	}
}

func (qc *QuestionController) Response(c *fiber.Ctx) error {
	var req request.Question
	err := c.BodyParser(&req)
	if err != nil {
		return nil
	}
	details, err := qc.sessionService.GetInterfaceByKeys(c, []string{"user"})
	if err != nil {
		config.Log.Error("error when loading session:" + err.Error())
		return fiber.DefaultErrorHandler(c, err)
	}
	if details["user"] == nil {
		return c.Redirect("/authentification", fiber.StatusSeeOther)
	}
	user, ok := details["user"].(entity.User)
	if !ok {
		return c.Redirect("/authentification", fiber.StatusSeeOther)
	}
	message, sessionAnswer, updatedUser, err := qc.questionService.ResponseAnalysis(req, user)
	if err != nil {
		config.Log.Error(message + ":" + err.Error())
		return fiber.DefaultErrorHandler(c, err)
	}
	err = qc.sessionService.CreateSession(c, map[string]interface{}{"user": updatedUser, "user_answer": sessionAnswer})
	if err != nil {
		return fiber.DefaultErrorHandler(c, err)
	}
	return c.Redirect("/question", fiber.StatusSeeOther)
}

func (qc *QuestionController) Home(c *fiber.Ctx) error {
	data := views.Question{}
	details, err := qc.sessionService.GetInterfaceByKeys(c, []string{"state", "user_answer"})
	if err != nil {
		config.Log.Error("Error when read the session:" + err.Error())
		data.HasError = true
		data.ErrorDetail = util.CreateError(fiber.ErrInternalServerError.Message, "error when loading your data")
	} else {
		if details["state"] == nil {
			data.State = true
		} else {
			state, ok := details["state"].(bool)
			if !ok {
				config.Log.Error("error when loading interface state:")
				data.HasError = true
				data.ErrorDetail = util.CreateError(fiber.ErrInternalServerError.Message, "error when recuperate your data")
			}
			data.State = state
		}
		if data.State {
			options, question, err := qc.questionService.GetRandomQuestion()
			if err != nil {
				data.HasError = true
				data.ErrorDetail = util.CreateError(fiber.ErrInternalServerError.Message, "error in the question service")
			} else {
				data.ID = question.ID
				data.Content = question.Content
				for _, option := range options {
					data.Options = append(data.Options, views.Option{ID: option.ID, Text: option.Text})
				}
				if question.Image.Path != "" {
					data.HasImage = true
					// Remove leading slash if present, as the route /image/:path_of_image will handle it
					imagePath := question.Image.Path
					if len(imagePath) > 0 && imagePath[0] == '/' {
						imagePath = imagePath[1:]
					}
					data.ImagePath = imagePath
				} else {
					data.HasImage = false
				}
			}
		} else if !data.State {
			userAnswer, ok := details["user_answer"].(session.UserAnswer)
			if !ok {
				data.HasError = true
				data.ErrorDetail = util.CreateError(fiber.ErrInternalServerError.Message, "error when recuperate user answer")
			} else {
				data.NewUserScore = userAnswer.NewUserScore
				data.UserHasCorrect = userAnswer.UserHasCorrect
				data.OldUserScore = userAnswer.OldUserScore
			}
		}
	}
	c.Set("Content-Type", "text/html; charset=utf-8")
	err = qc.template.ExecuteTemplate(c.Response().BodyWriter(), "question", data)
	if err != nil {
		config.Log.Error("Error when rendering the template:" + err.Error())
		return fiber.DefaultErrorHandler(c, err)
	}
	return nil
}

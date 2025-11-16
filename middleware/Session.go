package middleware

import (
	"github.com/ZdarkBlackShadow/the-gamers-trial/service"
	"github.com/gofiber/fiber/v2"
)

func SessionMiddleware(sessionService *service.SessionService) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		sess, err := sessionService.GetSession(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		sess.Save()
		return c.Next()
	}
}

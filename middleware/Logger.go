package middleware

import (
	"time"

	"github.com/ZdarkBlackShadow/the-gamers-trial/config"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func Logger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		err := c.Next()
		stop := time.Now()
		latency := stop.Sub(start)

		status := c.Response().StatusCode()

		entry := config.Log.WithFields(logrus.Fields{
			"method":  c.Method(),
			"path":    c.Path(),
			"status":  status,
			"latency": latency,
			"ip":      c.IP(),
		})

		switch {
		case status >= 500:
			entry.Error("Internal server error")
		case status >= 400:
			entry.Warn("Client error")
		default:
			entry.Info("Request handled")
		}

		return err
	}
}

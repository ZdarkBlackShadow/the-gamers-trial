package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type RateLimiter struct {
	visits    map[string]int
	limit     int
	window    time.Duration
	resetTime time.Time
}

func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		visits:    make(map[string]int),
		limit:     limit,
		window:    window,
		resetTime: time.Now().Add(window),
	}
}

func (rl *RateLimiter) Handle() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ip := c.IP()

		if time.Now().After(rl.resetTime) {
			rl.visits = make(map[string]int)
			rl.resetTime = time.Now().Add(rl.window)
		}

		rl.visits[ip]++

		if rl.visits[ip] > rl.limit {
			return c.Status(fiber.StatusTooManyRequests).JSON(map[string]string{"error": "too many request, try again later"})
		}

		return c.Next()
	}
}

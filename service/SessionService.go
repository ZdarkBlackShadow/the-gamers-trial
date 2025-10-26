package service

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type SessionService struct {
	sessionStore *session.Store
}

func InitSessionService() *SessionService {
	store := session.New(session.Config{
		Expiration:        24 * time.Hour,
		KeyLookup:         "cookie:session_the_gamers_trial",
		CookieDomain:      "",
		CookiePath:        "",
		CookieSecure:      false,
		CookieHTTPOnly:    true,
		CookieSameSite:    "Lax",
		CookieSessionOnly: false,
	})

	return &SessionService{
		sessionStore: store,
	}
}

func (s *SessionService) GetSession(c *fiber.Ctx) (*session.Session, error) {
	return s.sessionStore.Get(c)
}

func (s *SessionService) CreateSession(c *fiber.Ctx, data map[string]interface{}) error {
	session, err := s.sessionStore.Get(c)
	if err != nil {
		return err
	}

	for key, value := range data {
		session.Set(key, value)
	}

	return session.Save()
}

func (s *SessionService) DeleteKey(sess *session.Session, key string) error {
	sess.Delete(key)
	return sess.Save()
}
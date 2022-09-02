package middleware

import (
	"OctaneServer/config"
	"OctaneServer/definitions"

	"github.com/gofiber/fiber/v2"
)

func New() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if len(config.CurrentConfig.Token) == 0 {
			return c.Next()
		}
		token := c.Get(definitions.HeaderToken)
		if token == "" {
			c.Status(404)
			return nil
		}
		for _, v := range config.CurrentConfig.Token {
			if v == token {
				return c.Next()
			}
		}
		c.Status(401)
		return nil
	}
}

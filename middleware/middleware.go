package middleware

import (
	"OctaneServer/config"
	"OctaneServer/data"
	"OctaneServer/definitions"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func New() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if config.CurrentConfig.Debug {
			log.Debug().Str("body", c.Request().String()).Msg("Got request")
		}
		c.Response().Header.Add("Access-Control-Allow-Origin", "*")
		c.Response().Header.Add("Access-Control-Allow-Headers", definitions.HeaderToken+","+"Content-Type")
		if string(c.Request().Header.Method()) == "OPTIONS" {
			c.Status(200).Send(nil)
			return nil
		}
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
		c.Status(401).JSON(data.Error{
			Code:   definitions.ERR_UNAUTHORIZED,
			Reason: config.Msg[config.CurrentConfig.Lang].API.Error.Unauthorized,
		})
		return nil
	}
}

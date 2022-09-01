package handlers

import (
	"OctaneServer/data"
	"OctaneServer/definitions"

	"github.com/gofiber/fiber/v2"
)

const Health = "health"

func HealthGet(ctx *fiber.Ctx) error {
	return ctx.JSON(data.Health{Health: definitions.Healthy})
}

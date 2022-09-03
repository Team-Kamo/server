package handlers

import (
	"OctaneServer/data"
	"OctaneServer/definitions"

	"github.com/gofiber/fiber/v2"
)

const Health = "health"

// HealthGet godoc
// @Summary     Get health status
// @Description get health status of the api server
// @Tags        health
// @Accept      json,xml,x-www-form-urlencoded
// @Produce     json
// @Success     200 {object} data.Health
// @Failure     401 {object} data.Error
// @Failure     500 {object} data.Error
// @Router      /health [get]
func HealthGet(ctx *fiber.Ctx) error {
	return ctx.JSON(data.Health{Health: definitions.Healthy})
}

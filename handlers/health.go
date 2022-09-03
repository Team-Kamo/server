package handlers

import (
	"OctaneServer/config"
	"OctaneServer/data"
	"OctaneServer/definitions"
	"strconv"

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
	health := data.Health{Health: definitions.Healthy}
	//ToDo: Implement more smarter
	if exceptionsCount > 50 {
		health.Health = definitions.Faulty
		health.Message = config.Msg[config.CurrentConfig.Lang].API.HealthTooManyError + strconv.Itoa(exceptionsCount)
	} else if exceptionsCount > 10 {
		health.Health = definitions.Degraded
		health.Message = config.Msg[config.CurrentConfig.Lang].API.HealthTooManyError + strconv.Itoa(exceptionsCount)
	}
	return ctx.JSON(health)
}

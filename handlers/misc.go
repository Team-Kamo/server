package handlers

import (
	"OctaneServer/config"
	"OctaneServer/data"
	"OctaneServer/definitions"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

var exceptionsCount = 0

func internalException(ctx *fiber.Ctx, err error) {
	errObj := data.Error{
		Code:   definitions.ERR_INTERNAL_EXCEPTION,
		Reason: err.Error(),
	}
	err = ctx.Status(500).JSON(errObj)
	if err != nil {
		log.Error().Err(err).Msg(config.Msg[config.CurrentConfig.Lang].Console.Error.FiberResponse)
	}
	exceptionsCount++
}

func HouseKeeping() {
	exceptionsCount = 0
}

func idConvert(ctx *fiber.Ctx) int64 {
	id, err := strconv.ParseInt(ctx.Params("*"), 10, 64)
	if err != nil {
		errObj := data.Error{
			Code:   definitions.ERR_BAD_REQUEST,
			Reason: err.Error(),
		}
		err = ctx.Status(400).JSON(errObj)
		if err != nil {
			log.Error().Err(err).Msg(config.Msg[config.CurrentConfig.Lang].Console.Error.FiberResponse)
		}
		return -1
	}
	return id
}

func readBody(ctx *fiber.Ctx, out interface{}) bool {
	err := ctx.BodyParser(&out)
	if err != nil {
		badRequest(ctx, err.Error())
		return false
	}
	return true
}

func badRequest(ctx *fiber.Ctx, why string) {
	errObj := data.Error{
		Code:   definitions.ERR_BAD_REQUEST,
		Reason: why,
	}
	err := ctx.Status(400).JSON(errObj)
	if err != nil {
		log.Error().Err(err).Msg(config.Msg[config.CurrentConfig.Lang].Console.Error.FiberResponse)
	}
}

func genericError(ctx *fiber.Ctx, status int, code string, why string) {
	errObj := data.Error{
		Code:   code,
		Reason: why,
	}
	err := ctx.Status(status).JSON(errObj)
	if err != nil {
		log.Error().Err(err).Msg(config.Msg[config.CurrentConfig.Lang].Console.Error.FiberResponse)
	}
}

func entryCheck(ctx *fiber.Ctx, err error, code string) bool {
	if err == data.ErrNoEntry {
		errObj := data.Error{
			Code:   code,
			Reason: err.Error(),
		}
		ctx.Status(400).JSON(errObj)
		return false
	} else if err != nil {
		internalException(ctx, err)
		return false
	}
	return true
}

package handlers

import (
	"OctaneServer/config"
	"OctaneServer/data"
	"OctaneServer/definitions"

	"github.com/gofiber/fiber/v2"
)

const Status = "room/*/status"

func StatusGet(ctx *fiber.Ctx) error {
	id := idConvert(ctx)
	if id == -1 {
		return nil
	}
	status, err := data.GetStatus(id)
	ok := entryCheck(ctx, err, definitions.ERR_NO_STATUS_SET)
	if !ok {
		return nil
	}
	return ctx.JSON(status)
}

func StatusPut(ctx *fiber.Ctx) error {
	id := idConvert(ctx)
	if id == -1 {
		return nil
	}
	var req = data.Status{}
	ok := readBody(ctx, &req)
	if !ok {
		return nil
	}
	room, err := data.GetRoom(id)
	ok = entryCheck(ctx, err, definitions.ERR_NO_SUCH_ROOM)
	if !ok {
		return nil
	}
	exists := false
	for _, v := range room.Devices {
		if req.Device == v.Name {
			exists = true
			break
		}
	}
	if !exists {
		genericError(ctx, 400, definitions.ERR_DEVICE_NOT_CONNECTED, config.Msg[config.CurrentConfig.Lang].API.Error.DeviceNotConnected)
		return nil
	}
	// BLAKE2b-64
	if len(req.Hash) != 64 {
		genericError(ctx, 400, definitions.ERR_BAD_HASH, config.Msg[config.CurrentConfig.Lang].API.Error.BadHash)
		return nil
	}
	err = data.SetStatus(id, req)
	ok = entryCheck(ctx, err, definitions.ERR_NO_SUCH_ROOM)
	if !ok {
		return nil
	}
	return ctx.Send(nil)
}

func StatusDelete(ctx *fiber.Ctx) error {
	id := idConvert(ctx)
	if id == -1 {
		return nil
	}
	err := data.DeleteStatus(id)
	ok := entryCheck(ctx, err, definitions.ERR_NO_STATUS_SET)
	if !ok {
		return nil
	}
	return ctx.Send(nil)
}

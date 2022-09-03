package handlers

import (
	"OctaneServer/config"
	"OctaneServer/data"
	"OctaneServer/definitions"

	"github.com/gofiber/fiber/v2"
)

const Status = "room/*/status"

// StatusGet godoc
// @Summary     Get content status
// @Description Get content status of the room
// @Tags        status
// @Accept      json,xml,x-www-form-urlencoded,mpfd
// @Produce     json
// @Param       id  path     int true "Room ID"
// @Success     200 {object} data.Status
// @Failure     401 {object} data.Error
// @Failure     500 {object} data.Error
// @Router      /room/{id}/status [get]
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

// StatusPut godoc
// @Summary     Set content status
// @Description Set content status of the room
// @Tags        status
// @Accept      json,xml,x-www-form-urlencoded,mpfd
// @Param       Status body data.Status true "Room status"
// @Param       id     path int         true "Room ID"
// @Success     200
// @Failure     401 {object} data.Error
// @Failure     500 {object} data.Error
// @Router      /room/{id}/status [put]
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
	// BLAKE2b-256
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

// StatusDelete godoc
// @Summary     Delete content status
// @Description Delete content status of the room
// @Tags        status
// @Accept      json,xml,x-www-form-urlencoded,mpfd
// @Param       id path int true "Room ID"
// @Success     200
// @Failure     401 {object} data.Error
// @Failure     500 {object} data.Error
// @Router      /room/{id}/status [delete]
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

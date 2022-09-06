package handlers

import (
	"OctaneServer/config"
	"OctaneServer/data"
	"OctaneServer/definitions"

	"github.com/gofiber/fiber/v2"
)

var RoomID = "room/*/"

// RoomIDGet godoc
// @Summary     Get room status
// @Description Get status of the room
// @Tags        room
// @Accept      json,xml,x-www-form-urlencoded,mpfd
// @Produce     json
// @Success     200 {object} data.Room
// @Failure     401 {object} data.Error
// @Failure     500 {object} data.Error
// @Router      /room/{id} [get]
func RoomIDGet(ctx *fiber.Ctx) error {
	id := idConvert(ctx)
	if id == -1 {
		return nil
	}
	room, err := data.GetRoom(id)
	ok := entryCheck(ctx, err, definitions.ERR_NO_SUCH_ROOM)
	if !ok {
		return nil
	}
	return ctx.JSON(room)
}

// RoomIDDelete godoc
// @Summary     Delete room
// @Description Delete room
// @Tags        room
// @Accept      json,xml,x-www-form-urlencoded,mpfd
// @Param       id                 path int                     true "Room ID"
// @Success     200
// @Failure     401 {object} data.Error
// @Failure     500 {object} data.Error
// @Router      /room/{id} [delete]
func RoomIDDelete(ctx *fiber.Ctx) error {
	id := idConvert(ctx)
	if id == -1 {
		return nil
	}
	err := data.DeleteRoom(id)
	ok := entryCheck(ctx, err, definitions.ERR_NO_SUCH_ROOM)
	if !ok {
		return nil
	}
	return ctx.Send(nil)
}

// RoomIDPost godoc
// @Summary     Connect or disconnect to/from the room
// @Description Connect or disconnect device to/from the room
// @Tags        room
// @Accept      json,xml,x-www-form-urlencoded,mpfd
// @Param       RoomConnectRequest body data.RoomConnectRequest true "Room connect request"
// @Param       id path int true "Room ID"
// @Success     200
// @Failure     401 {object} data.Error
// @Failure     500 {object} data.Error
// @Router      /room/{id} [post]
func RoomIdPost(ctx *fiber.Ctx) error {
	id := idConvert(ctx)
	if id == -1 {
		return nil
	}
	var req = data.RoomConnectRequest{}
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
		if v.Name == req.Name {
			exists = true
			break
		}
	}
	if req.Request == definitions.RequestConnect {
		if exists {
			genericError(ctx, 400, definitions.ERR_DUP_DEVICE, config.Msg[config.CurrentConfig.Lang].API.Error.DupDevice)
			return nil
		} else {
			err = data.ConnectRoom(req.Name, id)
			ok = entryCheck(ctx, err, definitions.ERR_NO_SUCH_ROOM)
			if !ok {
				return nil
			}
			return nil
		}
	} else if req.Request == definitions.RequestDisconnect {
		if exists {
			err = data.DisconnectRoom(req.Name, id)
			ok = entryCheck(ctx, err, definitions.ERR_NO_SUCH_ROOM)
			if !ok {
				return nil
			}
			return nil
		} else {
			genericError(ctx, 400, definitions.ERR_DEVICE_NOT_CONNECTED, config.Msg[config.CurrentConfig.Lang].API.Error.DeviceNotConnected)
			return nil
		}
	}
	genericError(ctx, 400, definitions.ERR_BAD_CONNECTION_REQUEST, config.Msg[config.CurrentConfig.Lang].API.Error.ConnectionRequest)
	return nil
}

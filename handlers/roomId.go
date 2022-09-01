package handlers

import (
	"OctaneServer/config"
	"OctaneServer/data"
	"OctaneServer/definitions"

	"github.com/gofiber/fiber/v2"
)

var RoomID = "room/*/"

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
	if exists {
		genericError(ctx, 400, definitions.ERR_DUP_DEVICE, config.Msg[config.CurrentConfig.Lang].API.Error.DupDevice)
		return nil
	}
	err = data.ConnectRoom(req.Name, id)
	ok = entryCheck(ctx, err, definitions.ERR_NO_SUCH_ROOM)
	if !ok {
		return nil
	}
	return nil
}

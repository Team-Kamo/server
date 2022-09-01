package handlers

import (
	"OctaneServer/config"
	"OctaneServer/data"

	"github.com/gofiber/fiber/v2"
)

const Room = "room"

func RoomPost(ctx *fiber.Ctx) error {
	req := data.RoomCreateRequest{}
	ok := readBody(ctx, &req)
	if !ok {
		return nil
	}
	if req.Name == "" {
		badRequest(ctx, config.Msg[config.CurrentConfig.Lang].API.Error.EmptyRoomName)
		return nil
	}

	id, err := data.NewRoom(req.Name)
	if err != nil {
		internalException(ctx, err)
		return nil
	}
	res := data.RoomCreate{
		Id: id,
	}
	return ctx.JSON(res)
}

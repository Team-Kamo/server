package handlers

import (
	"OctaneServer/config"
	"OctaneServer/data"

	"github.com/gofiber/fiber/v2"
)

const Room = "room"

// RoomPost godoc
// @Summary     Create room
// @Description Create new room
// @Tags        create
// @Accept      json,xml,x-www-form-urlencoded,mpfd
// @Param       RoomCreateRequest body     data.RoomCreateRequest true "Room create request"
// @Success     200               {object} data.RoomCreate
// @Failure     401               {object} data.Error
// @Failure     500               {object} data.Error
// @Router      /room [post]
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

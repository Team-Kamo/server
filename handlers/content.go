package handlers

import (
	"OctaneServer/config"
	"OctaneServer/data"
	"OctaneServer/definitions"
	"OctaneServer/storage"
	"encoding/hex"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/blake2b"
)

const Content = "room/*/content"

// ContentPut godoc
// @Summary      Set room content
// @Description  Set content to the room
// @Tags         content
// @Param        id   path      int  true  "Room ID"
// @Success      200
// @Failure      401  {object}  data.Error
// @Failure      500  {object}  data.Error
// @Router       /room/{id}/content [put]
func ContentPut(ctx *fiber.Ctx) error {
	body := ctx.Body()
	id := idConvert(ctx)
	if id == -1 {
		return nil
	}
	status, err := data.GetStatus(id)
	ok := entryCheck(ctx, err, definitions.ERR_NO_STATUS_SET)
	if !ok {
		return nil
	}
	sum := blake2b.Sum256(body)
	if hex.EncodeToString(sum[:]) != status.Hash {
		genericError(ctx, 400, definitions.ERR_CONTENT_MISMATCH, config.Msg[config.CurrentConfig.Lang].API.Error.ContentMismatch)
		return nil
	}
	mime := string(ctx.Request().Header.ContentType())
	if mime != status.Mime {
		genericError(ctx, 400, definitions.ERR_CONTENT_MISMATCH, config.Msg[config.CurrentConfig.Lang].API.Error.MimeMismatch)
		return nil
	}
	_, err = data.GetRoom(id)
	if !entryCheck(ctx, err, definitions.ERR_BAD_REQUEST) {
		return nil
	}
	if status.Type == definitions.Clipboard {
		err = data.SetTextContent(id, string(body))
		if err != nil {
			internalException(ctx, err)
			return nil
		}
	} else {
		err = storage.Storage.Save(strconv.FormatInt(id, 10)+"_"+status.Hash, body)
		if err != nil {
			internalException(ctx, err)
			return nil
		}
	}
	return ctx.Send(nil)
}

// ContentDelete godoc
// @Summary      Delete room content
// @Description  Delete content to the room
// @Tags         content
// @Param        id   path      int  true  "Room ID"
// @Success      200
// @Failure      401  {object}  data.Error
// @Failure      500  {object}  data.Error
// @Router       /room/{id}/content [delete]
func ContentDelete(ctx *fiber.Ctx) error {
	id := idConvert(ctx)
	if id == -1 {
		return nil
	}
	status, err := data.GetStatus(id)
	ok := entryCheck(ctx, err, definitions.ERR_NO_STATUS_SET)
	if !ok {
		return nil
	}
	_, err = data.GetRoom(id)
	if !entryCheck(ctx, err, definitions.ERR_BAD_REQUEST) {
		return nil
	}
	if status.Type == definitions.Clipboard {
		err = data.DeleteTextContent(id)
		if err != nil {
			internalException(ctx, err)
			return nil
		}
	} else {
		err = storage.Storage.Delete(strconv.FormatInt(id, 10) + "_" + status.Hash)
		if err != nil {
			internalException(ctx, err)
			return nil
		}
	}
	return ctx.Send(nil)
}

// ContentGet godoc
// @Summary      Get room content
// @Description  Get content to the room
// @Tags         content
// @Param        id   path      int  true  "Room ID"
// @Success      200
// @Failure      401  {object}  data.Error
// @Failure      500  {object}  data.Error
// @Router       /room/{id}/content [get]
func ContentGet(ctx *fiber.Ctx) error {
	id := idConvert(ctx)
	if id == -1 {
		return nil
	}
	status, err := data.GetStatus(id)
	ok := entryCheck(ctx, err, definitions.ERR_NO_STATUS_SET)
	if !ok {
		return nil
	}
	_, err = data.GetRoom(id)
	if !entryCheck(ctx, err, definitions.ERR_BAD_REQUEST) {
		return nil
	}
	if status.Type == definitions.Clipboard {
		text, err := data.GetTextContent(id)
		if err != nil {
			internalException(ctx, err)
			return nil
		}
		ctx.Set("Content-Type", status.Mime)
		return ctx.Send([]byte(text))
	} else {
		return ctx.Redirect(storage.Uri(strconv.FormatInt(id, 10)+"_"+status.Hash), 302)
	}
}

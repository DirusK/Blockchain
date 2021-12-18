package blockchain

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// respond - helper func for respond.
func (h Handler) respond(ctx *fiber.Ctx, code int, payload interface{}) error {
	var err error

	ctx.Response().SetStatusCode(code)

	if err = ctx.JSON(payload); err != nil {
		log.Println("failed: write response: ", err)
	}

	return err
}

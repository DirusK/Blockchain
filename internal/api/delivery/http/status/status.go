package status

import (
	"github.com/gofiber/fiber/v2"
)

type (
	// Handler defines a handler for HTTP requests for checking status.
	Handler struct {
		resp response
	}
)

// NewHandler defines a handler constructor.
func NewHandler(appName, version, nodeID string) *Handler {
	return &Handler{
		resp: response{
			App:     appName,
			Version: version,
			NodeID:  nodeID,
			Status:  fiber.StatusOK,
		},
	}
}

// CheckStatus -  HTTP GET handler for status endpoint.
func (h Handler) CheckStatus(ctx *fiber.Ctx) error {
	return ctx.JSON(h.resp)
}

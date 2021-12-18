package delivery

import (
	"github.com/gofiber/fiber/v2"
)

type (
	// StatusHTTP - describes an interface for work with service status over HTTP.
	StatusHTTP interface {
		CheckStatus(ctx *fiber.Ctx) error
	}

	// BlockchainHTTP - describes an interface for work with blockchain over HTTP.
	BlockchainHTTP interface {
		Mine(ctx *fiber.Ctx) error
		CreateTransaction(ctx *fiber.Ctx) error
		Explore(ctx *fiber.Ctx) error
		NodesRegister(ctx *fiber.Ctx) error
		NodesResolve(ctx *fiber.Ctx) error
	}
)

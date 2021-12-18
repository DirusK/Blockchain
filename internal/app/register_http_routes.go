package app

import (
	"blockchain/internal/api/delivery/http/blockchain"
	"blockchain/internal/api/delivery/http/status"

	"github.com/gofiber/fiber/v2"
)

func (t *App) registerHTTPRoutes(app *fiber.App) {
	API := app.Group("/v1/blockchain")
	API.Get("/status", t.statusHTTPHandler.CheckStatus)
	API.Get("/explore", t.blockchainHTTPHandler.Explore)
	API.Get("/mine", t.blockchainHTTPHandler.Mine)

	transactionsGroup := API.Group("/transactions")
	transactionsGroup.Post("/create", t.blockchainHTTPHandler.CreateTransaction)

	nodesGroup := API.Group("/nodes")
	nodesGroup.Get("/resolve", t.blockchainHTTPHandler.NodesResolve)
	nodesGroup.Post("/register", t.blockchainHTTPHandler.NodesRegister)
}

func registerHTTPHandlers(a *App) {
	a.statusHTTPHandler = status.NewHandler(
		a.meta.Info.AppName,
		a.meta.Info.Version,
		a.nodeID,
	)

	a.blockchainHTTPHandler = blockchain.NewHandler(a.blockchain, a.nodeID)
}

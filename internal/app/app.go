package app

import (
	"blockchain/pkg/blockchain"
	"context"
	"log"
	"sync"

	"blockchain/internal/api/delivery"
	"blockchain/internal/config"
)

type (
	// Meta defines meta for application.
	Meta struct {
		Info       Info
		ConfigPath string
	}

	// Info defines metadata of application.
	Info struct {
		AppName string
		Version string
		Status  string
	}

	// App defines main application struct.
	App struct {
		nodeID string

		// meta information about application.
		meta Meta

		// tech dependencies
		config     *config.Config
		blockchain *blockchain.Blockchain

		// delivery dependencies.
		statusHTTPHandler     delivery.StatusHTTP
		blockchainHTTPHandler delivery.BlockchainHTTP
	}

	worker func(ctx context.Context, a *App)
)

// New - app constructor without init for components.
func New(meta Meta, nodeID string) *App {
	return &App{
		meta:   meta,
		nodeID: nodeID,
	}
}

// Run registers graceful shutdown.
// populate configuration and application dependencies.
// run workers.
func (t *App) Run(ctx context.Context) {
	// initialize configuration
	t.populateConfiguration()

	// init dependencies
	t.registerDependencies(
		initBlockchain,
		registerHTTPHandlers,
	)

	t.runWorkers(ctx, t.initWorkers())
}

// PopulateConfiguration ...
func (t *App) populateConfiguration() {
	var err error

	if t.config, err = config.New(t.meta.ConfigPath); err != nil {
		log.Fatalln(err)
	}
}

// RunWorkers ...
func (t *App) runWorkers(ctx context.Context, workers []worker) {
	wg := new(sync.WaitGroup)
	wg.Add(len(workers))

	for _, work := range workers {
		go func(ctx context.Context, work func(context.Context, *App), t *App) {
			work(ctx, t)
			wg.Done()
		}(ctx, work, t)
	}

	wg.Wait()
}

// RegisterDependencies ...
func (t *App) registerDependencies(deps ...func(a *App)) {
	for _, fn := range deps {
		fn(t)
	}
}

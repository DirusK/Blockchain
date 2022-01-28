package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"blockchain/internal/app"
	"blockchain/internal/config"

	"github.com/google/uuid"
)

const (
	appName = "blockchain"
	version = "v1.0.0"
)

func main() {
	cfgPath := flag.String("c", config.DefaultPath, "configuration file")
	flag.Parse()

	nodeID, err := uuid.NewUUID()
	if err != nil {
		log.Fatalln(err)
	}

	app.New(
		app.Meta{
			Info: app.Info{
				AppName: appName,
				Version: version,
			},
			ConfigPath: *cfgPath,
		}, nodeID.String(),
	).Run(registerGracefulHandle())

}

func registerGracefulHandle() context.Context {
	ctx, cancel := context.WithCancel(context.Background())

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		cancel()
	}()

	return ctx
}

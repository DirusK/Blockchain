package app

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func serveHTTP(ctx context.Context, a *App) {
	router := fiber.New(fiber.Config{
		Prefork:      false,
		ReadTimeout:  a.config.HTTPServer.ReadTimeout,
		WriteTimeout: a.config.HTTPServer.WriteTimeout,
		Network:      "tcp4",
		BodyLimit:    a.config.HTTPServer.HeaderLimit,
		AppName:      a.meta.Info.AppName,
	})

	router.Use(requestid.New())
	router.Use(loggerMiddleware(a))
	router.Use(recover.New())

	a.registerHTTPRoutes(router)

	// graceful shutdown listener.
	go func() {
		<-ctx.Done()

		if err := router.Shutdown(); err != nil {
			log.Println("http: server shutdown: ", err)
		}
	}()

	// starting server
	ip := a.config.HTTPServer.ListenAddress
	if err := router.Listen(ip); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("failed to start server: %v", err)
		}
	}
}

package app

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func loggerMiddleware(a *App) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		if err := c.Next(); err != nil {
			log.Println("ERROR: request error: ", err)
			return err
		}
		params := []interface{}{
			"latency", time.Since(start).Seconds(),
			"client_ip", c.IP(),
			"method", string(c.Context().Method()),
			"status_code", c.Response().StatusCode(),
			"body_size", len(c.Response().Body()),
			"path", string(c.Context().URI().Path()),
		}

		log.Printf("Request success. %v", params)

		return nil
	}
}

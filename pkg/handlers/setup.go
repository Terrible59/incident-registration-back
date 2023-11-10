package handlers

import "github.com/gofiber/fiber/v2"

type SetupConfig struct {
	App *fiber.App
}

func SetupRoutes(config *SetupConfig) {
	api := config.App.Group("/api/v1")

	api.Get("/hello", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "hello",
		})
	})
}

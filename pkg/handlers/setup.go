package handlers

import (
	"github.com/gofiber/fiber/v2"
	db "incidents_back/db/sqlc"
)

type SetupConfig struct {
	App  *fiber.App
	Repo *db.Repo
}

func SetupRoutes(config *SetupConfig) {
	api := config.App.Group("/api/v1")

	api.Get("/hello", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "hello",
		})
	})

	incidents := api.Group("/incidents")
	incidents.Get("/", func(ctx *fiber.Ctx) error {
		inc, err := config.Repo.GetAllIncidents(ctx.Context())
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Something went wrong",
			})
		}

		return ctx.Status(fiber.StatusOK).JSON(inc)
	})
}

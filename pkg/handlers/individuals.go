package handlers

import (
	"database/sql"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
	"incidents_back/pkg/services"
)

type IndividualParams struct {
	FirstName    string `json:"first_name" validate:"required,min=1,max=15"`
	LastName     string `json:"last_name" validate:"required,min=1,max=15"`
	MiddleName   string `json:"middle_name" validate:"required,min=1,max=15"`
	Address      string `json:"address" validate:"required,min=1,max=50"`
	RecordsCount int64  `json:"criminal_records_count" validate:"required"`
}

func (h *Handlers) getAllIndividuals(c *fiber.Ctx) error {
	response, err := services.GetIndividuals(h.Repo)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Something went wrong",
		})
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func (h *Handlers) getIndividual(c *fiber.Ctx) error {
	individualID, err := uuid.FromString(c.Params("individualID"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Individual not found",
		})
	}

	response, err := services.GetIndividual(individualID, h.Repo)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Individual not found",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Something went wrong",
		})
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func (h *Handlers) createIndividual(c *fiber.Ctx) error {
	params := new(IndividualParams)

	if err := c.BodyParser(params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request parameters",
		})
	}

	err := h.Validator.Struct(params)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "Validation failed",
		})
	}

	response, err := services.CreateIndividual(params.FirstName, params.LastName, params.MiddleName, params.Address, params.RecordsCount, h.Repo)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Something went wrong",
		})
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func (h *Handlers) updateIndividual(c *fiber.Ctx) error {
	params := new(IndividualParams)

	if err := c.BodyParser(params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request parameters",
		})
	}

	err := h.Validator.Struct(params)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "Validation failed",
		})
	}

	individualID, err := uuid.FromString(c.Params("individualID"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Individual not found",
		})
	}

	response, err := services.UpdateIndividual(individualID, params.FirstName, params.LastName, params.MiddleName, params.Address, params.RecordsCount, h.Repo)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Individual not found",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Something went wrong",
		})
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func (h *Handlers) deleteIndividual(c *fiber.Ctx) error {
	individualID, err := uuid.FromString(c.Params("individualID"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Individual not found",
		})
	}

	err = services.DeleteIndividual(individualID, h.Repo)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Individual not found",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Something went wrong",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "ok",
	})
}

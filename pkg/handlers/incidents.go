package handlers

import (
	"database/sql"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
	"incidents_back/pkg/services"
)

type IncidentParams struct {
	IncidentType string `json:"incident_type" validate:"required,min=3,max=32"`
	Summary      string `json:"summary" validate:"required,min=6,max=512"`
}

func (h *Handlers) getAllIncidents(c *fiber.Ctx) error {
	response, err := services.GetIncidents(h.Repo)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Something went wrong",
		})
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func (h *Handlers) getIncident(c *fiber.Ctx) error {
	incidentID, err := uuid.FromString(c.Params("incidentID"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Incident not found",
		})
	}

	response, err := services.GetIncident(incidentID, h.Repo)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Incident not found",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Something went wrong",
		})
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func (h *Handlers) createIncident(c *fiber.Ctx) error {
	params := new(IncidentParams)

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

	response, err := services.CreateIncident(params.Summary, params.IncidentType, h.Repo)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Something went wrong",
		})
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func (h *Handlers) updateIncident(c *fiber.Ctx) error {
	params := new(IncidentParams)

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

	incidentID, err := uuid.FromString(c.Params("incidentID"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Incident not found",
		})
	}

	response, err := services.UpdateIncident(incidentID, params.Summary, params.IncidentType, h.Repo)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Incident not found",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Something went wrong",
		})
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func (h *Handlers) deleteIncident(c *fiber.Ctx) error {
	incidentID, err := uuid.FromString(c.Params("incidentID"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Incident not found",
		})
	}

	err = services.DeleteIncident(incidentID, h.Repo)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Incident not found",
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

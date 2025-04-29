package handlers

import (
	"github.com/gofiber/fiber/v3"
)

func (h *Handlers) GetBalance(c fiber.Ctx) error {
	id := c.Params("student_id")
	if id == "" {
		return c.Status(400).JSON("ID must be provided")
	}

	balance, err := h.service.GetBalance(id)
	if err != nil {
		return c.Status(500).JSON(err)
	}

	return c.Status(200).JSON(balance)
}

func (h *Handlers) GetHistory(c fiber.Ctx) error {
	id := c.Params("student_id")

	if id == "" {
		return c.Status(400).JSON("ID must be provided")
	}

	balance, err := h.service.GetHistory(id)
	if err != nil {
		h.logger.Error("Error getting history: ", err)
		return c.Status(500).JSON(err)
	}

	if balance == nil {
		return c.Status(400).JSON("balance is nil")
	}

	return c.Status(200).JSON(balance)
}

func (h *Handlers) GetHistoryOfTrips(c fiber.Ctx) error {
	id := c.Params("student_id")
	if id == "" {
		return c.Status(400).JSON("ID must be provided")
	}

	history, err := h.service.GetHistoryOfTrips(id)
	if err != nil {
		h.logger.Error("Error getting history: ", err)
		return c.Status(500).JSON(err)
	}

	if history == nil {
		return c.Status(400).JSON("history is nil")
	}

	return c.Status(200).JSON(*history)
}

func (h *Handlers) ActiveTickets(c fiber.Ctx) error {
	id := c.Params("student_id")
	if id == "" {
		return c.Status(400).JSON("ID must be provided")
	}

	resp, err := h.service.ActiveTickets(id)
	if err != nil {
		h.logger.Error("Error getting balance: ", err)
	}

	return c.Status(200).JSON(resp)
}

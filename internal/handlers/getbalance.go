package handlers

import "github.com/gofiber/fiber/v3"

func (h *Handlers) GetBalance(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON("ID must be provided")
	}

	balance, err := h.service.GetBalance(id)
	if err != nil {
		return c.Status(500).JSON(err)
	}

	return c.Status(200).JSON(balance)
}

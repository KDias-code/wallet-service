package handlers

import (
	"diplom/wallet-service/internal/models"
	"encoding/json"
	"github.com/gofiber/fiber/v3"
)

func (h *Handlers) ChangeBalance(c fiber.Ctx) error {
	body := c.Body()
	req := new(models.ChangeBalanceRequest)

	err := json.Unmarshal(body, &req)
	if err != nil {
		h.logger.Error("error unmarshalling body", "err", err)
		return c.Status(400).JSON("invalid request body")
	}

	isChanged, err := h.service.ChangeBalance(*req)
	if err != nil {
		h.logger.Error("error calling service", "err", err)
		return c.Status(500).JSON("error calling service")
	}

	if isChanged == 0.0 {
		return c.Status(400).JSON("Not enough balance to change")
	}

	return c.Status(200).JSON(isChanged)
}

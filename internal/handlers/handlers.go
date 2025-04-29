package handlers

import (
	"diplom/wallet-service/internal/service"
	"github.com/gofiber/fiber/v3"
	"github.com/hashicorp/go-hclog"
)

type IHandlers interface {
	GetBalance(c fiber.Ctx) error
	ChangeBalance(c fiber.Ctx) error
	GetHistory(c fiber.Ctx) error
	GetHistoryOfTrips(c fiber.Ctx) error
	ActiveTickets(c fiber.Ctx) error
}
type Handlers struct {
	logger  hclog.Logger
	service service.IService
}

func NewHandlers(logger hclog.Logger, service service.IService) *Handlers {
	return &Handlers{
		logger:  logger,
		service: service,
	}
}

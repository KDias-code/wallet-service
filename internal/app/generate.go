package app

import (
	"diplom/wallet-service/internal/handlers"
	"diplom/wallet-service/internal/service"
	"diplom/wallet-service/internal/store"
	"diplom/wallet-service/pkg/configs"
	"diplom/wallet-service/pkg/db"
	"github.com/gofiber/fiber/v3"
	"github.com/hashicorp/go-hclog"
)

type server struct {
	app      *fiber.App
	logger   hclog.Logger
	handlers handlers.IHandlers
}

func (s *server) generate(conf *configs.Config) {
	s.app = fiber.New()

	s.logger = hclog.New(&hclog.LoggerOptions{
		Level:      2,
		JSONFormat: true,
	})

	dbConnection, err := db.ConnectPostgres(conf.Db)
	if err != nil {
		s.logger.Error("Failed to connect to database", "error", err)
		return
	}

	stores := store.NewStore(dbConnection)
	services := service.NewServices(stores)
	s.handlers = handlers.NewHandlers(s.logger, services)

	s.router()
}

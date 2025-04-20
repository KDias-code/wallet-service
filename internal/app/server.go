package app

import "diplom/wallet-service/pkg/configs"

func Start(conf *configs.Config) error {
	s := new(server)

	s.generate(conf)

	err := s.app.Listen(":" + conf.Port)
	s.logger.Info("server started")
	if err != nil {
		s.logger.Error("Failed to start server", "error", err)
		return err
	}

	return err
}

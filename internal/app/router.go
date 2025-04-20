package app

func (s *server) router() {
	s.app.Get("/get-balance", s.handlers.GetBalance)
	s.app.Put("/change-balance", s.handlers.ChangeBalance)
}

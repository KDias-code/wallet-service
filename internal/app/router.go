package app

func (s *server) router() {
	s.app.Get("/get-balance/:student_id", s.handlers.GetBalance)
	s.app.Put("/change-balance", s.handlers.ChangeBalance)
	s.app.Get("/get-history/:student_id", s.handlers.GetHistory)
	s.app.Get("/get-history-trip/:student_id", s.handlers.GetHistoryOfTrips)
	s.app.Get("/get-active-tickets/:student_id", s.handlers.ActiveTickets)
}

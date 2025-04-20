package store

import "diplom/wallet-service/internal/models"

func (s *Store) GetBalance(id string) (models.GetBalanceResponse, error) {
	query := `SELECT id, balance FROM users WHERE id = ?`
	var resp models.GetBalanceResponse

	err := s.db.Select(&resp, query, id)
	if err != nil {
		return models.GetBalanceResponse{}, err
	}

	return resp, nil
}

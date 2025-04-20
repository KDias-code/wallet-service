package store

import "diplom/wallet-service/internal/models"

func (s *Store) ChangeBalance(req models.ChangeBalanceRequest) error {
	query := `UPDATE users SET balance = $1 WHERE id = $2`

	_, err := s.db.NamedExec(query, req)
	if err != nil {
		return err
	}

	return nil
}

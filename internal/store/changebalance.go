package store

import (
	"diplom/wallet-service/internal/models"
	"errors"
	"time"
)

func (s *Store) ChangeBalance(req models.ChangeBalanceRequest) error {
	busCheckQuery := `SELECT id FROM buses WHERE bus_number = $1`

	_, err := s.db.Query(busCheckQuery, req.BusNumber)
	if err != nil {
		return errors.New("не существует такого номера автобуса")
	}

	query := `UPDATE users SET balance = $1 WHERE student_id = $2`

	_, err = s.db.Exec(query, req.Change, req.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) HistoryAdd(amount float64, studentId, method, busNum string) error {
	now := time.Now()

	query := `INSERT INTO history (student_id, replenishment, change_date, method, bus_num) VALUES ($1, $2, $3, $4, $5)`

	_, err := s.db.Exec(query, studentId, amount, now, method, busNum)
	if err != nil {
		return err
	}

	return nil
}

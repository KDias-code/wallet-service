package store

import (
	"diplom/wallet-service/internal/models"
	"log"
)

func (s *Store) GetBalance(id string) (models.GetBalanceResponse, error) {
	query := `SELECT student_id, balance FROM users WHERE student_id = $1`
	var resp models.GetBalanceResponse

	err := s.db.Get(&resp, query, id)
	if err != nil {
		return models.GetBalanceResponse{}, err
	}

	return resp, nil
}

func (s *Store) GetHistory(studentId string) (*[]models.HistoryResponse, error) {
	query := `
		SELECT id, student_id, replenishment, change_date
		FROM history
		WHERE student_id = ?`

	var history []models.HistoryResponse

	err := s.db.Select(&history, query, studentId)
	if err != nil {
		log.Printf("Ошибка при получении истории для студента %s: %v", studentId, err)
		return nil, err
	}

	return &history, nil
}

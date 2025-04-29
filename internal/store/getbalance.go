package store

import (
	"diplom/wallet-service/internal/models"
	"log"
	"time"
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
		SELECT id, student_id, replenishment, change_date, method, bus_num
		FROM history
		WHERE student_id = $1`

	var history []models.HistoryResponse

	err := s.db.Select(&history, query, studentId)
	if err != nil {
		log.Printf("Ошибка при получении истории для студента %s: %v", studentId, err)
		return nil, err
	}

	return &history, nil
}

func (s *Store) HistoryOfTrips(studentId string) (*[]models.HistoryResponse, error) {
	query := `
		SELECT id, student_id, replenishment, change_date, method
		FROM history
		WHERE student_id = $1 AND method = 'paid'`

	var history []models.HistoryResponse

	err := s.db.Select(&history, query, studentId)
	if err != nil {
		log.Printf("Ошибка при получении истории для студента %s: %v", studentId, err)
		return nil, err
	}

	return &history, nil
}

func (s *Store) ActiveTickets(studentId string, timeBefore time.Time) ([]models.HistoryResponse, error) {
	query := `
        SELECT id, student_id, replenishment, change_date, bus_num, method
        FROM history
        WHERE student_id = $1
          AND change_date >= $2
          AND change_date <= $3
          AND method = 'paid'
        ORDER BY change_date DESC
    `

	rows, err := s.db.Query(query, studentId, timeBefore.Format("15:04:05"), time.Now().Format("15:04:05"))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []models.HistoryResponse
	for rows.Next() {
		var h models.HistoryResponse
		err := rows.Scan(&h.ID, &h.StudentID, &h.Replenishment, &h.ChangeDate, &h.BusNum, &h.Method)
		if err != nil {
			return nil, err
		}
		result = append(result, h)
	}

	return result, nil
}

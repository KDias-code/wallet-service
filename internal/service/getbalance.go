package service

import (
	"diplom/wallet-service/internal/models"
	"time"
)

func (s *Service) GetBalance(id string) (models.GetBalanceResponse, error) {
	response, err := s.store.GetBalance(id)
	if err != nil {
		return models.GetBalanceResponse{}, err
	}

	return response, nil
}

func (s *Service) GetHistory(studentId string) (*[]models.HistoryResponse, error) {
	response, err := s.store.GetHistory(studentId)
	if err != nil {
		return nil, err
	}

	if response == nil {
		return nil, nil
	}

	return response, nil
}

func (s *Service) GetHistoryOfTrips(studentId string) (*[]models.HistoryResponse, error) {
	resp, err := s.store.HistoryOfTrips(studentId)
	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, nil
	}

	return resp, nil
}

func (s *Service) ActiveTickets(studentId string) ([]models.HistoryResponse, error) {
	now := time.Now()
	thirtyMinutesAgo := now.Add(-30 * time.Minute)

	resp, err := s.store.ActiveTickets(studentId, thirtyMinutesAgo)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

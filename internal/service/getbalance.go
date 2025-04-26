package service

import (
	"diplom/wallet-service/internal/models"
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

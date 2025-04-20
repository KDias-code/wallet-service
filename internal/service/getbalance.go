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

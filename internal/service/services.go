package service

import (
	"diplom/wallet-service/internal/models"
	"diplom/wallet-service/internal/store"
)

type IService interface {
	GetBalance(id string) (models.GetBalanceResponse, error)
	ChangeBalance(req models.ChangeBalanceRequest) (float64, error)
}
type Service struct {
	store store.IStore
}

func NewServices(store store.IStore) *Service {
	return &Service{
		store: store,
	}
}

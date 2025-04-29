package service

import (
	"diplom/wallet-service/internal/models"
	"diplom/wallet-service/internal/store"
)

type IService interface {
	GetBalance(id string) (models.GetBalanceResponse, error)
	ChangeBalance(req models.ChangeBalanceRequest) (float64, error)
	GetHistory(studentId string) (*[]models.HistoryResponse, error)
	GetHistoryOfTrips(studentId string) (*[]models.HistoryResponse, error)
	ActiveTickets(studentId string) ([]models.HistoryResponse, error)
}
type Service struct {
	store        store.IStore
	secretStripe string
}

func NewServices(store store.IStore, secretStripe string) *Service {
	return &Service{
		store:        store,
		secretStripe: secretStripe,
	}
}

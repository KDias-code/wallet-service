package store

import (
	"diplom/wallet-service/internal/models"
	"github.com/jmoiron/sqlx"
	"time"
)

type IStore interface {
	GetBalance(id string) (models.GetBalanceResponse, error)
	ChangeBalance(req models.ChangeBalanceRequest) error
	HistoryAdd(amount float64, studentId, method, busNum string) error
	GetHistory(studentId string) (*[]models.HistoryResponse, error)
	HistoryOfTrips(studentId string) (*[]models.HistoryResponse, error)
	ActiveTickets(studentId string, now time.Time) ([]models.HistoryResponse, error)
}
type Store struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) *Store {
	return &Store{db: db}
}

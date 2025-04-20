package store

import (
	"diplom/wallet-service/internal/models"
	"github.com/jmoiron/sqlx"
)

type IStore interface {
	GetBalance(id string) (models.GetBalanceResponse, error)
	ChangeBalance(req models.ChangeBalanceRequest) error
}
type Store struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) *Store {
	return &Store{db: db}
}

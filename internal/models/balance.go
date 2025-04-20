package models

type GetBalanceResponse struct {
	ID      string  `json:"id" db:"id"`
	Balance float64 `json:"balance" db:"balance"`
}

type ChangeBalanceRequest struct {
	ID     string  `json:"id" db:"id"`
	Change float64 `json:"change" db:"balance"`
}

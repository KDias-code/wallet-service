package models

type GetBalanceResponse struct {
	ID      string  `json:"student_id" db:"student_id"`
	Balance float64 `json:"balance" db:"balance"`
}

type ChangeBalanceRequest struct {
	ID        string  `json:"student_id" db:"student_id"`
	Change    float64 `json:"change" db:"balance"`
	BusNumber string  `json:"bus_number" db:"bus_number"`
}

type HistoryResponse struct {
	ID            int     `json:"id" db:"id"`
	StudentID     string  `json:"student_id" db:"student_id"`
	Replenishment float64 `json:"replenishment" db:"replenishment"`
	ChangeDate    string  `json:"change_date" db:"change_date"`
	Method        string  `json:"method" db:"method"`
}

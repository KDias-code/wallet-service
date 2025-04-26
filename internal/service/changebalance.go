package service

import (
	"diplom/wallet-service/internal/models"
	"diplom/wallet-service/pkg/stripe"
)

func (s *Service) ChangeBalance(req models.ChangeBalanceRequest) (float64, error) {
	balance, err := s.store.GetBalance(req.ID)
	if err != nil {
		return 0, err
	}

	if req.BusNumber == "" {
		req.Change += balance.Balance
		err = s.store.ChangeBalance(req)
		if err != nil {
			return 0, err
		}

		err = s.store.HistoryAdd(req.Change, req.ID)
		if err != nil {
			return 0, err
		}
	} else {
		amount := balance.Balance - req.Change
		if amount < 0 {
			return 0, nil
		}

		req.Change = amount

		err = stripe.Payment(int(req.Change), s.secretStripe)
		if err != nil {
			return 0, err
		}

		err = s.store.ChangeBalance(req)
		if err != nil {
			return 0, err
		}

		err = s.store.HistoryAdd(req.Change, req.ID)
		if err != nil {
			return 0, err
		}
	}

	return req.Change, nil
}

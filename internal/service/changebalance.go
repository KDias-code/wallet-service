package service

import "diplom/wallet-service/internal/models"

func (s *Service) ChangeBalance(req models.ChangeBalanceRequest) (float64, error) {
	balance, err := s.store.GetBalance(req.ID)
	if err != nil {
		return 0, err
	}

	if req.Change > 0 {
		req.Change += balance.Balance
		err = s.store.ChangeBalance(req)
		if err != nil {
			return 0, err
		}
	} else if req.Change < 0 {
		amount := balance.Balance - req.Change
		if amount < 0 {
			return 0, nil
		}

		req.Change = amount
		err = s.store.ChangeBalance(req)
		if err != nil {
			return 0, err
		}
	}

	return req.Change, nil
}

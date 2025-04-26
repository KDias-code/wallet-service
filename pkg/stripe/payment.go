package stripe

import (
	"github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/paymentintent"
)

func Payment(amount int, secret string) error {
	stripe.Key = secret
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(30000),
		Currency: stripe.String("kzt"),
	}

	_, err := paymentintent.New(params)
	if err != nil {
		return err
	}

	return nil
}

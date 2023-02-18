package domain

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type CreditCard struct {
	id              string
	name            string
	number          string
	expirationMonth int32
	expirationYear  int32
	cvv             int32
	balance         float64
	limit           float64
	createdAt       time.Time
}

func NewCreditCard() CreditCard {
	c := CreditCard{}

	c.id = uuid.NewV4().String()
	c.createdAt = time.Now()
}

package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type TransactionRepository interface {
	SaveTransaction(transaction Transaction, creditCard CreditCard) error
	GetCreditCard(creditCard CreditCard) (CreditCard, error)
	CreateCreditCard(creditCard CreditCard) error
}

type Transaction struct {
	id           string
	amount       float64
	status       string
	description  string
	store        string
	creditCardId string
	createdAt    time.Time
}

func NewTransaction() *Transaction {
	t := &Transaction{}
	t.id = uuid.NewV4().String()
	t.createdAt = time.Now()
	return t
}

func (t *Transaction) ProcessAndValidate(creditcard *CreditCard) {
	if t.amount+creditcard.balance > creditcard.limit {
		t.status = "rejected"
	} else {
		t.status = "approved"
		creditcard.balance += t.amount
	}

}

package repository

import (
	"database/sql"

	"github.com/pedro77h/code-bamk/domain"
)

type TransactionRepositoryDb struct {
	db *sql.DB
}

func NewTransactionRepositoryDB(db *sql.DB) *TransactionRepositoryDb {
	return &TransactionRepositoryDb{db: db}
}

func (t *TransactionRepositoryDb) SaveTransaction(transaction domain.Transaction, creditCard domain.Transaction) error {
	stmt, err := t.db.Prepare(`insert into transactions(id , credit_card_id , amount, status, description, store , created_at)
							values($1 ,$2, $3, $4 , $5, $6, $7 )`)
}

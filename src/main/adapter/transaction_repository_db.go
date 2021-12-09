package adapter

import (
	"database/sql"
	"time"
)

type TransactionRepositoryDB struct {
	db *sql.DB
}

func NewTransactionRepositoryDB(db *sql.DB) *TransactionRepositoryDB {
	return &TransactionRepositoryDB{
		db: db,
	}
}

func (t *TransactionRepositoryDB) Insert(ID string, AccountID string, Amount float64, Status string, ErrorMessage string) error {
	query := `
		INSERT INTO transaction(id, account_id, amount, status, error_message, created_at, updated_at)
		VALUES($1, $2, $3, $4, $5, $6, $7)
	`
	stmt, err := t.db.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		ID,
		AccountID,
		Amount,
		Status,
		ErrorMessage,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return err
	}

	return nil

}

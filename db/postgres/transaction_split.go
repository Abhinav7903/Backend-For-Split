package postgres

import (
	"fmt"

	"github.com/Abhinav7903/split/factory"
)

func (p *Postgres) CreateTransactionSplit(transaction *factory.TransactionSplit) error {
	query := `
        INSERT INTO transaction_splits (transaction_id, user_id, amount)
        VALUES ($1, $2, $3)
        RETURNING transaction_split_id
    `

	tx, err := p.dbConn.Begin() // Start a transaction
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	// Perform the query within the transaction
	err = tx.QueryRow(query, transaction.TransactionID, transaction.UserID, transaction.Amount).
		Scan(&transaction.TransactionSplitID)
	if err != nil {
		tx.Rollback() // Roll back the transaction on error
		return fmt.Errorf("error creating transaction split: %w", err)
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (p *Postgres) GetTransactionSplits(transactionID int) ([]factory.TransactionSplit, error) {
	query := `
		SELECT transaction_split_id, transaction_id, user_id, amount
		FROM transaction_splits
		WHERE transaction_id = $1
	`

	rows, err := p.dbConn.Query(query, transactionID)
	if err != nil {
		return nil, fmt.Errorf("error querying transaction splits: %w", err)
	}
	defer rows.Close()

	// Preallocate slice if you know approximate size (optional)
	transactionSplits := make([]factory.TransactionSplit, 0)

	for rows.Next() {
		var transactionSplit factory.TransactionSplit
		if err := rows.Scan(
			&transactionSplit.TransactionSplitID,
			&transactionSplit.TransactionID,
			&transactionSplit.UserID,
			&transactionSplit.Amount,
		); err != nil {
			return nil, fmt.Errorf("error scanning transaction split: %w", err)
		}
		transactionSplits = append(transactionSplits, transactionSplit)
	}

	// Check for errors encountered during iteration
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating transaction splits: %w", err)
	}

	return transactionSplits, nil
}

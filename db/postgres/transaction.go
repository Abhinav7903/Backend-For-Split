package postgres

import (
	"database/sql"
	"errors"

	"github.com/Abhinav7903/split/factory"
)

func (p *Postgres) CreateTransaction(transaction *factory.Transaction) (int, error) {
	query := `INSERT INTO transactions 
              (lender_id, borrower_id, group_id, amount, status, purpose, payment_method_id, retry_count, failure_reason)
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING transaction_id`

	// Use QueryRow to execute the query
	row := p.dbConn.QueryRow(
		query,
		transaction.LenderID,
		transaction.BorrowerID,
		transaction.GroupID, // Allow for nullable fields
		transaction.Amount,
		transaction.Status,
		transaction.Purpose,         // Allow for nullable fields
		transaction.PaymentMethodID, // Allow for nullable fields
		transaction.RetryCount,
		transaction.FailureReason, // Allow for nullable fields
	)

	var transactionID int
	err := row.Scan(&transactionID) // Scan the RETURNING result into transactionID
	if err != nil {
		return 0, err // Return an error if the operation fails
	}

	return transactionID, nil
}

func (p *Postgres) GetTransactionByID(transactionID int) (*factory.Transaction, error) {
	query := `SELECT 
                 transaction_id, lender_id, borrower_id, group_id, amount, 
                 status, purpose, payment_method_id, retry_count, failure_reason
              FROM transactions 
              WHERE transaction_id = $1`

	row := p.dbConn.QueryRow(query, transactionID)

	// Define a variable to hold the transaction
	var transaction factory.Transaction

	// Scan the result into the transaction struct
	err := row.Scan(
		&transaction.TransactionID,
		&transaction.LenderID,
		&transaction.BorrowerID,
		&transaction.GroupID, // Nullable field
		&transaction.Amount,
		&transaction.Status,
		&transaction.Purpose,         // Nullable field
		&transaction.PaymentMethodID, // Nullable field
		&transaction.RetryCount,
		&transaction.FailureReason, // Nullable field
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Return nil if no transaction is found
		}
		return nil, err // Return the error for other cases
	}

	return &transaction, nil
}

func (p *Postgres) GetTransactionsByLenderID(lenderID int) ([]factory.Transaction, error) {
	query := `SELECT 
				 transaction_id, lender_id, borrower_id, group_id, amount, 
				 status, purpose, payment_method_id, retry_count, failure_reason
			  FROM transactions 
			  WHERE lender_id = $1`

	rows, err := p.dbConn.Query(query, lenderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []factory.Transaction
	for rows.Next() {
		var transaction factory.Transaction
		err = rows.Scan(
			&transaction.TransactionID,
			&transaction.LenderID,
			&transaction.BorrowerID,
			&transaction.GroupID, // Nullable field
			&transaction.Amount,
			&transaction.Status,
			&transaction.Purpose,         // Nullable field
			&transaction.PaymentMethodID, // Nullable field
			&transaction.RetryCount,
			&transaction.FailureReason, // Nullable field
		)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (p *Postgres) GetTransactionsByBorrowerID(borrowerID int) ([]factory.Transaction, error) {
	query := `SELECT 
				 transaction_id, lender_id, borrower_id, group_id, amount, 
				 status, purpose, payment_method_id, retry_count, failure_reason
			  FROM transactions 
			  WHERE borrower_id = $1`

	rows, err := p.dbConn.Query(query, borrowerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []factory.Transaction
	for rows.Next() {
		var transaction factory.Transaction
		err = rows.Scan(
			&transaction.TransactionID,
			&transaction.LenderID,
			&transaction.BorrowerID,
			&transaction.GroupID, // Nullable field
			&transaction.Amount,
			&transaction.Status,
			&transaction.Purpose,         // Nullable field
			&transaction.PaymentMethodID, // Nullable field
			&transaction.RetryCount,
			&transaction.FailureReason, // Nullable field
		)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (p *Postgres) UpdateTransactionStatus(transactionID int, status string) error {
	query := `UPDATE transactions SET status = $1 WHERE transaction_id = $2`

	_, err := p.dbConn.Exec(query, status, transactionID)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) DeleteTransaction(transactionID int) error {
	query := `DELETE FROM transactions WHERE transaction_id = $1`

	_, err := p.dbConn.Exec(query, transactionID)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) SearchTransactions(filters factory.TransactionFilters) ([]factory.Transaction, error) {
	query := `SELECT 
                 transaction_id, lender_id, borrower_id, group_id, amount, 
                 status, purpose, payment_method_id, retry_count, failure_reason
              FROM transactions 
              WHERE lender_id = COALESCE($1, lender_id)
              AND borrower_id = COALESCE($2, borrower_id)
              AND group_id = COALESCE($3, group_id)
              AND status = COALESCE($4, status)
              AND amount >= COALESCE($5, amount)
              AND amount <= COALESCE($6, amount)
              AND payment_method_id = COALESCE($7, payment_method_id)`

	// Ensure at least one filter is provided
	if filters.LenderID == nil && filters.BorrowerID == nil && filters.GroupID == nil &&
		filters.Status == nil && filters.MinAmount == nil && filters.MaxAmount == nil &&
		filters.PaymentMethodID == nil {
		return nil, errors.New("at least one filter must be provided")
	}

	// Execute the query with the filters
	rows, err := p.dbConn.Query(
		query,
		filters.LenderID,
		filters.BorrowerID,
		filters.GroupID,
		filters.Status,
		filters.MinAmount,
		filters.MaxAmount,
		filters.PaymentMethodID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []factory.Transaction
	for rows.Next() {
		var transaction factory.Transaction
		err := rows.Scan(
			&transaction.TransactionID,
			&transaction.LenderID,
			&transaction.BorrowerID,
			&transaction.GroupID, // Nullable field
			&transaction.Amount,
			&transaction.Status,
			&transaction.Purpose,         // Nullable field
			&transaction.PaymentMethodID, // Nullable field
			&transaction.RetryCount,
			&transaction.FailureReason, // Nullable field
		)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	// Check for row iteration errors
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}

package postgres

import (
	"database/sql"
	"errors"

	"github.com/Abhinav7903/split/factory"
)

func (p *Postgres) CreateTransaction(transaction *factory.Transaction) (int, error) {
	// Validate that lender exists
	lenderExists, err := p.CheckUserExists(transaction.LenderID)
	if err != nil || !lenderExists {
		return 0, errors.New("lender does not exist")
	}

	// Validate that borrower exists
	borrowerExists, err := p.CheckUserExists(transaction.BorrowerID)
	if err != nil || !borrowerExists {
		return 0, errors.New("borrower does not exist")
	}

	// Validate that group exists
	groupExists, err := p.CheckGroupExists(transaction.GroupID)
	if err != nil || !groupExists {
		return 0, errors.New("group does not exist")
	}

	// Validate that payment method exists
	paymentMethodExists, err := p.CheckPaymentMethodExists(transaction.PaymentMethodID)
	if err != nil || !paymentMethodExists {
		return 0, errors.New("payment method does not exist")
	}

	// Proceed with inserting the transaction if all foreign key checks pass
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
	err = row.Scan(&transactionID) // Scan the RETURNING result into transactionID
	if err != nil {
		return 0, err // Return an error if the operation fails
	}

	return transactionID, nil
}

func (p *Postgres) GetTransactionByID(transactionID int) (*factory.Transaction, error) {
	// Check if the transaction exists in the database before fetching it
	var exists bool
	checkQuery := `SELECT EXISTS (SELECT 1 FROM transactions WHERE transaction_id = $1)`
	err := p.dbConn.QueryRow(checkQuery, transactionID).Scan(&exists)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("transaction not found") // Return a more descriptive error
	}

	query := `SELECT 
				 transaction_id, lender_id, borrower_id, group_id, amount, 
				 status, purpose, payment_method_id, retry_count, failure_reason
			  FROM transactions 
			  WHERE transaction_id = $1`

	row := p.dbConn.QueryRow(query, transactionID)

	// Define a variable to hold the transaction
	var transaction factory.Transaction

	// Scan the result into the transaction struct
	err = row.Scan(
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
	// Ensure the lender exists before fetching transactions
	var exists bool
	checkQuery := `SELECT EXISTS (SELECT 1 FROM users WHERE user_id = $1)`
	err := p.dbConn.QueryRow(checkQuery, lenderID).Scan(&exists)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("lender not found")
	}

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
	// Ensure the borrower exists before fetching transactions
	var exists bool
	checkQuery := `SELECT EXISTS (SELECT 1 FROM users WHERE user_id = $1)`
	err := p.dbConn.QueryRow(checkQuery, borrowerID).Scan(&exists)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("borrower not found")
	}

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
	// Ensure the transaction exists before updating
	var exists bool
	checkQuery := `SELECT EXISTS (SELECT 1 FROM transactions WHERE transaction_id = $1)`
	err := p.dbConn.QueryRow(checkQuery, transactionID).Scan(&exists)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("transaction not found")
	}

	query := `UPDATE transactions SET status = $1 WHERE transaction_id = $2`
	_, err = p.dbConn.Exec(query, status, transactionID)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) DeleteTransaction(transactionID int) error {
	// Ensure the transaction exists before deleting
	var exists bool
	checkQuery := `SELECT EXISTS (SELECT 1 FROM transactions WHERE transaction_id = $1)`
	err := p.dbConn.QueryRow(checkQuery, transactionID).Scan(&exists)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("transaction not found")
	}

	query := `DELETE FROM transactions WHERE transaction_id = $1`
	_, err = p.dbConn.Exec(query, transactionID)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) SearchTransactions(filters factory.TransactionFilters) ([]factory.Transaction, error) {
	// Ensure at least one filter is provided
	if filters.LenderID == nil && filters.BorrowerID == nil && filters.GroupID == nil &&
		filters.Status == nil && filters.MinAmount == nil && filters.MaxAmount == nil &&
		filters.PaymentMethodID == nil {
		return nil, errors.New("at least one filter must be provided")
	}

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

// Helper function to check if a user exists
func (p *Postgres) CheckUserExists(userID int) (bool, error) {
	query := `SELECT EXISTS (SELECT 1 FROM users WHERE user_id = $1)`
	var exists bool
	err := p.dbConn.QueryRow(query, userID).Scan(&exists)
	return exists, err
}

// Helper function to check if a group exists
func (p *Postgres) CheckGroupExists(groupID int) (bool, error) {
	query := `SELECT EXISTS (SELECT 1 FROM groups WHERE group_id = $1)`
	var exists bool
	err := p.dbConn.QueryRow(query, groupID).Scan(&exists)
	return exists, err
}

// Helper function to check if a payment method exists
func (p *Postgres) CheckPaymentMethodExists(paymentMethodID int) (bool, error) {
	query := `SELECT EXISTS (SELECT 1 FROM payment_methods WHERE payment_id = $1)`
	var exists bool
	err := p.dbConn.QueryRow(query, paymentMethodID).Scan(&exists)
	return exists, err
}

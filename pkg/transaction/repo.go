package transaction

import "github.com/Abhinav7903/split/factory"

type Repository interface {
	// Create a new transaction
	CreateTransaction(transaction *factory.Transaction) (int, error)

	// Retrieve a transaction by ID
	GetTransactionByID(transactionID int) (*factory.Transaction, error)

	// Retrieve transactions by Lender ID
	GetTransactionsByLenderID(lenderID int) ([]factory.Transaction, error)

	// Retrieve transactions by Borrower ID
	GetTransactionsByBorrowerID(borrowerID int) ([]factory.Transaction, error)

	// Update the status of a transaction
	UpdateTransactionStatus(transactionID int, status string) error

	// Delete a transaction by ID
	DeleteTransaction(transactionID int) error

	// Custom queries for filtering or searching
	SearchTransactions(filters factory.TransactionFilters) ([]factory.Transaction, error)
}

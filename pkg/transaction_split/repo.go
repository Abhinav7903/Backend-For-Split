package transactionsplit

import "github.com/Abhinav7903/split/factory"

type Repository interface {
	CreateTransactionSplit(transaction *factory.TransactionSplit) error
	GetTransactionSplits(transactionID int) ([]factory.TransactionSplit, error)
}

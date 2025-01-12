package balance

import "github.com/Abhinav7903/split/factory"

type Repository interface {
	CreateBalance(balance *factory.Balance) (int, error)
	GetBalanceByID(balanceID int) (*factory.Balance, error)
	GetBalancesByUserID(userID int) ([]factory.Balance, error)
	GetBalancesByGroupID(groupID int) ([]factory.Balance, error)
	UpdateBalanceAmounts(balanceID int, owedAmount, lentAmount *float64) error
	DeleteBalance(balanceID int) error
}

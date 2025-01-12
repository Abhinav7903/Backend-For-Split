package postgres

import (
	"database/sql"
	"errors"

	"github.com/Abhinav7903/split/factory"
)

func (p *Postgres) CreateBalance(balance *factory.Balance) (int, error) {
	// Validate that user exists
	userExists, err := p.CheckUserExists(balance.UserID)
	if err != nil || !userExists {
		return 0, errors.New("user does not exist")
	}

	// Validate that group exists (if group_id is provided)
	if balance.GroupID != nil {
		groupExists, err := p.CheckGroupExists(*balance.GroupID)
		if err != nil || !groupExists {
			return 0, errors.New("group does not exist")
		}
	}

	// Insert balance into the database
	query := `INSERT INTO balances (user_id, group_id, owed_amount, lent_amount)
	          VALUES ($1, $2, $3, $4) RETURNING balance_id`

	var balanceID int
	err = p.dbConn.QueryRow(query, balance.UserID, balance.GroupID, balance.OwedAmount, balance.LentAmount).Scan(&balanceID)
	if err != nil {
		return 0, err
	}

	return balanceID, nil
}

func (p *Postgres) GetBalanceByID(balanceID int) (*factory.Balance, error) {
	query := `SELECT balance_id, user_id, group_id, owed_amount, lent_amount 
	          FROM balances WHERE balance_id = $1`

	var balance factory.Balance
	row := p.dbConn.QueryRow(query, balanceID)

	// Scan the result into the balance struct
	err := row.Scan(&balance.BalanceID, &balance.UserID, &balance.GroupID, &balance.OwedAmount, &balance.LentAmount)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("balance not found")
		}
		return nil, err
	}

	return &balance, nil
}

func (p *Postgres) GetBalancesByUserID(userID int) ([]factory.Balance, error) {
	// Ensure the user exists
	userExists, err := p.CheckUserExists(userID)
	if err != nil || !userExists {
		return nil, errors.New("user not found")
	}

	query := `SELECT balance_id, user_id, group_id, owed_amount, lent_amount 
	          FROM balances WHERE user_id = $1`

	rows, err := p.dbConn.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var balances []factory.Balance
	for rows.Next() {
		var balance factory.Balance
		err = rows.Scan(&balance.BalanceID, &balance.UserID, &balance.GroupID, &balance.OwedAmount, &balance.LentAmount)
		if err != nil {
			return nil, err
		}
		balances = append(balances, balance)
	}

	return balances, nil
}

func (p *Postgres) GetBalancesByGroupID(groupID int) ([]factory.Balance, error) {
	// Ensure the group exists
	groupExists, err := p.CheckGroupExists(groupID)
	if err != nil || !groupExists {
		return nil, errors.New("group not found")
	}

	query := `SELECT balance_id, user_id, group_id, owed_amount, lent_amount 
	          FROM balances WHERE group_id = $1`

	rows, err := p.dbConn.Query(query, groupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var balances []factory.Balance
	for rows.Next() {
		var balance factory.Balance
		err = rows.Scan(&balance.BalanceID, &balance.UserID, &balance.GroupID, &balance.OwedAmount, &balance.LentAmount)
		if err != nil {
			return nil, err
		}
		balances = append(balances, balance)
	}

	return balances, nil
}

func (p *Postgres) UpdateBalanceAmounts(balanceID int, owedAmount, lentAmount *float64) error {
	// Ensure the balance exists
	var exists bool
	checkQuery := `SELECT EXISTS (SELECT 1 FROM balances WHERE balance_id = $1)`
	err := p.dbConn.QueryRow(checkQuery, balanceID).Scan(&exists)
	if err != nil || !exists {
		return errors.New("balance not found")
	}

	query := `UPDATE balances 
	          SET owed_amount = COALESCE($1, owed_amount), 
	              lent_amount = COALESCE($2, lent_amount)
	          WHERE balance_id = $3`

	_, err = p.dbConn.Exec(query, owedAmount, lentAmount, balanceID)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) DeleteBalance(balanceID int) error {
	// Ensure the balance exists
	var exists bool
	checkQuery := `SELECT EXISTS (SELECT 1 FROM balances WHERE balance_id = $1)`
	err := p.dbConn.QueryRow(checkQuery, balanceID).Scan(&exists)
	if err != nil || !exists {
		return errors.New("balance not found")
	}

	query := `DELETE FROM balances WHERE balance_id = $1`
	_, err = p.dbConn.Exec(query, balanceID)
	if err != nil {
		return err
	}

	return nil
}

// Helper function to check if a balance exists
func (p *Postgres) CheckBalanceExists(balanceID int) (bool, error) {
	query := `SELECT EXISTS (SELECT 1 FROM balances WHERE balance_id = $1)`
	var exists bool
	err := p.dbConn.QueryRow(query, balanceID).Scan(&exists)
	return exists, err
}

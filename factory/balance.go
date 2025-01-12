package factory

type Balance struct {
	BalanceID  int     `json:"balance_id"`
	UserID     int     `json:"user_id"`
	GroupID    *int    `json:"group_id,omitempty"`
	OwedAmount float64 `json:"owed_amount"`
	LentAmount float64 `json:"lent_amount"`
}

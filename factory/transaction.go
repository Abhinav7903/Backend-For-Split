package factory

type Transaction struct {
	TransactionID   int     `json:"transaction_id"`
	LenderID        int     `json:"lender_id"`
	BorrowerID      int     `json:"borrower_id"`
	GroupID         *int    `json:"group_id"` // Pointer for nullable field
	Amount          float64 `json:"amount"`
	Status          string  `json:"status"`
	Purpose         *string `json:"purpose"`           // Pointer for nullable field
	PaymentMethodID *int    `json:"payment_method_id"` // Pointer for nullable field
	RetryCount      int     `json:"retry_count"`
	FailureReason   *string `json:"failure_reason"` // Pointer for nullable field
}

type TransactionFilters struct {
	LenderID        *int
	BorrowerID      *int
	GroupID         *int
	Status          *string
	MinAmount       *float64
	MaxAmount       *float64
	PaymentMethodID *int
}

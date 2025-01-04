package factory

import (
	"errors"
	"time"
)

// PaymentMethod represents a payment method in the system
type PaymentMethod struct {
	PaymentID      int       `json:"payment_id"`                // Primary Key
	UserID         int       `json:"user_id"`                   // ID of the user owning this payment method
	Email          string    `json:"email,omitempty"`           // Email of the user owning this payment method
	PaymentType    string    `json:"payment_type"`              // Either 'UPI' or 'Bank Account'
	UPIID          *string   `json:"upi_id,omitempty"`          // UPI ID for UPI payment methods
	AccountNumber  *string   `json:"account_number,omitempty"`  // Bank account number
	IFSCCode       *string   `json:"ifsc_code,omitempty"`       // IFSC code for bank accounts
	WalletProvider *string   `json:"wallet_provider,omitempty"` // Optional, e.g., Google Pay, Paytm
	IsPrimary      bool      `json:"is_primary"`                // Whether this is the user's primary payment method
	CreatedAt      time.Time `json:"created_at"`                // Timestamp when the method was created
}

// ValidPaymentTypes defines the allowed types of payments
var ValidPaymentTypes = []string{"UPI", "Bank Account"}

// Validate checks the validity of a PaymentMethod object
	func (p *PaymentMethod) Validate() error {
		if p.UserID <= 0 {
			return errors.New("user_id is required and must be positive")
		}
		if p.PaymentType == "" {
			return errors.New("payment_type is required")
		}

		switch p.PaymentType {
		case "UPI":
			if p.UPIID == nil || *p.UPIID == "" {
				return errors.New("upi_id is required for UPI payment methods")
			}
		case "Bank Account":
			if p.AccountNumber == nil || *p.AccountNumber == "" {
				return errors.New("account_number is required for bank accounts")
			}
			if p.IFSCCode == nil || *p.IFSCCode == "" {
				return errors.New("ifsc_code is required for bank accounts")
			}
		default:
			return errors.New("payment_type must be 'UPI' or 'Bank Account'")
		}
		return nil
	}

// Helper function to check if a string exists in a slice
func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
